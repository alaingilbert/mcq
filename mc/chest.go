package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type ContainerEntity struct {
	blockEntity
	items BaseItems
}

func (c *ContainerEntity) Items() IItems { return c.items }

func ContainerEntityFromNbt(node *nbt.TagNodeCompound) *ContainerEntity {
	e := new(ContainerEntity)
	e.blockEntity = *BlockEntityFromNbt(node)
	e.items = make(BaseItems, 0)
	if items, ok := node.Entries["Items"].(*nbt.TagNodeList); ok {
		items.Each(func(node nbt.ITagNode) {
			item := node.(*nbt.TagNodeCompound)
			parsedItem := parseItemFromNbt(item)
			e.items = append(e.items, parsedItem)
		})
	}
	return e
}

type BaseChest struct {
	ContainerEntity
	customName string
}

func BaseChestFromNbt(node *nbt.TagNodeCompound) *BaseChest {
	c := new(BaseChest)
	c.ContainerEntity = *ContainerEntityFromNbt(node)
	if customName, ok := node.Entries["CustomName"].(*nbt.TagNodeString); ok {
		c.customName = customName.String()
	}
	return c
}

func (c *BaseChest) CustomName() string { return c.customName }

type Chest struct {
	BaseChest
}

func ChestFromNbt(node *nbt.TagNodeCompound) *Chest {
	c := new(Chest)
	c.BaseChest = *BaseChestFromNbt(node)
	return c
}
