package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type MagmaCube struct{ Mob }

func MagmaCubeFromNbt(node *nbt.TagNodeCompound) *MagmaCube {
	return &MagmaCube{Mob: *MobFromNbt(node)}
}
