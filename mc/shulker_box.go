package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type ShulkerBox struct {
	blockEntity
	items BaseItems
}

func ShulkerBoxFromNbt(node *nbt.TagNodeCompound) *ShulkerBox {
	s := new(ShulkerBox)
	s.blockEntity = *BlockEntityFromNbt(node)
	s.items = make(BaseItems, 0)
	if items, ok := node.Entries["Items"].(*nbt.TagNodeList); ok {
		items.Each(func(node nbt.ITagNode) {
			item := node.(*nbt.TagNodeCompound)
			parsedItem := parseItemFromNbt(item)
			s.items = append(s.items, parsedItem)
		})
	}
	return s
}

func (s *ShulkerBox) Items() IItems { return s.items }

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
