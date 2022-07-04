package main

import (
	"fmt"
	"github.com/alaingilbert/mcq"
	"github.com/alaingilbert/mcq/mc"
	"os"
	"sort"
)

func main() {
	world := mcq.NewWorld(os.Getenv("WORLD_PATH"))

	type BlockStat struct {
		ID    mc.ID
		Count int
	}

	// Block that we want to count
	targets := []string{mc.DarkOakLeavesID, mc.SpruceLogID, mc.BirchLeavesID, mc.SpruceWoodID}
	// 3D area in which we want to count blocks
	bbox := mcq.New3DBBox(mc.Overworld, 0, 0, 0, 50, 50, 50)
	statsMap := make(map[mc.ID]int)
	mcq.Q(world).BBox(bbox).Targets(targets...).Find(func(result mcq.Result) {
		statsMap[result.Item.ID()]++
	}, mcq.WithBlocks)

	// Map to slice and sort
	statsArr := make([]BlockStat, 0)
	for blockID, blockCount := range statsMap {
		statsArr = append(statsArr, BlockStat{blockID, blockCount})
	}
	sort.Slice(statsArr, func(i, j int) bool { return statsArr[i].Count > statsArr[j].Count })

	// Display sorted result
	for _, el := range statsArr {
		fmt.Printf("%30s -> %d\n", el.ID, el.Count)
	}
}
