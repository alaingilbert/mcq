package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Barrel struct {
	BaseChest
}

func BarrelFromNbt(node *nbt.TagNodeCompound) *Barrel {
	b := new(Barrel)
	b.BaseChest = *BaseChestFromNbt(node)
	return b
}
