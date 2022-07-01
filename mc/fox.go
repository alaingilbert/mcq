package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Fox struct{ Mob }

func FoxFromNbt(node *nbt.TagNodeCompound) *Fox {
	return &Fox{Mob: *MobFromNbt(node)}
}
