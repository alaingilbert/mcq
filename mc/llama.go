package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Llama struct{ Mob }

func LlamaFromNbt(node *nbt.TagNodeCompound) *Llama {
	return &Llama{Mob: *MobFromNbt(node)}
}
