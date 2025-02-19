package request

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetAccountKey_InvalidParse(t *testing.T) {
	var getAccountKey GetAccountKey

	tests := []struct {
		name    string
		address string
		index   string
		height  string
		err     string
	}{
		{
			"parse with invalid address",
			"0xxxaddr",
			"1",
			"100",
			"invalid address",
		},
		{
			"parse with invalid keyIndex",
			"0xf8d6e0586b0a20c7",
			"-1.2",
			"100",
			"invalid key index: value must be an unsigned 64 bit integer",
		},
		{
			"parse with invalid height",
			"0xf8d6e0586b0a20c7",
			"2",
			"-100",
			"invalid height format",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := getAccountKey.Parse(test.address, test.index, test.height)
			assert.EqualError(t, err, test.err)
		})
	}
}

func Test_GetAccountKey_ValidParse(t *testing.T) {
	var getAccountKey GetAccountKey

	addr := "f8d6e0586b0a20c7"
	keyIndex := "5"
	height := "100"
	err := getAccountKey.Parse(addr, keyIndex, height)
	assert.NoError(t, err)
	assert.Equal(t, getAccountKey.Address.String(), addr)
	assert.Equal(t, getAccountKey.Index, uint64(5))
	assert.Equal(t, getAccountKey.Height, uint64(100))

	err = getAccountKey.Parse(addr, keyIndex, "")
	assert.NoError(t, err)
	assert.Equal(t, getAccountKey.Address.String(), addr)
	assert.Equal(t, getAccountKey.Index, uint64(5))
	assert.Equal(t, getAccountKey.Height, SealedHeight)

	err = getAccountKey.Parse(addr, keyIndex, "sealed")
	assert.NoError(t, err)
	assert.Equal(t, getAccountKey.Address.String(), addr)
	assert.Equal(t, getAccountKey.Index, uint64(5))
	assert.Equal(t, getAccountKey.Height, SealedHeight)

	err = getAccountKey.Parse(addr, keyIndex, "final")
	assert.NoError(t, err)
	assert.Equal(t, getAccountKey.Address.String(), addr)
	assert.Equal(t, getAccountKey.Index, uint64(5))
	assert.Equal(t, getAccountKey.Height, FinalHeight)
}
