package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Ravager struct{ Mob }

func RavagerFromNbt(node *nbt.TagNodeCompound) *Ravager {
	return &Ravager{Mob: *MobFromNbt(node)}
}
