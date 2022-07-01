package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Silverfish struct{ Mob }

func SilverfishFromNbt(node *nbt.TagNodeCompound) *Silverfish {
	return &Silverfish{Mob: *MobFromNbt(node)}
}
