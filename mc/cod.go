package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Cod struct{ Mob }

func CodFromNbt(node *nbt.TagNodeCompound) *Cod {
	return &Cod{Mob: *MobFromNbt(node)}
}
