package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Ocelot struct{ Mob }

func OcelotFromNbt(node *nbt.TagNodeCompound) *Ocelot {
	return &Ocelot{Mob: *MobFromNbt(node)}
}
