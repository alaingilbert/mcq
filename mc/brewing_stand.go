package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type BrewingStand struct {
	ContainerBlockEntity
	brewTime         int // The number of ticks the potions have to brew.
	fuel             int // Remaining fuel for the brewing stand. 20 when full, and counts down by 1 each time a potion is brewed.
	leftPotionSlot   IItem
	middlePotionSlot IItem
	rightPotionSlot  IItem
	ingredientSlot   IItem
	fuelSlot         IItem
}

func (b BrewingStand) BrewTime() int           { return b.brewTime }
func (b BrewingStand) Fuel() int               { return b.fuel }
func (b BrewingStand) LeftPotionSlot() IItem   { return b.leftPotionSlot }
func (b BrewingStand) MiddlePotionSlot() IItem { return b.middlePotionSlot }
func (b BrewingStand) RightPotionSlot() IItem  { return b.rightPotionSlot }
func (b BrewingStand) IngredientSlot() IItem   { return b.ingredientSlot }
func (b BrewingStand) FuelSlot() IItem         { return b.fuelSlot }

func BrewingStandFromNbt(node *nbt.TagNodeCompound) *BrewingStand {
	brewingStand := new(BrewingStand)
	brewingStand.blockEntity = *BlockEntityFromNbt(node)
	brewingStand.brewTime = node.Entries["BrewTime"].(*nbt.TagNodeShort).Int()
	brewingStand.fuel = node.Entries["Fuel"].(*nbt.TagNodeByte).Int()
	brewingStand.items = make(BaseItems, 0)
	if items, ok := node.Entries["Items"].(*nbt.TagNodeList); ok {
		items.Each(func(node nbt.ITagNode) {
			item := node.(*nbt.TagNodeCompound)
			parsedItem := parseItemFromNbt(item)
			if parsedItem.Slot() == 0 {
				brewingStand.leftPotionSlot = parsedItem
			} else if parsedItem.Slot() == 1 {
				brewingStand.middlePotionSlot = parsedItem
			} else if parsedItem.Slot() == 2 {
				brewingStand.rightPotionSlot = parsedItem
			} else if parsedItem.Slot() == 3 {
				brewingStand.ingredientSlot = parsedItem
			} else if parsedItem.Slot() == 4 {
				brewingStand.fuelSlot = parsedItem
			}
			brewingStand.items = append(brewingStand.items, parsedItem)
		})
	}
	return brewingStand
}
