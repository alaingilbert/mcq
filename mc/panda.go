package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Panda struct{ Mob }

func PandaFromNbt(node *nbt.TagNodeCompound) *Panda {
	return &Panda{Mob: *MobFromNbt(node)}
}
