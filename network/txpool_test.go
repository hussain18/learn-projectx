package network

import (
	"math/rand"
	"testing"

	"github.com/hussain18/learn-projectx/core"
	"github.com/stretchr/testify/assert"
)

func TestTxPoo(t *testing.T) {
	p := NewTxPool()

	assert.Equal(t, 0, p.Len())
}

func TestTxPoolAddTx(t *testing.T) {
	p := NewTxPool()
	tx := core.NewTransaction([]byte("Hey there"))

	assert.Nil(t, p.Add(tx))
	assert.Equal(t, 1, p.Len())

	core.NewTransaction([]byte("Hey there"))
	assert.Equal(t, 1, p.Len())

	p.Flush()

	assert.Equal(t, 0, p.Len())
}

func TestSortTransactions(t *testing.T) {
	p := NewTxPool()
	txLen := 1000

	for i := 0; i < txLen; i++ {
		tx := core.NewTransaction([]byte(string(rune(i))))
		tx.SetFirstSeen(int64(i * rand.Intn(10000)))
		assert.Nil(t, p.Add(tx))
	}

	txx := p.Transactions()
	assert.Equal(t, len(txx), txLen)

	for i := 1; i < txLen; i++ {
		assert.Greater(t, txx[i].FirstSeen(), txx[i-1].FirstSeen())
	}
}
