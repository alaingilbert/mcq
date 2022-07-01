package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Boat struct{ vehicle }

func BoatFromNbt(node *nbt.TagNodeCompound) *Boat {
	return &Boat{vehicle: *VehicleFromNbt(node)}
}
