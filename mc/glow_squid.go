package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type GlowSquid struct{ Mob }

func GlowSquidFromNbt(node *nbt.TagNodeCompound) *GlowSquid {
	return &GlowSquid{Mob: *MobFromNbt(node)}
}
