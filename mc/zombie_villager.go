package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type ZombieVillager struct{ Mob }

func ZombieVillagerFromNbt(node *nbt.TagNodeCompound) *ZombieVillager {
	return &ZombieVillager{Mob: *MobFromNbt(node)}
}
