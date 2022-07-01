package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type ZombifiedPiglin struct{ Mob }

func ZombifiedPiglinFromNbt(node *nbt.TagNodeCompound) *ZombifiedPiglin {
	return &ZombifiedPiglin{Mob: *MobFromNbt(node)}
}
