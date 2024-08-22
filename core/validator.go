package core

import "fmt"

type Validator interface {
	ValidateBlock(b *Block) error
}

type BlockValidator struct {
	bc *Blockchain
}

func NewBlockValidator(bc *Blockchain) *BlockValidator {
	return &BlockValidator{bc: bc}
}

func (v *BlockValidator) ValidateBlock(b *Block) error {
	if v.bc.HasBlock(b.Height) {
		return fmt.Errorf("chain already contained block with {%v} with hash {%v}", b.Height, b.Hash(BlockHasher{}))
	}

	if err := b.Verify(); err != nil {
		return err
	}

	return nil
}
