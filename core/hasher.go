package core

import (
	"crypto/sha256"

	"github.com/hussain18/learn-projectx/types"
)

type Hasher[T any] interface {
	Hash(T) types.Hash
}

type BlockHasher struct{}

func (BlockHasher) Hash(b *Block) types.Hash {
	h := sha256.Sum256(b.HeaderBytes())
	return types.Hash(h)
}
