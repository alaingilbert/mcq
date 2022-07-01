package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Wolf struct{ Mob }

func WolfFromNbt(node *nbt.TagNodeCompound) *Wolf {
	return &Wolf{Mob: *MobFromNbt(node)}
}
