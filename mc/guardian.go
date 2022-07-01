package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Guardian struct{ Mob }

func GuardianFromNbt(node *nbt.TagNodeCompound) *Guardian {
	return &Guardian{Mob: *MobFromNbt(node)}
}
