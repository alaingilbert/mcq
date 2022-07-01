package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type ElderGuardian struct{ Mob }

func ElderGuardianFromNbt(node *nbt.TagNodeCompound) *ElderGuardian {
	return &ElderGuardian{Mob: *MobFromNbt(node)}
}
