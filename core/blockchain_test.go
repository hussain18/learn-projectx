package core

import (
	"testing"

	"github.com/hussain18/learn-projectx/types"
	"github.com/stretchr/testify/assert"
)

func newBlockChainWithGenesis(t *testing.T) *BlockChain {
	bc, err := NewBlockChain(randomBlock(t, 0, types.Hash{}))
	assert.NotNil(t, bc.validator)
	assert.Nil(t, err)

	return bc
}

func TestNewBlockchain(t *testing.T) {
	bc := newBlockChainWithGenesis(t)
	assert.Equal(t, uint32(0), bc.Height())
}

func TestHasBlock(t *testing.T) {
	bc := newBlockChainWithGenesis(t)
	assert.True(t, bc.HasBlock(uint32(0)))
}

func TestAddBlock(t *testing.T) {
	bc := newBlockChainWithGenesis(t)
	lenBlock := 1000

	for i := 0; i < lenBlock; i++ {
		b := randomBlockWithSignature(t, uint32(i+1), getPrevBlockHash(t, bc, uint32(i+1)))
		assert.Nil(t, bc.AddBlock(b))
	}
	assert.Equal(t, uint32(lenBlock), bc.Height())
	assert.Equal(t, lenBlock+1, len(bc.headers))

	tooSmallBlockHeight := uint32(88)
	assert.NotNil(t, bc.AddBlock(randomBlockWithSignature(t, tooSmallBlockHeight, getPrevBlockHash(t, bc, tooSmallBlockHeight-1))))
}

func TestAddBlockTooHigh(t *testing.T) {
	bc := newBlockChainWithGenesis(t)
	lenBlock := 1000

	for i := 0; i < lenBlock; i++ {
		b := randomBlockWithSignature(t, uint32(i+1), getPrevBlockHash(t, bc, uint32(i+1)))
		assert.Nil(t, bc.AddBlock(b))
	}

	assert.NotNil(t, bc.AddBlock(randomBlockWithSignature(t, uint32(lenBlock+2), getPrevBlockHash(t, bc, uint32(lenBlock+1)))))
}

func TestGetHeader(t *testing.T) {
	bc := newBlockChainWithGenesis(t)
	lenBlock := 1000

	// prevHeader, err := bc.GetHeader(0)
	// assert.Nil(t, err)

	// prevBlockHash := BlockHasher{}.Hash(prevHeader)

	for i := 0; i < lenBlock; i++ {
		b := randomBlockWithSignature(t, uint32(i+1), getPrevBlockHash(t, bc, uint32(i+1)))
		assert.Nil(t, bc.AddBlock(b))

		header, err := bc.GetHeader(b.Height)
		assert.Nil(t, err)
		assert.Equal(t, b.Header, header)

		// prevBlockHash = BlockHasher{}.Hash(header)
	}
}

func getPrevBlockHash(t *testing.T, bc *BlockChain, height uint32) types.Hash {
	prevHeader, err := bc.GetHeader(height - 1)
	assert.Nil(t, err)

	return BlockHasher{}.Hash(prevHeader)
}
