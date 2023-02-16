package distributor_test

import (
	"context"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/onflow/flow-go/module/irrecoverable"
	"github.com/onflow/flow-go/network/p2p"
	"github.com/onflow/flow-go/network/p2p/distributor"
	mockp2p "github.com/onflow/flow-go/network/p2p/mock"
	p2ptest "github.com/onflow/flow-go/network/p2p/test"
	"github.com/onflow/flow-go/utils/unittest"
)

// TestGossipSubInspectorNotification tests the gossip sub inspector notification by adding two consumers to the
// notification distributor component and sending a random set of notifications to the notification component. The test
// verifies that the consumers receive the notifications.
func TestGossipSubInspectorNotification(t *testing.T) {
	g := distributor.DefaultGossipSubInspectorNotification(unittest.Logger())

	c1 := mockp2p.NewGossipSubRpcInspectorConsumer(t)
	c2 := mockp2p.NewGossipSubRpcInspectorConsumer(t)

	g.AddConsumer(c1)
	g.AddConsumer(c2)

	tt := invalidControlMessageNotificationListFixture(t, 100)

	c1Done := sync.WaitGroup{}
	c1Done.Add(len(tt))
	c1Seen := unittest.NewProtectedMap[peer.ID, struct{}]()
	c1.On("OnInvalidControlMessage", mock.Anything, mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		peerID, ok := args.Get(0).(peer.ID)
		require.True(t, ok)

		msgType, ok := args.Get(1).(p2p.ControlMessageType)
		require.True(t, ok)

		count, ok := args.Get(2).(uint64)
		require.True(t, ok)

		require.Contains(t, tt, p2p.InvalidControlMessageNotification{
			PeerID:  peerID,
			MsgType: msgType,
			Count:   count,
		})

		// ensure consumer see each peer once
		require.False(t, c1Seen.Has(peerID))
		c1Seen.Add(peerID, struct{}{})

		c1Done.Done()
	}).Return()

	c2Done := sync.WaitGroup{}
	c2Done.Add(len(tt))
	c2Seen := unittest.NewProtectedMap[peer.ID, struct{}]()
	c2.On("OnInvalidControlMessage", mock.Anything, mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		peerID, ok := args.Get(0).(peer.ID)
		require.True(t, ok)

		msgType, ok := args.Get(1).(p2p.ControlMessageType)
		require.True(t, ok)

		count, ok := args.Get(2).(uint64)
		require.True(t, ok)

		require.Contains(t, tt, p2p.InvalidControlMessageNotification{
			PeerID:  peerID,
			MsgType: msgType,
			Count:   count,
		})
		// ensure consumer see each peer once
		require.False(t, c2Seen.Has(peerID))
		c2Seen.Add(peerID, struct{}{})

		c2Done.Done()
	}).Return()

	cancelCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx, _ := irrecoverable.WithSignaler(cancelCtx)
	g.Start(ctx)

	unittest.RequireCloseBefore(t, g.Ready(), 100*time.Millisecond, "could not start distributor")

	for i := 0; i < len(tt); i++ {
		go func(i int) {
			g.OnInvalidControlMessage(tt[i].PeerID, tt[i].MsgType, tt[i].Count)
		}(i)
	}

	unittest.RequireReturnsBefore(t, c1Done.Wait, 1*time.Second, "events are not received by consumer 1")
	unittest.RequireReturnsBefore(t, c2Done.Wait, 1*time.Second, "events are not received by consumer 2")
	cancel()
	unittest.RequireCloseBefore(t, g.Done(), 100*time.Millisecond, "could not stop handler")
}

func invalidControlMessageNotificationListFixture(t *testing.T, n int) []p2p.InvalidControlMessageNotification {
	list := make([]p2p.InvalidControlMessageNotification, n)
	for i := 0; i < n; i++ {
		list[i] = invalidControlMessageNotificationFixture(t)
	}
	return list
}

func invalidControlMessageNotificationFixture(t *testing.T) p2p.InvalidControlMessageNotification {
	return p2p.InvalidControlMessageNotification{
		PeerID:  p2ptest.PeerIdFixture(t),
		MsgType: []p2p.ControlMessageType{p2p.CtrlMsgGraft, p2p.CtrlMsgPrune, p2p.CtrlMsgIHave, p2p.CtrlMsgIWant}[rand.Intn(4)],
		Count:   rand.Uint64(),
	}
}
