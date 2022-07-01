package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Pillager struct{ Mob }

func PillagerFromNbt(node *nbt.TagNodeCompound) *Pillager {
	return &Pillager{Mob: *MobFromNbt(node)}
}
