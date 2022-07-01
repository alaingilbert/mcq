package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Skeleton struct{ Mob }

func SkeletonFromNbt(node *nbt.TagNodeCompound) *Skeleton {
	return &Skeleton{Mob: *MobFromNbt(node)}
}
