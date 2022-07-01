package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Horse struct{ Mob }

func HorseFromNbt(node *nbt.TagNodeCompound) *Horse {
	return &Horse{Mob: *MobFromNbt(node)}
}
