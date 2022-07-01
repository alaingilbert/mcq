package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type TropicalFish struct{ Mob }

func TropicalFishFromNbt(node *nbt.TagNodeCompound) *TropicalFish {
	return &TropicalFish{Mob: *MobFromNbt(node)}
}
