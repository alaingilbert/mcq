package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Witch struct{ Mob }

func WitchFromNbt(node *nbt.TagNodeCompound) *Witch {
	return &Witch{Mob: *MobFromNbt(node)}
}
