package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Rabbit struct{ Mob }

func RabbitFromNbt(node *nbt.TagNodeCompound) *Rabbit {
	return &Rabbit{Mob: *MobFromNbt(node)}
}
