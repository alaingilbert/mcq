package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Endermite struct{ Mob }

func EndermiteFromNbt(node *nbt.TagNodeCompound) *Endermite {
	return &Endermite{Mob: *MobFromNbt(node)}
}
