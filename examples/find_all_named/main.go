package main

import (
	"fmt"
	"github.com/alaingilbert/mcq"
	"github.com/alaingilbert/mcq/mc"
	"os"
)

func main() {
	world := mcq.NewWorld(os.Getenv("WORLD_PATH"))
	mcq.Q(world).Find(func(result mcq.Result) {
		if item, ok := result.Item.(mc.IItem); ok {
			fmt.Printf("%s %v %s\n", result.Coord(), result, item.Display())
		} else if entity, ok := result.Item.(mc.IEntity); ok {
			fmt.Printf("%s %v %s\n", result.Coord(), result, entity.CustomName())
		}
	}, mcq.WithCustomName(true))
}
