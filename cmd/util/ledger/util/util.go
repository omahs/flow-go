package util

import (
	"fmt"

	"github.com/onflow/atree"
	"github.com/onflow/cadence/runtime/common"

	"github.com/onflow/flow-go/engine/execution/state"
	"github.com/onflow/flow-go/fvm/environment"
	"github.com/onflow/flow-go/fvm/storage/snapshot"
	"github.com/onflow/flow-go/ledger"
	"github.com/onflow/flow-go/model/flow"
)

func KeyToRegisterID(key ledger.Key) (flow.RegisterID, error) {
	if len(key.KeyParts) != 2 ||
		key.KeyParts[0].Type != state.KeyPartOwner ||
		key.KeyParts[1].Type != state.KeyPartKey {
		return flow.RegisterID{}, fmt.Errorf("key not in expected format %s", key.String())
	}

	return flow.NewRegisterID(
		string(key.KeyParts[0].Value),
		string(key.KeyParts[1].Value),
	), nil
}

func RegisterIDToKey(registerID flow.RegisterID) ledger.Key {
	newKey := ledger.Key{}
	newKey.KeyParts = []ledger.KeyPart{
		{
			Type:  state.KeyPartOwner,
			Value: []byte(registerID.Owner),
		},
		{
			Type:  state.KeyPartKey,
			Value: []byte(registerID.Key),
		},
	}
	return newKey
}

type AccountsAtreeLedger struct {
	Accounts environment.Accounts
}

func NewAccountsAtreeLedger(accounts environment.Accounts) *AccountsAtreeLedger {
	return &AccountsAtreeLedger{Accounts: accounts}
}

var _ atree.Ledger = &AccountsAtreeLedger{}

func (a *AccountsAtreeLedger) GetValue(owner, key []byte) ([]byte, error) {
	v, err := a.Accounts.GetValue(
		flow.NewRegisterID(
			string(flow.BytesToAddress(owner).Bytes()),
			string(key)))
	if err != nil {
		return nil, fmt.Errorf("getting value failed: %w", err)
	}
	return v, nil
}

func (a *AccountsAtreeLedger) SetValue(owner, key, value []byte) error {
	err := a.Accounts.SetValue(
		flow.NewRegisterID(
			string(flow.BytesToAddress(owner).Bytes()),
			string(key)),
		value)
	if err != nil {
		return fmt.Errorf("setting value failed: %w", err)
	}
	return nil
}

func (a *AccountsAtreeLedger) ValueExists(owner, key []byte) (exists bool, err error) {
	v, err := a.GetValue(owner, key)
	if err != nil {
		return false, fmt.Errorf("checking value existence failed: %w", err)
	}

	return len(v) > 0, nil
}

// AllocateStorageIndex allocates new storage index under the owner accounts to store a new register
func (a *AccountsAtreeLedger) AllocateStorageIndex(owner []byte) (atree.StorageIndex, error) {
	v, err := a.Accounts.AllocateStorageIndex(flow.BytesToAddress(owner))
	if err != nil {
		return atree.StorageIndex{}, fmt.Errorf("storage address allocation failed: %w", err)
	}
	return v, nil
}

type PayloadSnapshot struct {
	Payloads map[flow.RegisterID]flow.RegisterValue
}

var _ snapshot.StorageSnapshot = (*PayloadSnapshot)(nil)

func NewPayloadSnapshot(payloads []ledger.Payload) (*PayloadSnapshot, error) {
	l := &PayloadSnapshot{
		Payloads: make(map[flow.RegisterID][]byte, len(payloads)),
	}
	for _, payload := range payloads {
		key, err := payload.Key()
		if err != nil {
			return nil, err
		}
		id, err := KeyToRegisterID(key)
		if err != nil {
			return nil, err
		}
		l.Payloads[id] = payload.Value()
	}
	return l, nil
}

func (p PayloadSnapshot) Get(id flow.RegisterID) (flow.RegisterValue, error) {
	value := p.Payloads[id]
	return value, nil
}

// NopMemoryGauge is a no-op implementation of the MemoryGauge interface
type NopMemoryGauge struct{}

func (n NopMemoryGauge) MeterMemory(common.MemoryUsage) error {
	return nil
}

var _ common.MemoryGauge = (*NopMemoryGauge)(nil)
