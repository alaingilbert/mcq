package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type SkeletonHorse struct{ Mob }

func SkeletonHorseFromNbt(node *nbt.TagNodeCompound) *SkeletonHorse {
	return &SkeletonHorse{Mob: *MobFromNbt(node)}
}
