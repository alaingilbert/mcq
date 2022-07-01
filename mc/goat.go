package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Goat struct{ Mob }

func GoatFromNbt(node *nbt.TagNodeCompound) *Goat {
	return &Goat{Mob: *MobFromNbt(node)}
}
