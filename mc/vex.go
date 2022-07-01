package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Vex struct{ Mob }

func VexFromNbt(node *nbt.TagNodeCompound) *Vex {
	return &Vex{Mob: *MobFromNbt(node)}
}
