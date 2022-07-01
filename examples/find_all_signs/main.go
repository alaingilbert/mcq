package main

import (
	"fmt"
	"github.com/alaingilbert/mcq"
	"github.com/alaingilbert/mcq/mc"
	"os"
)

func main() {
	world := mcq.NewWorld(os.Getenv("WORLD_PATH"))
	mcq.Q(world).Targets(mc.SignID).Find(func(result mcq.Result) {
		if sign, ok := result.Item.(*mc.Sign); ok {
			if sign.HasText() {
				fmt.Printf("%s %v\n", result.Coord(), sign.InlineText())
			}
		}
	})
}
