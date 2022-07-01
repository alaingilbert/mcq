package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Cow struct{ Mob }

func CowFromNbt(node *nbt.TagNodeCompound) *Cow {
	return &Cow{Mob: *MobFromNbt(node)}
}
