package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Drowned struct{ Mob }

func DrownedFromNbt(node *nbt.TagNodeCompound) *Drowned {
	return &Drowned{Mob: *MobFromNbt(node)}
}
