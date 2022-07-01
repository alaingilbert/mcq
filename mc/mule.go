package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Mule struct{ Mob }

func MuleFromNbt(node *nbt.TagNodeCompound) *Mule {
	return &Mule{Mob: *MobFromNbt(node)}
}
