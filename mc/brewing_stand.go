package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type BrewingStand struct {
	ContainerBlockEntity
	BrewTime         int // The number of ticks the potions have to brew.
	Fuel             int // Remaining fuel for the brewing stand. 20 when full, and counts down by 1 each time a potion is brewed.
	LeftPotionSlot   IItem
	MiddlePotionSlot IItem
	RightPotionSlot  IItem
	IngredientSlot   IItem
	FuelSlot         IItem
}

func BrewingStandFromNbt(node *nbt.TagNodeCompound) *BrewingStand {
	brewingStand := new(BrewingStand)
	brewingStand.blockEntity = *BlockEntityFromNbt(node)
	brewingStand.BrewTime = node.Entries["BrewTime"].(*nbt.TagNodeShort).Int()
	brewingStand.Fuel = node.Entries["Fuel"].(*nbt.TagNodeByte).Int()
	brewingStand.items = make(BaseItems, 0)
	if items, ok := node.Entries["Items"].(*nbt.TagNodeList); ok {
		items.Each(func(node nbt.ITagNode) {
			item := node.(*nbt.TagNodeCompound)
			parsedItem := parseItemFromNbt(item)
			if parsedItem.Slot() == 0 {
				brewingStand.LeftPotionSlot = parsedItem
			} else if parsedItem.Slot() == 1 {
				brewingStand.MiddlePotionSlot = parsedItem
			} else if parsedItem.Slot() == 2 {
				brewingStand.RightPotionSlot = parsedItem
			} else if parsedItem.Slot() == 3 {
				brewingStand.IngredientSlot = parsedItem
			} else if parsedItem.Slot() == 4 {
				brewingStand.FuelSlot = parsedItem
			}
			brewingStand.items = append(brewingStand.items, parsedItem)
		})
	}
	return brewingStand
}
