package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Blaze struct{ Mob }

func BlazeFromNbt(node *nbt.TagNodeCompound) *Blaze {
	return &Blaze{Mob: *MobFromNbt(node)}
}
