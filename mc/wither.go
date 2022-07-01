package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Wither struct{ Mob }

func WitherFromNbt(node *nbt.TagNodeCompound) *Wither {
	return &Wither{Mob: *MobFromNbt(node)}
}
