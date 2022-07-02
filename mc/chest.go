package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type container struct {
	items BaseItems
}

func (c *container) Items() IItems { return c.items }

func ContainerFromNbt(node *nbt.TagNodeCompound) *container {
	c := new(container)
	c.items = make(BaseItems, 0)
	if items, ok := node.Entries["Items"].(*nbt.TagNodeList); ok {
		items.Each(func(node nbt.ITagNode) {
			item := node.(*nbt.TagNodeCompound)
			parsedItem := parseItemFromNbt(item)
			c.items = append(c.items, parsedItem)
		})
	}
	return c
}

type ContainerEntity struct {
	entity
	container
}

func ContainerEntityFromNbt(node *nbt.TagNodeCompound) *ContainerEntity {
	e := new(ContainerEntity)
	e.entity = *EntityFromNbt(node)
	e.container = *ContainerFromNbt(node)
	return e
}

type ContainerBlockEntity struct {
	blockEntity
	container
}

func ContainerBlockEntityFromNbt(node *nbt.TagNodeCompound) *ContainerBlockEntity {
	e := new(ContainerBlockEntity)
	e.blockEntity = *BlockEntityFromNbt(node)
	e.container = *ContainerFromNbt(node)
	return e
}

type BaseChest struct {
	ContainerBlockEntity
}

func BaseChestFromNbt(node *nbt.TagNodeCompound) *BaseChest {
	c := new(BaseChest)
	c.ContainerBlockEntity = *ContainerBlockEntityFromNbt(node)
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
