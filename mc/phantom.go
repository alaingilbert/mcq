package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Phantom struct{ Mob }

func PhantomFromNbt(node *nbt.TagNodeCompound) *Phantom {
	return &Phantom{Mob: *MobFromNbt(node)}
}
