package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type PolarBear struct{ Mob }

func PolarBearFromNbt(node *nbt.TagNodeCompound) *PolarBear {
	return &PolarBear{Mob: *MobFromNbt(node)}
}
