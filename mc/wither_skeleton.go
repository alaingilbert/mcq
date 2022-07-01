package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type WitherSkeleton struct{ Mob }

func WitherSkeletonFromNbt(node *nbt.TagNodeCompound) *WitherSkeleton {
	return &WitherSkeleton{Mob: *MobFromNbt(node)}
}
