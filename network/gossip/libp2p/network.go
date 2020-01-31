package libp2p

import (
	"fmt"
	"hash"
	"strconv"
	"sync"

	"github.com/dchest/siphash"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/dapperlabs/flow-go/model/flow"
	"github.com/dapperlabs/flow-go/module"
	"github.com/dapperlabs/flow-go/network"
	"github.com/dapperlabs/flow-go/network/gossip/libp2p/cache"
	"github.com/dapperlabs/flow-go/network/gossip/libp2p/message"
	"github.com/dapperlabs/flow-go/network/gossip/libp2p/middleware"
	"github.com/dapperlabs/flow-go/network/gossip/libp2p/topology"
	"github.com/dapperlabs/flow-go/protocol"
)

// Network represents the overlay network of our peer-to-peer network, including
// the protocols for handshakes, authentication, gossiping and heartbeats.
type Network struct {
	sync.RWMutex
	logger  zerolog.Logger
	codec   network.Codec
	state   protocol.State
	me      module.Local
	mw      middleware.Middleware
	cache   *cache.Cache
	top     *topology.Topology
	sip     hash.Hash
	engines map[uint8]network.Engine
}

// NewNetwork creates a new naive overlay network, using the given middleware to
// communicate to direct peers, using the given codec for serialization, and
// using the given state & cache interfaces to track volatile information.
func NewNetwork(log zerolog.Logger, codec network.Codec, state protocol.State, me module.Local, mw middleware.Middleware) (*Network, error) {

	top, err := topology.New()
	if err != nil {
		return nil, errors.Wrap(err, "could not initialize topology")
	}

	ca, err := cache.New()
	if err != nil {
		return nil, errors.Wrap(err, "could not initialize cache")
	}

	o := &Network{
		logger:  log,
		codec:   codec,
		state:   state,
		me:      me,
		mw:      mw,
		top:     top,
		cache:   ca,
		sip:     siphash.New([]byte("daflowtrickleson")),
		engines: make(map[uint8]network.Engine),
	}

	return o, nil
}

// Ready returns a channel that will close when the network stack is ready.
func (n *Network) Ready() <-chan struct{} {
	ready := make(chan struct{})
	go func() {
		err := n.mw.Start(n)
		n.logger.Err(err).Msg("failed to start middleware")
		close(ready)
	}()
	return ready
}

// Done returns a channel that will close when shutdown is complete.
func (n *Network) Done() <-chan struct{} {
	done := make(chan struct{})
	go func() {
		n.mw.Stop()
		close(done)
	}()
	return done
}

// Register will register the given engine with the given unique engine engineID,
// returning a conduit to directly submit messages to the message bus of the
// engine.
func (n *Network) Register(channelID uint8, engine network.Engine) (network.Conduit, error) {
	n.Lock()
	defer n.Unlock()

	// check if the engine engineID is already taken
	_, ok := n.engines[channelID]
	if ok {
		return nil, fmt.Errorf("engine already registered (%d)", engine)
	}

	// Register the middleware for the channelID topic
	err := n.mw.Subscribe(channelID)
	if err != nil {
		return nil, fmt.Errorf("failed to subscribe to channel %d: %w", channelID, err)
	}

	// add the engine ID to the cache
	n.cache.Add(channelID)

	// create the conduit
	conduit := &Conduit{
		channelID: channelID,
		submit:    n.submit,
	}

	// register engine with provided engineID
	n.engines[channelID] = engine
	return conduit, nil
}

// Identity returns a map of all flow.Identifier to flow identity by querying the flow state
func (n *Network) Identity() (map[flow.Identifier]flow.Identity, error) {
	ids, err := n.state.Final().Identities()

	if err != nil {
		return nil, fmt.Errorf("could not get identities: %w", err)
	}
	identifierToID := make(map[flow.Identifier]flow.Identity)
	for _, id := range ids {
		identifierToID[id.NodeID] = *id
	}
	return identifierToID, nil
}

// Cleanup implements a callback to handle peers that have been dropped
// by the middleware layer.
func (n *Network) Cleanup(nodeID flow.Identifier) error {
	// drop the peer state using the ID we registered
	n.top.Down(nodeID)
	return nil
}

func (n *Network) Receive(nodeID flow.Identifier, msg interface{}) error {

	var err error
	switch m := msg.(type) {
	case *message.Message:
		err = n.processNetworkMessage(nodeID, m)
	default:
		err = fmt.Errorf("network received invalid message type (%T)", m)
	}
	if err != nil {
		err = fmt.Errorf("could not process message: %w", err)
	}
	return err
}

func (n *Network) processNetworkMessage(senderID flow.Identifier, message *message.Message) error {

	// Extract channel id and find the registered engine
	channelID := uint8(message.ChannelID)
	en, found := n.engines[channelID]
	if !found {
		n.logger.Debug().Str("sender", senderID.String()).
			Uint8("channel", channelID).
			Msg(" dropping message since no engine to receive it was found")
		return nil
	}

	// Convert message payload to a known message type
	decodedMessage, err := n.codec.Decode(message.Payload)
	if err != nil {
		return fmt.Errorf("could not decode event: %w", err)
	}

	// call the engine with the message payload
	err = en.Process(senderID, decodedMessage)
	if err != nil {
		n.logger.Error().Str("sender", senderID.String()).Uint8("channel", channelID).Err(err)
		return fmt.Errorf("failed to process message from %s: %w", senderID.String(), err)
	}
	return nil
}

// genNetworkMessage uses the codec to encode an event into a NetworkMessage
func (n *Network) genNetworkMessage(channelID uint8, event interface{}, targetIDs ...flow.Identifier) (*message.Message, error) {
	// encode the payload using the configured codec
	payload, err := n.codec.Encode(event)
	if err != nil {
		return nil, errors.Wrap(err, "could not encode event")
	}

	// use a hash with an engine-specific salt to get the payload hash
	sip := siphash.New([]byte("libp2ppacking" + fmt.Sprintf("%03d", channelID)))

	var emTargets [][]byte
	for _, t := range targetIDs {
		emTargets = append(emTargets, t[:])
	}

	// get origin ID (inplace slicing n.me.NodeID()[:] doesn't work)
	selfID := n.me.NodeID()
	originID := selfID[:]

	//cast event to a libp2p.Message
	message := &message.Message{
		ChannelID: uint32(channelID),
		EventID:   sip.Sum(payload),
		OriginID:  originID,
		TargetIDs: emTargets,
		Payload:   payload,
	}

	return message, nil
}

// submit method submits the given event for the given channel to the overlay layer
// for processing; it is used by engines through conduits.
func (n *Network) submit(channelID uint8, event interface{}, targetIDs ...flow.Identifier) error {
	// genNetworkMessage the event to get payload and event ID
	message, err := n.genNetworkMessage(channelID, event, targetIDs...)
	if err != nil {
		return errors.Wrap(err, "could not cast the event into network message")
	}

	// checks if the event is already in the cache
	ok := n.cache.Has(channelID, message.EventID)
	if ok {
		// returns nil and terminates sending the message since
		// the message already submitted to the network
		return nil
	}
	// storing event in the cache
	n.cache.Set(channelID, message.EventID, message)

	// TODO: debup the message here

	err = n.send(channelID, message, targetIDs...)
	if err != nil {
		return errors.Wrap(err, "could not gossip event")
	}

	return nil
}

// send sends the message to the set of target ids through the middleware
// send is the last method within the pipeline of message shipping in network layer
// once it is called, the message slips through the network layer towards the middleware
// If there is only one target NodeID, then a direct 1-1 connection is used by calling middleware.send
// Otherwise, middleware.Publish is used, which uses the PubSub method of communication.
// TODO: Move this decision making to the Middleware Issue#2246
func (n *Network) send(channelID uint8, msg *message.Message, nodeIDs ...flow.Identifier) error {
	var err error
	switch len(nodeIDs) {
	case 0:
		return fmt.Errorf("list of target node IDs empty")
	case 1:
		if nodeIDs[0] == n.me.NodeID() {
			// to avoid self dial by the underlay
			n.logger.Debug().Msg("self dial attempt")
			return nil
		}
		err = n.mw.Send(nodeIDs[0], msg)
	default:
		err = n.mw.Publish(strconv.Itoa(int(channelID)), msg)
	}
	if err != nil {
		err = fmt.Errorf("failed to send message to %s:%w", nodeIDs, err)
	}
	return err
}
