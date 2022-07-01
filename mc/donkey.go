package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Donkey struct{ Mob }

func DonkeyFromNbt(node *nbt.TagNodeCompound) *Donkey {
	return &Donkey{Mob: *MobFromNbt(node)}
}
