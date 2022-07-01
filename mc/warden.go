package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Warden struct{ Mob }

func WardenFromNbt(node *nbt.TagNodeCompound) *Warden {
	return &Warden{Mob: *MobFromNbt(node)}
}
