package core

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type BlockChain struct {
	headers   []*Header
	store     Storage
	validator Validator
}

func NewBlockChain(genesis *Block) (*BlockChain, error) {
	bc := &BlockChain{
		headers: []*Header{},
	}

	bc.store = NewMemoryStore()
	bc.validator = newBlockValidator(bc)
	err := bc.addBlockWithoutValidation(genesis)

	return bc, err
}

func (bc *BlockChain) SetValidator(v Validator) {
	bc.validator = v
}

func (bc *BlockChain) Height() uint32 {
	return uint32(len(bc.headers) - 1)
}

func (bc *BlockChain) AddBlock(b *Block) error {
	if err := bc.validator.ValidateBlock(b); err != nil {
		return err
	}

	return bc.addBlockWithoutValidation(b)
}

func (bc *BlockChain) HasBlock(height uint32) bool {
	return height <= bc.Height()
}

func (bc *BlockChain) GetHeader(h uint32) (*Header, error) {
	if h > bc.Height() {
		return nil, fmt.Errorf("block (%d) too high", bc.Height())
	}

	return bc.headers[len(bc.headers)-1], nil
}

func (bc *BlockChain) addBlockWithoutValidation(b *Block) error {
	logrus.WithField("NEW_BLOCK", logrus.Fields{
		"Height": b.Height,
		"Hash":   b.Hash(BlockHasher{}),
	}).Info("new block added")

	bc.headers = append(bc.headers, b.Header)

	return bc.store.Put(b)
}
