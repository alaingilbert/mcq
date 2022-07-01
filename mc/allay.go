package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Allay struct{ Mob }

func AllayFromNbt(node *nbt.TagNodeCompound) *Allay {
	return &Allay{Mob: *MobFromNbt(node)}
}
