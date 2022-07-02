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
		if i, ok := result.Item.(mc.INamed); ok {
			fmt.Printf("%s %v %s\n", result.Coord(), result, i.CustomName())
		}
	}, mcq.WithCustomName(true))
}
