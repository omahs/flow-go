package trie_test

import (
	"encoding/binary"
	"fmt"
	"testing"

	"github.com/onflow/flow-go/ledger"
	"github.com/onflow/flow-go/ledger/common/hash"
	"github.com/onflow/flow-go/ledger/common/testutils"
	"github.com/onflow/flow-go/ledger/complete/mtrie/node"
	"github.com/onflow/flow-go/ledger/complete/mtrie/trie"
	"github.com/onflow/flow-go/utils/unittest"
	"github.com/stretchr/testify/require"
)

func Test_PathCreation(t *testing.T) {
	paths := []ledger.Path{
		testutils.PathByUint8(1 << 7),   // 1000...
		testutils.PathByUint8(1 << 6),   // 1000...
		testutils.PathByUint16(1 << 15), // 1000...
		testutils.PathByUint16(1 << 14), // 1000...
	}

	// emptyTrie := trie.NewEmptyMTrie()
	for _, path := range paths {
		fmt.Println(path)
	}
	// payload := testutils.LightPayload(11, 12345)
	var p ledger.Path
	binary.BigEndian.PutUint32(p[28:32], 1)
	fmt.Println(p)
}

// a full trie with 4 payloads at bottom
func Test_FullTrie4_Update(t *testing.T) {
	// prepare paths for a trie with only 2 levels, which has 4 payloads in total
	paths := []ledger.Path{
		testutils.PathByUint8(0),           // 00000000...
		testutils.PathByUint8(1 << 6),      // 01000000...
		testutils.PathByUint8(1 << 7),      // 10000000...
		testutils.PathByUint8(1<<7 + 1<<6), // 11000000...
	}

	// prepare payloads, the value doesn't matter
	payloads := make([]ledger.Payload, 0, len(paths))
	for i := range paths {
		payloads = append(payloads, *payloadFromInt16(uint16(i)))
	}

	t0 := trie.NewEmptyMTrie()
	payloadStorage := unittest.CreateMockPayloadStore()

	// create a full trie by adding payloads at all 4 paths to an empty trie
	t1, depth, err := trie.NewTrieWithUpdatedRegisters(t0, paths, payloads, true, payloadStorage)
	require.NoError(t, err)
	require.Equal(t, uint16(2), depth)

	// create a updated trie that is updated at one path,
	// so prepare the updated path first
	paths2 := []ledger.Path{
		paths[2], // 10000000...
	}
	payloads2 := []ledger.Payload{
		*payloadFromInt16(uint16(100)),
	}
	// create the updated trie at the given path
	t2, _, err := trie.NewTrieWithUpdatedRegisters(t1, paths2, payloads2, true, payloadStorage)
	require.NoError(t, err)

	// empty trie should have nothing stored at the path
	from0, err := t0.ReadSinglePayload(paths2[0], payloadStorage)
	require.NoError(t, err)
	require.True(t, from0.IsEmpty())

	// the full trie should have original payload at the path
	from1, err := t1.ReadSinglePayload(paths2[0], payloadStorage)
	require.NoError(t, err)
	require.True(t, from1.Equals(&payloads[2]))

	// the updated trie should have updated payload at the path
	from2, err := t2.ReadSinglePayload(paths2[0], payloadStorage)
	require.NoError(t, err)
	require.True(t, from2.Equals(&payloads2[0]))

}

func payloadFromInt16(i uint16) *ledger.Payload {
	return testutils.LightPayload(i, i)
}

func getAllPathsAtLevel2() []ledger.Path {
	return []ledger.Path{
		testutils.PathByUint8(0),           // 00000000...
		testutils.PathByUint8(1 << 6),      // 01000000...
		testutils.PathByUint8(1 << 7),      // 10000000...
		testutils.PathByUint8(1<<7 + 1<<6), // 11000000...
	}
}

// adding payloads to the trie until full
func Test_AddingPayloadUntilFull(t *testing.T) {
	// prepare paths for a trie with only 2 levels, which has 4 payloads in total
	paths := getAllPathsAtLevel2()

	// prepare payloads, the value doesn't matter
	payloads := make([]ledger.Payload, 0, len(paths))
	for i := range paths {
		payloads = append(payloads, *payloadFromInt16(uint16(i)))
	}

	t0 := trie.NewEmptyMTrie()
	payloadStorage := unittest.CreateMockPayloadStore()
	fmt.Println(t0)

	fmt.Println("=== adding 1 payload at 00000000")
	t1, depth, err := trie.NewTrieWithUpdatedRegisters(t0, paths[0:1], payloads[0:1], true, payloadStorage)
	require.NoError(t, err)
	require.Equal(t, uint16(0), depth)
	fmt.Println(t1)

	fmt.Println("=== adding 1 payload at 01000000")
	t2, depth, err := trie.NewTrieWithUpdatedRegisters(t1, paths[1:2], payloads[1:2], true, payloadStorage)
	require.NoError(t, err)
	require.Equal(t, uint16(2), depth)
	fmt.Println(t2)

	emptyPayloads := []ledger.Payload{
		*ledger.EmptyPayload(),
	}

	fmt.Println("=== removing 1 payload at 01000000")
	t3, depth, err := trie.NewTrieWithUpdatedRegisters(t2, paths[1:2], emptyPayloads, true, payloadStorage)
	require.NoError(t, err)
	require.Equal(t, uint16(2), depth)
	fmt.Println(t3)
	require.True(t, t1.Equals(t3))

	fmt.Println("=== removing 1 payload at 00000000")
	t4, depth, err := trie.NewTrieWithUpdatedRegisters(t3, paths[0:1], emptyPayloads, true, payloadStorage)
	require.NoError(t, err)
	require.Equal(t, uint16(0), depth)
	fmt.Println(t4)
	require.True(t, t0.Equals(t4))
}

type pathAndPayload struct {
	Paths    []ledger.Path
	Payloads []ledger.Payload
}

func Test_Lookup_AddUntilFull(t *testing.T) {
	// prepare paths for a trie with only 2 levels, which has 4 payloads in total
	paths := getAllPathsAtLevel2()
	payloads := testutils.RandomPayloads(4, 10, 20)

	groups := getPermutationPathsPayloads(paths, payloads)
	for _, group := range groups {
		payloadStorage := unittest.CreateMockPayloadStore()
		updated, _, err := trie.NewTrieWithUpdatedRegisters(
			trie.NewEmptyMTrie(), group.Paths, group.Payloads, true, payloadStorage)
		require.NoError(t, err)

		nonExistingPaths := findNonExistingPaths(paths, group.Paths)

		for i, path := range group.Paths {
			payload := group.Payloads[i]
			foundPayload, err := updated.ReadSinglePayload(path, payloadStorage)
			require.NoError(t, err)
			require.True(t, foundPayload.Equals(&payload))
		}

		for _, nonExistingPath := range nonExistingPaths {
			shouldBeEmpty, err := updated.ReadSinglePayload(nonExistingPath, payloadStorage)
			require.NoError(t, err)
			require.True(t, shouldBeEmpty.IsEmpty())
		}
	}
}

func getPermutationPathsPayloads(paths []ledger.Path, payloads []*ledger.Payload) []*pathAndPayload {
	permutation := make([]*pathAndPayload, 0)
	for i, path := range paths {
		payload := payloads[i]

		count := len(permutation)
		for i := 0; i < count; i++ {
			group := permutation[i]
			permutation = append(permutation, &pathAndPayload{
				Paths:    append(group.Paths, path),
				Payloads: append(group.Payloads, *payload),
			})
		}

		permutation = append(permutation, &pathAndPayload{
			Paths:    []ledger.Path{path},
			Payloads: []ledger.Payload{*payload},
		})
	}
	return permutation
}

func findNonExistingPaths(all []ledger.Path, paths []ledger.Path) []ledger.Path {
	lookup := make(map[ledger.Path]struct{})
	for _, path := range paths {
		lookup[path] = struct{}{}
	}

	missing := make([]ledger.Path, 0, len(all))
	for _, path := range all {
		_, ok := lookup[path]
		if !ok {
			missing = append(missing, path)
		}
	}
	return missing
}

func nEmptyPayloads(n int) []ledger.Payload {
	emptyPayloads := make([]ledger.Payload, 0, n)
	for i := 0; i < n; i++ {
		emptyPayloads = append(emptyPayloads, *ledger.EmptyPayload())
	}
	return emptyPayloads
}

func toPayloads(pointers []*ledger.Payload) []ledger.Payload {
	payloads := make([]ledger.Payload, 0, len(pointers))
	for _, p := range pointers {
		payloads = append(payloads, *p)
	}
	return payloads
}

func Test_Lookup_RemoveUntilEmpty(t *testing.T) {
	// prepare paths for a trie with only 2 levels, which has 4 payloads in total
	paths := getAllPathsAtLevel2()
	randoms := testutils.RandomPayloads(4, 10, 20)
	payloads := toPayloads(randoms)

	groups := getPermutationPathsPayloads(paths, randoms)

	for _, group := range groups {
		payloadStorage := unittest.CreateMockPayloadStore()

		// create a full trie with all 4 payloads added
		fullTrie, _, err := trie.NewTrieWithUpdatedRegisters(
			trie.NewEmptyMTrie(), paths, payloads, true, payloadStorage)
		require.NoError(t, err)

		// remove some payloads at the given paths
		updated, _, err := trie.NewTrieWithUpdatedRegisters(
			fullTrie, group.Paths, nEmptyPayloads(len(group.Paths)), true, payloadStorage)
		require.NoError(t, err)

		// the following 2 steps verifies that for the removed paths, the payload no longer exists,
		// for not removed paths, the payload still exists
		// 1. build a lookup
		lookup := make(map[ledger.Path]*ledger.Payload)
		for i, path := range group.Paths {
			payload := group.Payloads[i]
			lookup[path] = &payload
		}

		// 2. use the lookup to check
		for _, path := range group.Paths {
			payloadRead, err := updated.ReadSinglePayload(path, payloadStorage)
			require.NoError(t, err)

			deletedPayload, isDeleted := lookup[path]
			if isDeleted {
				require.True(t, payloadRead.IsEmpty())
			} else {
				require.True(t, payloadRead.Equals(deletedPayload))
			}
		}

		// verify the same path can be deleted again:
		// deleting all paths again regardless the path has payload or not, verify no path has payload
		allDeleted, _, err := trie.NewTrieWithUpdatedRegisters(
			updated, paths, nEmptyPayloads(len(paths)), true, payloadStorage)

		for _, path := range paths {
			payloadRead, err := allDeleted.ReadSinglePayload(path, payloadStorage)
			require.NoError(t, err)

			require.True(t, payloadRead.IsEmpty())
		}
	}
}

func Test_RootHash_AddUntilFull(t *testing.T) {
}

func Test_RootHash_RemoveUntilEmpty(t *testing.T) {
}

func Test_Prove_AddUntilFull(t *testing.T) {
}

func Test_Prove_RemoveUntilEmpty(t *testing.T) {
}

func computeRootHash(paths []ledger.Path, payloads []*ledger.Payload) hash.Hash {
	lookup := make(map[ledger.Path]*ledger.Payload)

	for i, path := range paths {
		lookup[path] = payloads[i]
	}
	path1, path2, path3, path4 := paths[0], paths[1], paths[2], paths[3]

	//     h7
	//   /    \
	//  h5    h6
	// /  \  /  \
	// h1 h2 h3 h4
	h1 := getLeafHash(path1, lookup)
	h2 := getLeafHash(path2, lookup)
	h3 := getLeafHash(path3, lookup)
	h4 := getLeafHash(path4, lookup)
	h5 := hash.HashInterNode(h1, h2)
	h6 := hash.HashInterNode(h3, h4)
	h7 := hash.HashInterNode(h5, h6)
	return h7
}

func getLeafHash(path ledger.Path, lookup map[ledger.Path]*ledger.Payload) hash.Hash {
	payload, ok := lookup[path]
	height := 256 - 2
	if !ok {
		return ledger.GetDefaultHashForHeight(height)
	}

	leaf := node.NewLeaf(path, payload, height)
	return leaf.Hash()
}
