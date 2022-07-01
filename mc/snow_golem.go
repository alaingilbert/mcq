package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type SnowGolem struct{ Mob }

func SnowGolemFromNbt(node *nbt.TagNodeCompound) *SnowGolem {
	return &SnowGolem{Mob: *MobFromNbt(node)}
}
