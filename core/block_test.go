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

func randomBlockWithSignature(t *testing.T, height uint32) *Block {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(height)

	assert.Nil(t, b.Sign(privKey))
	return b
}

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
