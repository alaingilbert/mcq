package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Barrel struct{ ContainerBlockEntity }

func BarrelFromNbt(node *nbt.TagNodeCompound) *Barrel {
	return &Barrel{ContainerBlockEntity: *ContainerBlockEntityFromNbt(node)}
}
