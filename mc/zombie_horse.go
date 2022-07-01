package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type ZombieHorse struct{ Mob }

func ZombieHorseFromNbt(node *nbt.TagNodeCompound) *ZombieHorse {
	return &ZombieHorse{Mob: *MobFromNbt(node)}
}
