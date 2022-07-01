package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Ghast struct{ Mob }

func GhastFromNbt(node *nbt.TagNodeCompound) *Ghast {
	return &Ghast{Mob: *MobFromNbt(node)}
}
