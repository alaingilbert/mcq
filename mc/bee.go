package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Bee struct{ Mob }

func BeeFromNbt(node *nbt.TagNodeCompound) *Bee {
	return &Bee{Mob: *MobFromNbt(node)}
}
