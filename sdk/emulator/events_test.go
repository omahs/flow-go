package emulator_test

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/dapperlabs/flow-go/pkg/constants"
	"github.com/dapperlabs/flow-go/pkg/crypto"
	"github.com/dapperlabs/flow-go/pkg/types"
	"github.com/dapperlabs/flow-go/sdk/emulator"
	"github.com/dapperlabs/flow-go/sdk/emulator/execution"
)

func TestEventEmitted(t *testing.T) {
	t.Run("EmittedFromTransaction", func(t *testing.T) {
		events := make([]types.Event, 0)

		b := emulator.NewEmulatedBlockchain(emulator.EmulatedBlockchainOptions{
			OnEventEmitted: func(event types.Event, blockNumber uint64, txHash crypto.Hash) {
				events = append(events, event)
			},
		})

		script := []byte(`
			event MyEvent(x: Int, y: Int)
			
			fun main() {
			  emit MyEvent(x: 1, y: 2)
			}
		`)

		tx := &types.Transaction{
			Script:             script,
			ReferenceBlockHash: nil,
			Nonce:              getNonce(),
			ComputeLimit:       10,
			PayerAccount:       b.RootAccountAddress(),
		}

		tx.AddSignature(b.RootAccountAddress(), b.RootKey())

		err := b.SubmitTransaction(tx)
		assert.Nil(t, err)

		require.Len(t, events, 1)

		expectedID := fmt.Sprintf("tx.%s.MyEvent", tx.Hash().Hex())

		assert.Equal(t, expectedID, events[0].ID)
		assert.Equal(t, big.NewInt(1), events[0].Values["x"])
		assert.Equal(t, big.NewInt(2), events[0].Values["y"])
	})

	t.Run("EmittedFromScript", func(t *testing.T) {
		events := make([]types.Event, 0)

		b := emulator.NewEmulatedBlockchain(emulator.EmulatedBlockchainOptions{
			OnEventEmitted: func(event types.Event, blockNumber uint64, txHash crypto.Hash) {
				events = append(events, event)
			},
		})

		script := []byte(`
			event MyEvent(x: Int, y: Int)
			
			fun main() {
			  emit MyEvent(x: 1, y: 2)
			}
		`)

		_, err := b.CallScript(script)
		assert.Nil(t, err)

		require.Len(t, events, 1)

		expectedID := fmt.Sprintf("script.%s.MyEvent", execution.ScriptHash(script).Hex())

		assert.Equal(t, expectedID, events[0].ID)
		assert.Equal(t, big.NewInt(1), events[0].Values["x"])
		assert.Equal(t, big.NewInt(2), events[0].Values["y"])
	})

	t.Run("EmittedFromAccount", func(t *testing.T) {
		events := make([]types.Event, 0)

		b := emulator.NewEmulatedBlockchain(emulator.EmulatedBlockchainOptions{
			OnEventEmitted: func(event types.Event, blockNumber uint64, txHash crypto.Hash) {
				events = append(events, event)
			},
		})

		accountScript := []byte(`
			event MyEvent(x: Int, y: Int)
		`)

		publicKey := b.RootKey().PublicKey()

		accountKey := types.AccountKey{
			PublicKey: publicKey,
			SignAlgo:  crypto.ECDSA_P256,
			HashAlgo:  crypto.SHA3_256,
			Weight:    constants.AccountKeyWeightThreshold,
		}

		address, err := b.CreateAccount([]types.AccountKey{accountKey}, accountScript, getNonce())
		assert.Nil(t, err)

		script := []byte(fmt.Sprintf(`
			import 0x%s
			
			fun main() {
				emit MyEvent(x: 1, y: 2)
			}
		`, address.Hex()))

		tx := &types.Transaction{
			Script:             script,
			ReferenceBlockHash: nil,
			Nonce:              getNonce(),
			ComputeLimit:       10,
			PayerAccount:       b.RootAccountAddress(),
		}

		tx.AddSignature(b.RootAccountAddress(), b.RootKey())

		err = b.SubmitTransaction(tx)
		assert.Nil(t, err)

		require.Len(t, events, 2)

		// first event is AccountCreated event
		expectedEvent := events[1]

		expectedID := fmt.Sprintf("account.%s.MyEvent", address.Hex())

		assert.Equal(t, expectedID, expectedEvent.ID)
		assert.Equal(t, big.NewInt(1), expectedEvent.Values["x"])
		assert.Equal(t, big.NewInt(2), expectedEvent.Values["y"])
	})
}
