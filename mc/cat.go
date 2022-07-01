package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Cat struct{ Mob }

func CatFromNbt(node *nbt.TagNodeCompound) *Cat {
	return &Cat{Mob: *MobFromNbt(node)}
}
