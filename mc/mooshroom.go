package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Mooshroom struct{ Mob }

func MooshroomFromNbt(node *nbt.TagNodeCompound) *Mooshroom {
	return &Mooshroom{Mob: *MobFromNbt(node)}
}
