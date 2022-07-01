package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Axolotl struct{ Mob }

func AxolotlFromNbt(node *nbt.TagNodeCompound) *Axolotl {
	return &Axolotl{Mob: *MobFromNbt(node)}
}
