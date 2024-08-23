package network

import (
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
