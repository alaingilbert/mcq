package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Chest struct{ ContainerBlockEntity }

func ChestFromNbt(node *nbt.TagNodeCompound) *Chest {
	return &Chest{ContainerBlockEntity: *ContainerBlockEntityFromNbt(node)}
}
