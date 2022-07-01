package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Turtle struct{ Mob }

func TurtleFromNbt(node *nbt.TagNodeCompound) *Turtle {
	return &Turtle{Mob: *MobFromNbt(node)}
}
