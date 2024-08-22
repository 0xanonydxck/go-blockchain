package core

import (
	"io"

	"github.com/dxckboi/go-blockchain/types"
)

type Transaction struct {
	Data []byte

	From types.Address
}

func (t *Transaction) EncodeBinary(w io.Writer) error {
	return nil
}

func (t *Transaction) DecodeBinary(r io.Reader) error {
	return nil
}
