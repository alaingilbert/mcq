package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Dolphin struct{ Mob }

func DolphinFromNbt(node *nbt.TagNodeCompound) *Dolphin {
	return &Dolphin{Mob: *MobFromNbt(node)}
}
