package core

import (
	"testing"
	"time"

	"github.com/hussain18/learn-projectx/crypto"
	"github.com/hussain18/learn-projectx/types"
	"github.com/stretchr/testify/assert"
)

func randomBlock(height uint32) *Block {
	h := &Header{
		Version:       1,
		PrevBlockHash: types.RandomHashFromBytes(),
		Timestamp:     uint64(time.Now().UnixNano()),
		Height:        height,
	}

	tx := Transaction{
		Data: []byte("Hey there"),
	}

	return NewBlock(h, []Transaction{tx})
}

// func TestHeader_Encode_Decode(t *testing.T) {
// 	h := &Header{
// 		Version:   1,
// 		PrevBlock: types.RandomHashFromBytes(),
// 		Timestamp: uint64(time.Now().UnixNano()),
// 		Height:    10,
// 		Nonce:     24354,
// 	}

// 	buf := &bytes.Buffer{}
// 	assert.Nil(t, h.EncodeBinary(buf))

// 	hDecode := &Header{}
// 	assert.Nil(t, hDecode.DecodeBinary(buf))
// 	assert.Equal(t, h, hDecode)
// }

// func TestBlock_Encode_Decode(t *testing.T) {
// 	h := &Header{
// 		Version:   1,
// 		PrevBlock: types.RandomHashFromBytes(),
// 		Timestamp: uint64(time.Now().UnixNano()),
// 		Height:    10,
// 		Nonce:     24354,
// 	}

// 	b := &Block{
// 		Header:       *h,
// 		Transactions: nil,
// 	}

// 	buf := &bytes.Buffer{}
// 	assert.Nil(t, b.EncodeBinary(buf))

// 	bDecode := &Block{}
// 	assert.Nil(t, bDecode.DecodeBinary(buf))
// 	assert.Equal(t, b, bDecode)
// }

func TestBlockHash(t *testing.T) {
	b := randomBlock(1)

	hash := b.Hash(BlockHasher{})
	assert.False(t, hash.IsZero())
}

func TestSignBlock(t *testing.T) {

	privKey := crypto.GeneratePrivateKey()
	pubKey := privKey.PublicKey()

	b := randomBlock(1)

	assert.Nil(t, b.Sign(privKey))
	assert.NotNil(t, b.Signature)

	assert.Equal(t, pubKey, b.Validator)
	assert.True(t, b.Signature.Verify(pubKey, b.HeaderBytes()))
}

func TestVerifyBlock(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	pubKey := privKey.PublicKey()

	b := randomBlock(1)

	assert.Nil(t, b.Sign(privKey))
	assert.Nil(t, b.Verify())
	assert.Equal(t, pubKey, b.Validator)

	otherPrivKey := crypto.GeneratePrivateKey()
	otherPubKey := otherPrivKey.PublicKey()

	b.Validator = otherPubKey
	assert.NotNil(t, b.Verify())

	// b.Validator = pubKey // Note: Flaw here...
	b.Height = 100
	assert.NotNil(t, b.Verify())
}
