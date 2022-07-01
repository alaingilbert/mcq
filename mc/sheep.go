package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Sheep struct{ Mob }

func SheepFromNbt(node *nbt.TagNodeCompound) *Sheep {
	return &Sheep{Mob: *MobFromNbt(node)}
}
