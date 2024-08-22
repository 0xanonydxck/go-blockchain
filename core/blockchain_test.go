package core

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newBlockchainWithGenesis(t *testing.T) *Blockchain {
	bc, err := NewBlockchain(randomBlock(0))
	assert.Nil(t, err)
	return bc
}

func TestAddBlock(t *testing.T) {
	bc := newBlockchainWithGenesis(t)

	lenBlock := 1000
	for i := 1; i <= lenBlock; i++ {
		block := randomBlockWithSignature(t, uint32(i))
		assert.Nil(t, bc.AddBlock(block))
	}

	assert.Equal(t, uint32(lenBlock), bc.Height())
	assert.Equal(t, len(bc.headers), lenBlock+1)
	assert.NotNil(t, bc.AddBlock(randomBlock(89)))
}

func TestNewBlockchain(t *testing.T) {
	bc := newBlockchainWithGenesis(t)
	assert.NotNil(t, bc.validator)
	assert.Equal(t, uint32(0), bc.Height())

	fmt.Println(bc.Height())
}

func TestHasBlock(t *testing.T) {
	bc := newBlockchainWithGenesis(t)
	assert.True(t, bc.HasBlock(0))
}
