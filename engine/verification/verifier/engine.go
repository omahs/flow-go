package verifier

import (
	"fmt"

	"github.com/rs/zerolog"

	"github.com/dapperlabs/flow-go/crypto/hash"
	"github.com/dapperlabs/flow-go/engine"
	"github.com/dapperlabs/flow-go/engine/verification"
	"github.com/dapperlabs/flow-go/engine/verification/utils"
	chmodels "github.com/dapperlabs/flow-go/model/chunks"
	"github.com/dapperlabs/flow-go/model/flow"
	"github.com/dapperlabs/flow-go/model/flow/filter"
	"github.com/dapperlabs/flow-go/module"
	"github.com/dapperlabs/flow-go/network"
	"github.com/dapperlabs/flow-go/state/protocol"
	"github.com/dapperlabs/flow-go/utils/logging"
)

// Engine (verifier engine) verifies chunks, generates result approvals or raises challenges.
// as input it accepts verifiable chunks (chunk + all data needed) and perform verification by
// constructing a partial trie, executing transactions and check the final state commitment and
// other chunk meta data (e.g. tx count)
type Engine struct {
	unit    *engine.Unit               // used to control startup/shutdown
	log     zerolog.Logger             // used to log relevant actions
	metrics module.VerificationMetrics // used to capture the performance metrics
	conduit network.Conduit            // used to propagate result approvals
	me      module.Local               // used to access local node information
	state   protocol.State             // used to access the protocol state
	rah     hash.Hasher                // used as hasher to sign the result approvals
	chVerif module.ChunkVerifier       // used to verify chunks
}

// New creates and returns a new instance of a verifier engine.
func New(
	log zerolog.Logger,
	metrics module.VerificationMetrics,
	net module.Network,
	state protocol.State,
	me module.Local,
	chVerif module.ChunkVerifier,
) (*Engine, error) {

	e := &Engine{
		unit:    engine.NewUnit(),
		log:     log.With().Str("engine", "verifier").Logger(),
		metrics: metrics,
		state:   state,
		me:      me,
		chVerif: chVerif,
		rah:     utils.NewResultApprovalHasher(),
	}

	var err error
	e.conduit, err = net.Register(engine.ApprovalProvider, e)
	if err != nil {
		return nil, fmt.Errorf("could not register engine on approval provider channel: %w", err)
	}

	return e, nil
}

// Ready returns a channel that is closed when the verifier engine is ready.
func (e *Engine) Ready() <-chan struct{} {
	return e.unit.Ready()
}

// Done returns a channel that is closed when the verifier engine is done.
func (e *Engine) Done() <-chan struct{} {
	return e.unit.Done()
}

// SubmitLocal submits an event originating on the local node.
func (e *Engine) SubmitLocal(event interface{}) {
	e.Submit(e.me.NodeID(), event)
}

// Submit submits the given event from the node with the given origin ID
// for processing in a non-blocking manner. It returns instantly and logs
// a potential processing error internally when done.
func (e *Engine) Submit(originID flow.Identifier, event interface{}) {
	e.unit.Launch(func() {
		err := e.Process(originID, event)
		if err != nil {
			engine.LogError(e.log, err)
		}
	})
}

// ProcessLocal processes an event originating on the local node.
func (e *Engine) ProcessLocal(event interface{}) error {
	return e.Process(e.me.NodeID(), event)
}

// Process processes the given event from the node with the given origin ID in
// a blocking manner. It returns the potential processing error when done.
func (e *Engine) Process(originID flow.Identifier, event interface{}) error {
	return e.unit.Do(func() error {
		return e.process(originID, event)
	})
}

// process receives verifiable chunks, evaluate them and send them for chunk verifier
func (e *Engine) process(originID flow.Identifier, event interface{}) error {
	switch resource := event.(type) {
	case *verification.VerifiableChunkData:
		return e.verifyWithMetrics(originID, resource)
	default:
		return fmt.Errorf("invalid event type (%T)", event)
	}
}

// verify handles the core verification process. It accepts a verifiable chunk
// and all dependent resources, verifies the chunk, and emits a
// result approval if applicable.
//
// If any part of verification fails, an error is returned, indicating to the
// initiating engine that the verification must be re-tried.
func (e *Engine) verify(originID flow.Identifier, vc *verification.VerifiableChunkData) error {
	// log it first
	log := e.log.With().Timestamp().
		Hex("origin", logging.ID(originID)).
		Uint64("chunk_index", vc.Chunk.Index).
		Hex("execution_result_id", logging.Entity(vc.Result)).
		Logger()

	log.Info().Msg("verifiable chunk received by verifier engine")

	// only accept internal calls
	if originID != e.me.NodeID() {
		return fmt.Errorf("invalid remote origin for verify")
	}
	// extracts list of verifier nodes id
	//
	// TODO state extraction should be done based on block references
	consensusNodes, err := e.state.Final().
		Identities(filter.HasRole(flow.RoleConsensus))
	if err != nil {
		// TODO this error needs more advance handling after MVP
		return fmt.Errorf("could not load consensus node IDs: %w", err)
	}

	// extracts chunk ID
	ch, ok := vc.Result.Chunks.ByIndex(vc.Chunk.Index)
	if !ok {
		return engine.NewInvalidInputErrorf("chunk out of range requested: %v", vc.Chunk.Index)
	}
	log.With().Hex("chunk_id", logging.Entity(ch)).Logger()

	// execute the assigned chunk
	chFault, err := e.chVerif.Verify(vc)
	// Any err means that something went wrong when verify the chunk
	// the outcome of the verification is captured inside the chFault and not the err
	if err != nil {
		return err
	}

	// if any fault found with the chunk
	if chFault != nil {
		switch chFault.(type) {
		case *chmodels.CFMissingRegisterTouch:
		case *chmodels.CFNonMatchingFinalState:
			// TODO raise challenge
			e.log.Warn().Msg(chFault.String())
		case *chmodels.CFInvalidVerifiableChunk:
			// TODO raise challenge
			e.log.Error().Msg(chFault.String())
		default:
			return engine.NewInvalidInputErrorf("unknown type of chunk fault is recieved (type: %T) : %v", chFault, chFault.String())
		}
		// don't do anything else, but skip generating result approvals
		return nil
	}

	// Generate result approval
	approval, err := e.GenerateResultApproval(vc.Chunk.Index, vc.Result.ID(), vc.Header.ID())
	if err != nil {
		return fmt.Errorf("couldn't generate a result approval: %w", err)
	}

	// broadcast result approval to the consensus nodes
	err = e.conduit.Submit(approval, consensusNodes.NodeIDs()...)
	if err != nil {
		// TODO this error needs more advance handling after MVP
		return fmt.Errorf("could not submit result approval: %w", err)
	}
	log.Info().Msg("result approval submitted")
	// increases number of sent result approvals for sake of metrics
	e.metrics.OnResultApproval()

	return nil
}

// GenerateResultApproval generates result approval for specific chunk of a exec receipt
func (e *Engine) GenerateResultApproval(chunkIndex uint64, execResultID flow.Identifier, blockID flow.Identifier) (*flow.ResultApproval, error) {

	// attestation
	atst := flow.Attestation{
		BlockID:           blockID,
		ExecutionResultID: execResultID,
		ChunkIndex:        chunkIndex,
	}

	// generates a signature over the attestation part of approval
	atstID := atst.ID()
	atstSign, err := e.me.Sign(atstID[:], e.rah)
	if err != nil {
		return nil, fmt.Errorf("could not sign attestation: %w", err)
	}

	// result approval body
	body := flow.ResultApprovalBody{
		Attestation:          atst,
		ApproverID:           e.me.NodeID(),
		AttestationSignature: atstSign,
		Spock:                nil,
	}

	// generates a signature over result approval body
	bodyID := body.ID()
	bodySign, err := e.me.Sign(bodyID[:], e.rah)
	if err != nil {
		return nil, fmt.Errorf("could not sign result approval body: %w", err)
	}

	return &flow.ResultApproval{
		Body:              body,
		VerifierSignature: bodySign,
	}, nil
}

// verifyWithMetrics acts as a wrapper around the verify method that captures its performance-related metrics
func (e *Engine) verifyWithMetrics(originID flow.Identifier, ch *verification.VerifiableChunkData) error {
	// increments number of received verifiable chunks
	// for sake of metrics
	e.metrics.OnVerifiableChunkReceived()

	// starts verification performance metrics trackers
	if ch.ChunkDataPack != nil {
		e.metrics.OnChunkVerificationStarted(ch.ChunkDataPack.ChunkID)
	}
	// starts verification of chunk
	err := e.verify(originID, ch)
	// closes verification performance metrics trackers
	if ch.ChunkDataPack != nil {
		e.metrics.OnChunkVerificationFinished(ch.ChunkDataPack.ChunkID)
	}
	return err
}
