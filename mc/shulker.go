package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Shulker struct{ Mob }

func ShulkerFromNbt(node *nbt.TagNodeCompound) *Shulker {
	return &Shulker{Mob: *MobFromNbt(node)}
}
