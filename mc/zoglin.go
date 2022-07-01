package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Zoglin struct{ Mob }

func ZoglinFromNbt(node *nbt.TagNodeCompound) *Zoglin {
	return &Zoglin{Mob: *MobFromNbt(node)}
}
