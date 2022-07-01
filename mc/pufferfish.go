package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Pufferfish struct{ Mob }

func PufferfishFromNbt(node *nbt.TagNodeCompound) *Pufferfish {
	return &Pufferfish{Mob: *MobFromNbt(node)}
}
