package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Vindicator struct{ Mob }

func VindicatorFromNbt(node *nbt.TagNodeCompound) *Vindicator {
	return &Vindicator{Mob: *MobFromNbt(node)}
}
