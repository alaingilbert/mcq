package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Item struct {
	BaseItem
	ItemID ID
}

func ItemFromNbt(node *nbt.TagNodeCompound) *Item {
	e := new(Item)
	item := node.Entries["Item"].(*nbt.TagNodeCompound)
	e.ItemID = ID(item.Entries["id"].(*nbt.TagNodeString).String())
	e.BaseItem = *baseItemFromNbt(node)
	return e
}
