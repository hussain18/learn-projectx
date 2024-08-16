package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"io"

	"github.com/hussain18/learn-projectx/types"
)

type Header struct {
	Version   uint32
	PrevBlock types.Hash
	Timestamp uint64
	Height    uint32
	Nonce     uint64
}

func (h *Header) EncodeBinary(w io.Writer) error {

	if err := binary.Write(w, binary.LittleEndian, &h.Version); err != nil {
		return err
	}

	if err := binary.Write(w, binary.LittleEndian, &h.PrevBlock); err != nil {
		return err
	}

	if err := binary.Write(w, binary.LittleEndian, &h.Timestamp); err != nil {
		return err
	}

	if err := binary.Write(w, binary.LittleEndian, &h.Height); err != nil {
		return err
	}

	if err := binary.Write(w, binary.LittleEndian, &h.Nonce); err != nil {
		return err
	}

	return nil
}

func (h *Header) DecodeBinary(r io.Reader) error {
	if err := binary.Read(r, binary.LittleEndian, &h.Version); err != nil {
		return err
	}

	if err := binary.Read(r, binary.LittleEndian, &h.PrevBlock); err != nil {
		return err
	}

	if err := binary.Read(r, binary.LittleEndian, &h.Timestamp); err != nil {
		return err
	}

	if err := binary.Read(r, binary.LittleEndian, &h.Height); err != nil {
		return err
	}

	if err := binary.Read(r, binary.LittleEndian, &h.Nonce); err != nil {
		return err
	}

	return nil
}

type Block struct {
	Header
	Transactions []Transaction

	hash types.Hash
}

func (b *Block) Hash() types.Hash {
	buf := &bytes.Buffer{}
	b.Header.EncodeBinary(buf)

	if b.hash.IsZero() {
		b.hash = types.Hash(sha256.Sum256(buf.Bytes()))
	}

	return b.hash
}

func (b *Block) EncodeBinary(w io.Writer) error {

	if err := b.Header.EncodeBinary(w); err != nil {
		return err
	}

	for _, tx := range b.Transactions {
		if err := tx.EncodeBinary(w); err != nil {
			return err
		}
	}

	return nil
}

func (b *Block) DecodeBinary(r io.Reader) error {
	if err := b.Header.DecodeBinary(r); err != nil {
		return err
	}

	for _, tx := range b.Transactions {
		if err := tx.DecodeBinary(r); err != nil {
			return err
		}
	}

	return nil
}
