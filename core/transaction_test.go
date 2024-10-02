package core

import (
	"bytes"
	"testing"

	"github.com/hussain18/learn-projectx/crypto"
	"github.com/stretchr/testify/assert"
)

func TestSignTransaction(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	pubKey := privKey.PublicKey()

	tx := Transaction{
		Data: []byte("foo"),
	}

	assert.Nil(t, tx.Sign(privKey))
	assert.NotNil(t, tx.Signature)

	assert.Equal(t, pubKey, tx.From)
	assert.True(t, tx.Signature.Verify(pubKey, tx.Data))

	otherData := []byte("bar")
	otherPrivKey := crypto.GeneratePrivateKey()
	otherPubKey := otherPrivKey.PublicKey()

	assert.False(t, tx.Signature.Verify(otherPubKey, tx.Data))
	tx.Data = otherData
	assert.False(t, tx.Signature.Verify(pubKey, tx.Data))
}

func TestVerifyTransaction(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()

	tx := Transaction{
		Data: []byte("foo"),
	}

	assert.Nil(t, tx.Sign(privKey))
	assert.Nil(t, tx.Verify())

	otherData := []byte("bar")
	tx.Data = otherData
	assert.NotNil(t, tx.Verify())
}

func TestTxEncodeDecode(t *testing.T) {
	tx := randomTxWithSignature(t)
	buf := &bytes.Buffer{}

	assert.Nil(t, tx.Encode(NewGobTxEncoder(buf)))

	txDecoded := new(Transaction)
	assert.Nil(t, txDecoded.Decode(NewGobTxDecoder(buf)))
	assert.Equal(t, txDecoded, tx)
}

func randomTxWithSignature(t *testing.T) *Transaction {
	privKey := crypto.GeneratePrivateKey()

	tx := &Transaction{
		Data: []byte("Hi there"),
	}
	assert.Nil(t, tx.Sign(privKey))
	return tx
}
