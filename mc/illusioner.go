package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Illusioner struct{ Mob }

func IllusionerFromNbt(node *nbt.TagNodeCompound) *Illusioner {
	return &Illusioner{Mob: *MobFromNbt(node)}
}
