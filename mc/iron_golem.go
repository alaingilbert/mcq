package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type IronGolem struct{ Mob }

func IronGolemFromNbt(node *nbt.TagNodeCompound) *IronGolem {
	return &IronGolem{Mob: *MobFromNbt(node)}
}
