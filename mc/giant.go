package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Giant struct{ Mob }

func GiantFromNbt(node *nbt.TagNodeCompound) *Giant {
	return &Giant{Mob: *MobFromNbt(node)}
}
