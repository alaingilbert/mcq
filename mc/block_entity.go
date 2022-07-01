package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type blockEntity struct {
	id      ID
	x, y, z int
}

func BlockEntityFromNbt(node *nbt.TagNodeCompound) *blockEntity {
	e := new(blockEntity)
	e.id = ID(node.Entries["id"].(*nbt.TagNodeString).String())
	if x, ok := node.Entries["x"].(*nbt.TagNodeInt); ok {
		e.x = x.Int()
	}
	if y, ok := node.Entries["y"].(*nbt.TagNodeInt); ok {
		e.y = y.Int()
	}
	if z, ok := node.Entries["z"].(*nbt.TagNodeInt); ok {
		e.z = z.Int()
	}
	return e
}

func (e blockEntity) ID() ID { return e.id }
func (e blockEntity) X() int { return e.x }
func (e blockEntity) Y() int { return e.y }
func (e blockEntity) Z() int { return e.z }

func ParseBlockEntity(node *nbt.TagNodeCompound) (parsed IBlockEntity) {
	id := ID(node.Entries["id"].(*nbt.TagNodeString).String())
	switch id {
	case BarrelID:
		parsed = BarrelFromNbt(node)
	case LecternID:
		parsed = LecternFromNbt(node)
	case ChestID:
		parsed = ChestFromNbt(node)
	case TrappedChestID:
		parsed = TrappedChestFromNbt(node)
	case ShulkerBoxID:
		parsed = ShulkerBoxFromNbt(node)
	case SignID:
		parsed = SignFromNbt(node)
	default:
		parsed = BlockEntityFromNbt(node)
	}
	return
}
