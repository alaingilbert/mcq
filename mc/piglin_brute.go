package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type PiglinBrute struct{ Mob }

func PiglinBruteFromNbt(node *nbt.TagNodeCompound) *PiglinBrute {
	return &PiglinBrute{Mob: *MobFromNbt(node)}
}
