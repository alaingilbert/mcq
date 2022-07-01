package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Tadpole struct{ Mob }

func TadpoleFromNbt(node *nbt.TagNodeCompound) *Tadpole {
	return &Tadpole{Mob: *MobFromNbt(node)}
}
