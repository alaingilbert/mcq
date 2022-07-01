package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Stray struct{ Mob }

func StrayFromNbt(node *nbt.TagNodeCompound) *Stray {
	return &Stray{Mob: *MobFromNbt(node)}
}
