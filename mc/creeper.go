package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Creeper struct{ Mob }

func CreeperFromNbt(node *nbt.TagNodeCompound) *Creeper {
	return &Creeper{Mob: *MobFromNbt(node)}
}
