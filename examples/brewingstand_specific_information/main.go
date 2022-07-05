package main

import (
	"fmt"
	"github.com/alaingilbert/mcq"
	"github.com/alaingilbert/mcq/mc"
	"os"
)

func main() {
	world := mcq.NewWorld(os.Getenv("WORLD_PATH"))
	mcq.Q(world).Targets(mc.BrewingStandID).Find(func(result mcq.Result) {
		fmt.Println("----------------------------------------")
		fmt.Printf("BrewingStand at %s\n", result.Coord())
		fmt.Println("------")

		// We can handle the BrewingStand as a generic IContainerEntity and iterate all items
		// like we would a chest or a shulkerbox.
		if s, ok := result.Item.(mc.IContainerEntity); ok {
			s.Items().Each(func(item mc.IItem) {
				fmt.Printf("Slot: %d | Count: %d | %s\n", item.Slot(), item.Count(), item.ID())
			})
		}

		fmt.Println("------")

		// Or we can handle it as a *BrewingStand and access specific information such as Fuel or BrewingTime.
		if s, ok := result.Item.(*mc.BrewingStand); ok {
			displaySlot := func(i mc.IItem) string {
				if i != nil {
					return fmt.Sprintf("%s %d", i.ID(), i.Count())
				}
				return "<empty>"
			}
			fmt.Printf("BrewingStand: Brew-Time: %d | Fuel: %d\n", s.BrewTime(), s.Fuel())
			fmt.Printf("Left potion slot:   %s\n", displaySlot(s.LeftPotionSlot()))
			fmt.Printf("Middle potion slot: %s\n", displaySlot(s.MiddlePotionSlot()))
			fmt.Printf("Right potion slot:  %s\n", displaySlot(s.RightPotionSlot()))
			fmt.Printf("Ingredient slot:    %s\n", displaySlot(s.IngredientSlot()))
			fmt.Printf("Fuel slot:          %s\n", displaySlot(s.FuelSlot()))
		}
		fmt.Println("----------------------------------------")
	})
}
