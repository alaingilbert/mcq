package main

import (
	"fmt"
	"github.com/alaingilbert/mcq"
	"github.com/alaingilbert/mcq/mc"
	"os"
)

func main() {
	world := mcq.NewWorld(os.Getenv("WORLD_PATH"))
	bbox := mcq.New2DBBox(mc.Overworld, 0, 0, 100, 100)
	mcq.Q(world).BBox(bbox).Targets(mc.DiamondOreID).Find(func(result mcq.Result) {
		fmt.Printf("Found diamond ore at %s\n", result.Coord())
	}, mcq.WithBlocks)
}
