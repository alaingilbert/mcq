package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type TraderLlama struct{ Mob }

func TraderLlamaFromNbt(node *nbt.TagNodeCompound) *TraderLlama {
	return &TraderLlama{Mob: *MobFromNbt(node)}
}
