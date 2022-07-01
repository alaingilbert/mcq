package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Piglin struct{ Mob }

func PiglinFromNbt(node *nbt.TagNodeCompound) *Piglin {
	return &Piglin{Mob: *MobFromNbt(node)}
}
