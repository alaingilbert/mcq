package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Spider struct{ Mob }

func SpiderFromNbt(node *nbt.TagNodeCompound) *Spider {
	return &Spider{Mob: *MobFromNbt(node)}
}
