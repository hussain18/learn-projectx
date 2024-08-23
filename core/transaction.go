package core

import (
	"fmt"

	"github.com/hussain18/learn-projectx/crypto"
)

type Transaction struct {
	Data []byte

	From      crypto.PublicKey
	Signature *crypto.Signature
}

func (t *Transaction) Sign(privKey crypto.PrivateKey) error {
	sig, err := privKey.Sign(t.Data)
	if err != nil {
		return err
	}

	t.Signature = sig
	t.From = privKey.PublicKey()

	return nil
}

func (t *Transaction) Verify() error {
	if t.Signature == nil {
		return fmt.Errorf("transaction has no signature")
	}

	if !t.Signature.Verify(t.From, t.Data) {
		return fmt.Errorf("transaction has invalid signature")
	}

	return nil
}

// func (t *Transaction) EncodeBinary(w io.Writer) error { return nil }
// func (t *Transaction) DecodeBinary(r io.Reader) error { return nil }
