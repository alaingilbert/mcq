package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type ShulkerBox struct{ ContainerBlockEntity }

func ShulkerBoxFromNbt(node *nbt.TagNodeCompound) *ShulkerBox {
	return &ShulkerBox{ContainerBlockEntity: *ContainerBlockEntityFromNbt(node)}
}

type ShulkerBoxItem struct {
	BaseItem
	ShulkerBox ShulkerBox
}

func ShulkerBoxItemFromNbt(node *nbt.TagNodeCompound) *ShulkerBoxItem {
	i := new(ShulkerBoxItem)
	i.BaseItem = *baseItemFromNbt(node)
	if tag, ok := node.Entries["tag"].(*nbt.TagNodeCompound); ok {
		blockEntityTag := tag.Entries["BlockEntityTag"].(*nbt.TagNodeCompound)
		i.ShulkerBox = *ShulkerBoxFromNbt(blockEntityTag)
	}
	return i
}

// Items implements IContainerEntity
func (s ShulkerBoxItem) Items() IItems { return s.ShulkerBox.items }
