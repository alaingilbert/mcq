package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Chicken struct{ Mob }

func ChickenFromNbt(node *nbt.TagNodeCompound) *Chicken {
	return &Chicken{Mob: *MobFromNbt(node)}
}
