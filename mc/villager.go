package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Villager struct{ Mob }

func VillagerFromNbt(node *nbt.TagNodeCompound) *Villager {
	return &Villager{Mob: *MobFromNbt(node)}
}
