package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type TrappedChest struct {
	BaseChest
}

func TrappedChestFromNbt(node *nbt.TagNodeCompound) *TrappedChest {
	c := new(TrappedChest)
	c.BaseChest = *BaseChestFromNbt(node)
	return c
}
