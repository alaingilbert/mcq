package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Pig struct{ Mob }

func PigFromNbt(node *nbt.TagNodeCompound) *Pig {
	return &Pig{Mob: *MobFromNbt(node)}
}
