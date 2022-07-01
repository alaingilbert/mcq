package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Strider struct{ Mob }

func StriderFromNbt(node *nbt.TagNodeCompound) *Strider {
	return &Strider{Mob: *MobFromNbt(node)}
}
