package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Frog struct{ Mob }

func FrogFromNbt(node *nbt.TagNodeCompound) *Frog {
	return &Frog{Mob: *MobFromNbt(node)}
}
