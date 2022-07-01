package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Evoker struct{ Mob }

func EvokerFromNbt(node *nbt.TagNodeCompound) *Evoker {
	return &Evoker{Mob: *MobFromNbt(node)}
}
