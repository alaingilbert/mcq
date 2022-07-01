package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Salmon struct{ Mob }

func SalmonFromNbt(node *nbt.TagNodeCompound) *Salmon {
	return &Salmon{Mob: *MobFromNbt(node)}
}
