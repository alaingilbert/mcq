package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type EnderDragon struct{ Mob }

func EnderDragonFromNbt(node *nbt.TagNodeCompound) *EnderDragon {
	return &EnderDragon{Mob: *MobFromNbt(node)}
}
