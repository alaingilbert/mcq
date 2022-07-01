package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Husk struct{ Mob }

func HuskFromNbt(node *nbt.TagNodeCompound) *Husk {
	return &Husk{Mob: *MobFromNbt(node)}
}
