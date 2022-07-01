package mcq

import (
	"os"
)

type WorldCoordinate struct {
	x int
	y int
	z int
}

func (c WorldCoordinate) GetRegionCoordinate() (rx, rz int) {
	return 0, 0
}

// World represent a minecraft world.
type World struct {
	Path           string
	mRegionManager *RegionManager
}

// NewWorld instantiate a world object.
// It returns a pointer to a world.
func NewWorld(pPath string) *World {
	world := new(World)
	world.Path = pPath
	world.mRegionManager = NewRegionManager(world.Path)
	return world
}

// PathValid ...
func (w *World) PathValid() bool {
	_, err := os.Stat(w.Path)
	return err == nil
}

// RegionManager get the region manager.
// It returns a pointer to the region manager.
func (w *World) RegionManager() *RegionManager {
	return w.mRegionManager
}
