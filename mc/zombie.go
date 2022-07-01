package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Zombie struct{ Mob }

func ZombieFromNbt(node *nbt.TagNodeCompound) *Zombie {
	return &Zombie{Mob: *MobFromNbt(node)}
}
