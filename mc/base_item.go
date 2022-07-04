package mc

import (
	"encoding/json"
	"github.com/alaingilbert/mcq/nbt"
)

type BaseItems []IItem

func (i BaseItems) Each(clb func(IItem)) {
	for _, el := range i {
		clb(el)
	}
}

type BaseItem struct {
	id      ID
	slot    byte
	count   byte
	display string
}

func (i BaseItem) ID() ID             { return i.id }
func (i BaseItem) Display() string    { return i.display }
func (i BaseItem) CustomName() string { return i.display }
func (i BaseItem) Slot() byte         { return i.slot }
func (i BaseItem) Count() byte        { return i.count }

//func NewBaseItem(id mc.ID, item *nbt.TagNodeCompound) *BaseItem {
//	return &BaseItem{id: id, nbt: item}
//}

func baseItemFromNbt(node *nbt.TagNodeCompound) *BaseItem {
	type txt struct {
		Text string `json:"text"`
	}
	i := new(BaseItem)
	i.id = ID(node.Entries["id"].(*nbt.TagNodeString).String())
	if tag, ok := node.Entries["tag"].(*nbt.TagNodeCompound); ok {
		if display, ok := tag.Entries["display"].(*nbt.TagNodeCompound); ok {
			name := display.Entries["Name"].(*nbt.TagNodeString).String()
			var t txt
			_ = json.Unmarshal([]byte(name), &t)
			i.display = t.Text
		}
	}
	if slot, ok := node.Entries["Slot"].(*nbt.TagNodeByte); ok {
		i.slot = slot.Byte()
	}
	if count, ok := node.Entries["Count"].(*nbt.TagNodeByte); ok {
		i.count = count.Byte()
	}
	return i
}

func parseItemFromNbt(node *nbt.TagNodeCompound) IItem {
	id := ID(node.Entries["id"].(*nbt.TagNodeString).String())
	var parsedItem IItem
	switch id {
	case ItemID:
		parsedItem = ItemFromNbt(node)
	case WritableBookID:
		parsedItem = WritableBookItemFromNbt(node)
	case WrittenBookID:
		parsedItem = WrittenBookItemFromNbt(node)
	case ShulkerBoxID:
		parsedItem = ShulkerBoxItemFromNbt(node)
	default:
		parsedItem = baseItemFromNbt(node)
	}
	return parsedItem
}
