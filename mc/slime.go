package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Slime struct{ Mob }

func SlimeFromNbt(node *nbt.TagNodeCompound) *Slime {
	return &Slime{Mob: *MobFromNbt(node)}
}
