package core

import (
	"testing"
	"time"

	"github.com/dxckboi/go-blockchain/crypto"
	"github.com/dxckboi/go-blockchain/types"
	"github.com/stretchr/testify/assert"
)

func randomBlock(height uint32) *Block {
	tx := Transaction{
		Data: []byte("foo"),
	}

	return NewBlock(&Header{
		Version:       1,
		DataHash:      types.Hash{1, 2, 3, 4},
		PrevBlockHash: types.RandomHash(),
		Timestamp:     time.Now().UnixNano(),
		Height:        height,
	}, []Transaction{tx})
}

func TestSignBlock(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(0)

	assert.Nil(t, b.Sign(privKey))
	assert.NotNil(t, b.Signature)
}

func TestVerifyBlock(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(0)

	assert.Nil(t, b.Sign(privKey))
	assert.Nil(t, b.Verify())

	otherPrivKey := crypto.GeneratePrivateKey()
	b.Validator = otherPrivKey.PublicKey()
	assert.NotNil(t, b.Verify())

	b.Height = 100
	assert.NotNil(t, b.Verify())
}
