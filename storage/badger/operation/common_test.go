// (c) 2019 Dapper Labs - ALL RIGHTS RESERVED

package operation

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/dgraph-io/badger/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Entity struct {
	ID uint64
}

func TestInsertValid(t *testing.T) {

	dir := filepath.Join(os.TempDir(), fmt.Sprintf("flow-test-db-%d", rand.Uint64()))
	db, err := badger.Open(badger.DefaultOptions(dir).WithLogger(nil))
	require.Nil(t, err)

	e := Entity{ID: 1337}
	key := []byte{0x01, 0x02, 0x03}
	val := []byte(`{"ID":1337}`)

	err = db.Update(insertNew(key, e))
	require.Nil(t, err)

	var act []byte
	_ = db.View(func(tx *badger.Txn) error {
		item, err := tx.Get(key)
		require.Nil(t, err)
		act, err = item.ValueCopy(nil)
		require.Nil(t, err)
		return nil
	})

	assert.Equal(t, act, val)
}

func TestInsertDuplicate(t *testing.T) {
	// TODO
}

func TestUpdateValid(t *testing.T) {

	dir := filepath.Join(os.TempDir(), fmt.Sprintf("flow-test-db-%d", rand.Uint64()))
	db, err := badger.Open(badger.DefaultOptions(dir).WithLogger(nil))
	require.Nil(t, err)

	e := Entity{ID: 1337}
	key := []byte{0x01, 0x02, 0x03}
	val := []byte(`{"ID":1337}`)

	_ = db.Update(func(tx *badger.Txn) error {
		err := tx.Set(key, []byte{})
		require.Nil(t, err)
		return nil
	})

	err = db.Update(update(key, e))
	require.Nil(t, err)

	var act []byte
	_ = db.View(func(tx *badger.Txn) error {
		item, err := tx.Get(key)
		require.Nil(t, err)
		act, err = item.ValueCopy(nil)
		require.Nil(t, err)
		return nil
	})

	assert.Equal(t, act, val)
}

func TestUpdateMissing(t *testing.T) {
	// TODO
}

func TestRetrieveValid(t *testing.T) {

	dir := filepath.Join(os.TempDir(), fmt.Sprintf("flow-test-db-%d", rand.Uint64()))
	db, err := badger.Open(badger.DefaultOptions(dir).WithLogger(nil))
	require.Nil(t, err)

	e := Entity{ID: 1337}
	key := []byte{0x01, 0x02, 0x03}
	val := []byte(`{"ID":1337}`)

	_ = db.Update(func(tx *badger.Txn) error {
		err := tx.Set(key, val)
		require.Nil(t, err)
		return nil
	})

	var act Entity
	err = db.View(retrieve(key, &act))
	require.Nil(t, err)

	assert.Equal(t, act, e)
}

func TestRetrieveMissing(t *testing.T) {
	// TODO
}

func TestScan(t *testing.T) {
	// TODO
}

func TestTraverse(t *testing.T) {
	// TODO
}
