package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type TrappedChest struct{ ContainerBlockEntity }

func TrappedChestFromNbt(node *nbt.TagNodeCompound) *TrappedChest {
	return &TrappedChest{ContainerBlockEntity: *ContainerBlockEntityFromNbt(node)}
}
