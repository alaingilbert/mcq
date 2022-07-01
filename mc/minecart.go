package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Minecart struct{ vehicle }

func MinecartFromNbt(node *nbt.TagNodeCompound) *Minecart {
	return &Minecart{vehicle: *VehicleFromNbt(node)}
}
