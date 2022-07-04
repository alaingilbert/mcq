package main

import (
	"fmt"
	"github.com/alaingilbert/mcq"
	"github.com/alaingilbert/mcq/mc"
	"os"
)

func main() {
	world := mcq.NewWorld(os.Getenv("WORLD_PATH"))
	bbox := mcq.New2DBBox(mc.Overworld, 170, 160, 266, 226)
	mcq.Q(world).BBox(bbox).Targets(mc.ShulkerBoxID).Find(func(result mcq.Result) {
		if container, ok := result.Item.(mc.IContainerEntity); ok {
			fmt.Printf("SHULKER: %s %s %v\n", result.Coord(), result.Description, container)
		}
	})
}
