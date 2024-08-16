package core

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/hussain18/learn-projectx/types"
	"github.com/stretchr/testify/assert"
)

func TestHeader_Encode_Decode(t *testing.T) {
	h := &Header{
		Version:   1,
		PrevBlock: types.RandomHashFromBytes(),
		Timestamp: uint64(time.Now().UnixNano()),
		Height:    10,
		Nonce:     24354,
	}

	buf := &bytes.Buffer{}
	assert.Nil(t, h.EncodeBinary(buf))

	hDecode := &Header{}
	assert.Nil(t, hDecode.DecodeBinary(buf))
	assert.Equal(t, h, hDecode)
}

func TestBlock_Encode_Decode(t *testing.T) {
	h := &Header{
		Version:   1,
		PrevBlock: types.RandomHashFromBytes(),
		Timestamp: uint64(time.Now().UnixNano()),
		Height:    10,
		Nonce:     24354,
	}

	b := &Block{
		Header:       *h,
		Transactions: nil,
	}

	buf := &bytes.Buffer{}
	assert.Nil(t, b.EncodeBinary(buf))

	bDecode := &Block{}
	assert.Nil(t, bDecode.DecodeBinary(buf))
	assert.Equal(t, b, bDecode)
}

func TestBlockHash(t *testing.T) {
	h := &Header{
		Version:   1,
		PrevBlock: types.RandomHashFromBytes(),
		Timestamp: uint64(time.Now().UnixNano()),
		Height:    10,
		Nonce:     24354,
	}

	b := &Block{
		Header:       *h,
		Transactions: nil,
	}

	hash := b.Hash()
	fmt.Println(hash)
	assert.False(t, hash.IsZero())
}
