package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newBlockChainWithGenesis(t *testing.T) *BlockChain {
	bc, err := NewBlockChain(randomBlock(0))
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
		b := randomBlockWithSignature(t, uint32(i+1))
		assert.Nil(t, bc.AddBlock(b))
	}

	assert.Equal(t, uint32(lenBlock), bc.Height())
	assert.Equal(t, lenBlock+1, len(bc.headers))

	assert.NotNil(t, bc.AddBlock(randomBlockWithSignature(t, 89)))
}
