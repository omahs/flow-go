package extractpayloads

import (
	"crypto/rand"
	"path/filepath"
	"strings"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"

	"github.com/onflow/cadence/common"

	"github.com/onflow/flow-go/cmd/util/ledger/util"
	"github.com/onflow/flow-go/ledger"
	"github.com/onflow/flow-go/utils/unittest"
)

type keyPair struct {
	key   ledger.Key
	value ledger.Value
}

func TestExtractPayloads(t *testing.T) {

	t.Run("some payloads", func(t *testing.T) {

		unittest.RunWithTempDir(t, func(datadir string) {

			inputFile := filepath.Join(datadir, "input.payload")
			outputFile := filepath.Join(datadir, "output.payload")

			size := 10

			// Generate some data
			keysValues := make(map[string]keyPair)
			var payloads []*ledger.Payload

			for i := 0; i < size; i++ {
				keys, values := getSampleKeyValues(i)

				for j, key := range keys {
					keysValues[key.String()] = keyPair{
						key:   key,
						value: values[j],
					}

					payloads = append(payloads, ledger.NewPayload(key, values[j]))
				}
			}

			numOfPayloadWritten, err := util.CreatePayloadFile(zerolog.Nop(), inputFile, payloads, nil, false)
			require.NoError(t, err)
			require.Equal(t, len(payloads), numOfPayloadWritten)

			const selectedAddressCount = 10
			selectedAddresses := make(map[string]struct{})
			selectedKeysValues := make(map[string]keyPair)
			for k, kv := range keysValues {
				owner := kv.key.KeyParts[0].Value
				if len(owner) != common.AddressLength {
					continue
				}

				address, err := common.BytesToAddress(owner)
				require.NoError(t, err)

				if len(selectedAddresses) < selectedAddressCount {
					selectedAddresses[address.Hex()] = struct{}{}
				}

				if _, exist := selectedAddresses[address.Hex()]; exist {
					selectedKeysValues[k] = kv
				}
			}

			addresses := make([]string, 0, len(selectedAddresses))
			for address := range selectedAddresses {
				addresses = append(addresses, address)
			}

			// Export selected payloads
			Cmd.SetArgs([]string{
				"--input-filename", inputFile,
				"--output-filename", outputFile,
				"--addresses", strings.Join(addresses, ","),
			})

			err = Cmd.Execute()
			require.NoError(t, err)

			// Verify exported payloads.
			partialState, payloadsFromFile, err := util.ReadPayloadFile(zerolog.Nop(), outputFile)
			require.NoError(t, err)
			require.True(t, partialState)

			nonGlobalPayloads := make([]*ledger.Payload, 0, len(selectedKeysValues))
			for _, payloadFromFile := range payloadsFromFile {
				key, err := payloadFromFile.Key()
				require.NoError(t, err)

				owner := key.KeyParts[0].Value
				if len(owner) > 0 {
					nonGlobalPayloads = append(nonGlobalPayloads, payloadFromFile)
				}
			}

			require.Equal(t, len(selectedKeysValues), len(nonGlobalPayloads))

			for _, payloadFromFile := range nonGlobalPayloads {
				k, err := payloadFromFile.Key()
				require.NoError(t, err)

				kv, exist := selectedKeysValues[k.String()]
				require.True(t, exist)
				require.Equal(t, kv.value, payloadFromFile.Value())
			}
		})
	})

	t.Run("empty address", func(t *testing.T) {

		unittest.RunWithTempDir(t, func(datadir string) {

			inputFile := filepath.Join(datadir, "input.payload")
			outputFile := filepath.Join(datadir, "output.payload")

			size := 10

			// Generate some data
			keysValues := make(map[string]keyPair)
			var payloads []*ledger.Payload

			for i := 0; i < size; i++ {
				keys, values := getSampleKeyValues(i)

				for j, key := range keys {
					keysValues[key.String()] = keyPair{
						key:   key,
						value: values[j],
					}

					payloads = append(payloads, ledger.NewPayload(key, values[j]))
				}
			}

			numOfPayloadWritten, err := util.CreatePayloadFile(zerolog.Nop(), inputFile, payloads, nil, false)
			require.NoError(t, err)
			require.Equal(t, len(payloads), numOfPayloadWritten)

			// Export selected payloads
			Cmd.SetArgs([]string{
				"--input-filename", inputFile,
				"--output-filename", outputFile,
				"--addresses", ",",
			})

			err = Cmd.Execute()
			require.NoError(t, err)

			// Verify exported payloads.
			partialState, payloadsFromFile, err := util.ReadPayloadFile(zerolog.Nop(), outputFile)
			require.NoError(t, err)
			require.True(t, partialState)

			var nonGlobalPayloads []*ledger.Payload
			for _, payloadFromFile := range payloadsFromFile {
				key, err := payloadFromFile.Key()
				require.NoError(t, err)

				owner := key.KeyParts[0].Value
				if len(owner) > 0 {
					nonGlobalPayloads = append(nonGlobalPayloads, payloadFromFile)
				}
			}

			require.Equal(t, 0, len(nonGlobalPayloads))

		})
	})
}

func getSampleKeyValues(i int) ([]ledger.Key, []ledger.Value) {
	switch i {
	case 0:
		return []ledger.Key{getKey("", "uuid"), getKey("", "account_address_state")},
			[]ledger.Value{[]byte{'1'}, []byte{'A'}}
	case 1:
		return []ledger.Key{getKey("ADDRESS", "public_key_count"),
				getKey("ADDRESS", "public_key_0"),
				getKey("ADDRESS", "exists"),
				getKey("ADDRESS", "storage_used")},
			[]ledger.Value{[]byte{1}, []byte("PUBLICKEYXYZ"), []byte{1}, []byte{100}}
	case 2:
		// TODO change the contract_names to CBOR encoding
		return []ledger.Key{getKey("ADDRESS", "contract_names"), getKey("ADDRESS", "code.mycontract")},
			[]ledger.Value{[]byte("mycontract"), []byte("CONTRACT Content")}
	default:
		keys := make([]ledger.Key, 0)
		values := make([]ledger.Value, 0)
		for j := 0; j < 10; j++ {
			// address := make([]byte, 32)
			address := make([]byte, 8)
			_, err := rand.Read(address)
			if err != nil {
				panic(err)
			}
			keys = append(keys, getKey(string(address), "test"))
			values = append(values, getRandomCadenceValue())
		}
		return keys, values
	}
}

func getKey(owner, key string) ledger.Key {
	return ledger.Key{KeyParts: []ledger.KeyPart{
		{Type: uint16(0), Value: []byte(owner)},
		{Type: uint16(2), Value: []byte(key)},
	},
	}
}

func getRandomCadenceValue() ledger.Value {

	randomPart := make([]byte, 10)
	_, err := rand.Read(randomPart)
	if err != nil {
		panic(err)
	}
	valueBytes := []byte{
		// magic prefix
		0x0, 0xca, 0xde, 0x0, 0x4,
		// tag
		0xd8, 132,
		// array, 5 items follow
		0x85,

		// tag
		0xd8, 193,
		// UTF-8 string, length 4
		0x64,
		// t, e, s, t
		0x74, 0x65, 0x73, 0x74,

		// nil
		0xf6,

		// positive integer 1
		0x1,

		// array, 0 items follow
		0x80,

		// UTF-8 string, length 10
		0x6a,
		0x54, 0x65, 0x73, 0x74, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	}

	valueBytes = append(valueBytes, randomPart...)
	return ledger.Value(valueBytes)
}
