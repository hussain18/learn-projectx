package core

import (
	"testing"
	"time"

	"github.com/hussain18/learn-projectx/crypto"
	"github.com/hussain18/learn-projectx/types"
	"github.com/stretchr/testify/assert"
)

func randomBlock(t *testing.T, height uint32, prevBlockHash types.Hash) *Block {
	b := &Block{
		Header: &Header{
			Version:       1,
			PrevBlockHash: prevBlockHash,
			Timestamp:     uint64(time.Now().UnixNano()),
			Height:        height,
		},
	}
	b.AddTransaction(randomTxWithSignature(t))

	return b
}

func randomBlockWithSignature(t *testing.T, height uint32, prevBlockHash types.Hash) *Block {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(t, height, prevBlockHash)

	assert.Nil(t, b.Sign(privKey))
	return b
}

/* -------------------------- My Idea was stupid!!! ------------------------- */

// func randomBlockWithPrevHash(t *testing.T, height uint32, prevBlockHash types.Hash) *Block {
// 	b := randomBlockWithSignature(t, height)
// 	b.PrevBlockHash = prevBlockHash

// 	return b
// }

/* -------------------------- My Idea was stupid!!! ------------------------- */

func TestBlockHash(t *testing.T) {
	b := randomBlock(t, 1, types.Hash{})

	hash := b.Hash(BlockHasher{})
	assert.False(t, hash.IsZero())
}

func TestSignBlock(t *testing.T) {

	privKey := crypto.GeneratePrivateKey()
	pubKey := privKey.PublicKey()

	b := randomBlock(t, 1, types.Hash{})

	assert.Nil(t, b.Sign(privKey))
	assert.NotNil(t, b.Signature)

	assert.Equal(t, pubKey, b.Validator)
	assert.True(t, b.Signature.Verify(pubKey, b.Header.Bytes()))
}

func TestVerifyBlock(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	pubKey := privKey.PublicKey()

	b := randomBlock(t, 1, types.Hash{})

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
