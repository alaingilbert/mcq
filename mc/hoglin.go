package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Hoglin struct{ Mob }

func HoglinFromNbt(node *nbt.TagNodeCompound) *Hoglin {
	return &Hoglin{Mob: *MobFromNbt(node)}
}
