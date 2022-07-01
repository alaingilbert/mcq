package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Enderman struct{ Mob }

func EndermanFromNbt(node *nbt.TagNodeCompound) *Enderman {
	return &Enderman{Mob: *MobFromNbt(node)}
}
