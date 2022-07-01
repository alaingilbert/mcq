package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Parrot struct{ Mob }

func ParrotFromNbt(node *nbt.TagNodeCompound) *Parrot {
	return &Parrot{Mob: *MobFromNbt(node)}
}
