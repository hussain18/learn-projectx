package core

import "fmt"

type Validator interface {
	ValidateBlock(*Block) error
}

type BlockValidator struct {
	bc *BlockChain
}

func newBlockValidator(bc *BlockChain) *BlockValidator {
	return &BlockValidator{bc: bc}
}

func (v *BlockValidator) ValidateBlock(b *Block) error {
	if v.bc.HasBlock(b.Height) {
		return fmt.Errorf("chain already contains block (%d) with hash (%d)", b.Height, b.Hash(BlockHasher{}))
	}

	if err := b.Verify(); err != nil {
		return fmt.Errorf("invalid block")
	}

	return nil
}
