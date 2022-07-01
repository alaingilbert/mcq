package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Bat struct{ Mob }

func BatFromNbt(node *nbt.TagNodeCompound) *Bat {
	return &Bat{Mob: *MobFromNbt(node)}
}
