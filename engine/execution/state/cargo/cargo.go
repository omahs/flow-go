package cargo

import (
	"time"

	"go.uber.org/atomic"

	"github.com/onflow/flow-go/fvm/state"
	"github.com/onflow/flow-go/model/flow"

	"github.com/onflow/flow-go/engine/execution/state/cargo/payload"
	"github.com/onflow/flow-go/engine/execution/state/cargo/queue"
	"github.com/onflow/flow-go/engine/execution/state/cargo/storage"
)

// Cargo is responsible for synchronization
// between a block execution and block finalization events
// block execution results and finalization are two concurrent
// streams and requires synchronization and buffering of events
// in case one of them is ahead of the other one.
// payloadStore accepts block execution events in a fork-aware way and expect block finalization
// signals to commit the changes into persistant storage.
// blockqueue acts as a fixed-size buffer to hold on to the unprocessed block finaliziation events
// both blockqueue and payloadstore has some internal validation to prevent
// out of order events and support concurency.
type Cargo struct {
	blockQueue   *queue.FinalizedBlockQueue
	payloadStore *payload.PayloadStore

	syncFrequency  time.Duration
	syncInProgress *atomic.Bool
	stopSync       chan struct{}
	syncStopped    chan struct{}
	syncError      chan error
}

// NewCargo constructs a new Cargo
// if syncFrequency is set to zero, it won't run and requires manual triggering
func NewCargo(
	storage storage.Storage,
	blockQueueCapacity int,
	genesis *flow.Header,
	syncFrequency time.Duration,
) (*Cargo, error) {
	// TODO load genesis from storage
	payloadStore, err := payload.NewPayloadStore(storage)
	if err != nil {
		return nil, err
	}
	c := &Cargo{
		blockQueue:     queue.NewFinalizedBlockQueue(blockQueueCapacity, genesis),
		payloadStore:   payloadStore,
		syncFrequency:  syncFrequency,
		syncInProgress: atomic.NewBool(false),
		stopSync:       make(chan struct{}, 1),
		syncStopped:    make(chan struct{}, 1),
		syncError:      make(chan error, 1),
	}
	c.Start()
	return c, nil
}

// Reader returns a reader for the execution
func (c *Cargo) Reader(header *flow.Header) state.StorageSnapshot {
	return c.payloadStore.Reader(header)
}

// BlockFinalized is called every time a block is executed
func (c *Cargo) BlockExecuted(header *flow.Header, updates map[flow.RegisterID]flow.RegisterValue) error {
	// add updates to the payload store
	return c.payloadStore.BlockExecuted(header, updates)
}

// BlockFinalized is called every time a new block is finalized
func (c *Cargo) BlockFinalized(new *flow.Header) error {
	// enqueue the finalized header
	// TODO deal with errors (e.g. if already processed move on)
	return c.blockQueue.Enqueue(new)
}

// TrySync tries to sync the block finalization queue with the payload store
func (c *Cargo) TrySync() error {
	// if sync is in progress skip
	if c.syncInProgress.CompareAndSwap(false, true) {
		defer c.syncInProgress.Store(false)
		// we take a peak at the oldest block that has been finalized and not processed yet
		// and see if its commitable by payload storage (if results are available for that blockID),
		// if commitable (return true by payload store), we deuque the block and continue
		// doing the same for more blocks until payloadstore returns false, which mean it doesn't
		// have results for that block yet and we need to hold on until next trigger.
		for c.blockQueue.HasHeaders() {
			blockID, header := c.blockQueue.Peak()
			found, err := c.payloadStore.BlockFinalized(blockID, header)
			if err != nil {
				return err
			}
			if !found {
				break
			}
			c.blockQueue.Dequeue()
		}

	}
	return nil
}

// Start starts the cargo sync ticker
func (c *Cargo) Start() {
	if c.syncFrequency > 0 {
		ticker := time.NewTicker(c.syncFrequency)
		go func() {
			for {
				select {
				case <-c.stopSync:
					ticker.Stop()
					c.syncStopped <- struct{}{}
					return
				case <-ticker.C:
					err := c.TrySync()
					if err != nil {
						c.syncError <- err
						return
					}
				}
			}
		}()
		return
	}
	c.syncStopped <- struct{}{}
}

func (c *Cargo) Ready() <-chan struct{} {
	// TODO initate the storage here
	ready := make(chan struct{})
	defer close(ready)
	return ready
}

func (c *Cargo) Done() <-chan struct{} {
	done := make(chan struct{})
	defer close(done)

	// wait for sync to stop
	c.stopSync <- struct{}{}
	<-c.syncStopped

	// TODO check the error line
	// <- c.syncError

	close(c.stopSync)
	close(c.syncStopped)
	close(c.syncError)

	return done
}
