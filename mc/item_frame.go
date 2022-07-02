package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type BaseItemFrame struct {
	entity
	item IItem
}

func (e BaseItemFrame) Item() IItem { return e.item }

func BaseItemFrameFromNbt(node *nbt.TagNodeCompound) *BaseItemFrame {
	i := new(BaseItemFrame)
	i.entity = *EntityFromNbt(node)
	if item, ok := node.Entries["Item"].(*nbt.TagNodeCompound); ok {
		parsed := parseItemFromNbt(item)
		i.item = parsed
	}
	return i
}

type ItemFrame struct{ BaseItemFrame }

func ItemFrameFromNbt(node *nbt.TagNodeCompound) *ItemFrame {
	return &ItemFrame{BaseItemFrame: *BaseItemFrameFromNbt(node)}
}

type GlowItemFrame struct{ BaseItemFrame }

func GlowItemFrameFromNbt(node *nbt.TagNodeCompound) *GlowItemFrame {
	return &GlowItemFrame{BaseItemFrame: *BaseItemFrameFromNbt(node)}
}
