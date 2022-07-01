package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Squid struct{ Mob }

func SquidFromNbt(node *nbt.TagNodeCompound) *Squid {
	return &Squid{Mob: *MobFromNbt(node)}
}
