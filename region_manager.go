package mcq

import (
	"os"
	"path"
	"strconv"
	"strings"
)

// RegionDir name of the directory where region files are located
const RegionDir = "region"
const NetherDir = "DIM-1"
const TheEndDir = "DIM1"
const EntitiesDir = "entities"

// RegionManager is used to manage the regions.
type RegionManager struct {
	mWorldPath string
}

// NewRegionManager instantiate a new region manager.
// It returns a pointer to a region manager.
func NewRegionManager(pWorldPath string) *RegionManager {
	regionManager := &RegionManager{}
	regionManager.mWorldPath = pWorldPath
	return regionManager
}

// GetRegion get a specific region.
// pRegionX coordinate of the region on the X axis.
// pRegionZ coordinate of the region on the Z axis.
// It returns a pointer to a region.
func (r *RegionManager) GetRegion(dim Dimension, pRegionX, pRegionZ int) *Region {
	return NewRegion(r, dim, pRegionX, pRegionZ)
}

// GetRegionFromXYZ get a specific region from a global world coordinate.
// pX x world coordinate.
// pY y world coordinate.
// pZ z world coordinate.
// It returns a pointer to a region.
func (r *RegionManager) GetRegionFromXYZ(dim Dimension, pX, pY, pZ int) *Region {
	return NewRegionFromXYZ(r, dim, pX, pY, pZ)
}

func getFileNames(p string) []string {
	tilesDirectory, err := os.Open(p)
	if err != nil {
		return nil
	}
	defer tilesDirectory.Close()
	files, err := tilesDirectory.Readdirnames(0)
	var newFiles []string
	for _, file := range files {
		if !strings.HasSuffix(file, "mca") {
			continue
		}
		newFiles = append(newFiles, file)
	}
	return newFiles
}

func (r *RegionManager) Each(dim Dimension, clb func(region *Region)) {
	fNames := r.RegionFileNames(dim)
	for _, fName := range fNames {
		splits := strings.SplitN(fName, ".", 4)
		regionX, _ := strconv.Atoi(splits[1])
		regionZ, _ := strconv.Atoi(splits[2])
		region := r.GetRegion(dim, regionX, regionZ)
		clb(region)
	}
}

func (r *RegionManager) EachEntities(dim Dimension, clb func(region *Region)) {
	fNames := r.EntitiesFileNames(dim)
	for _, fName := range fNames {
		splits := strings.SplitN(fName, ".", 4)
		regionX, _ := strconv.Atoi(splits[1])
		regionZ, _ := strconv.Atoi(splits[2])
		region := r.GetRegion(dim, regionX, regionZ)
		clb(region)
	}
}

// RegionFileNames ...
func (r *RegionManager) RegionFileNames(dim Dimension) []string {
	var p string
	switch dim {
	case Overworld:
		p = path.Join(r.mWorldPath, RegionDir)
	case Nether:
		p = path.Join(r.mWorldPath, NetherDir, RegionDir)
	case TheEnd:
		p = path.Join(r.mWorldPath, TheEndDir, RegionDir)
	}
	return getFileNames(p)
}

func (r *RegionManager) EntitiesFileNames(dim Dimension) []string {
	var p string
	switch dim {
	case Overworld:
		p = path.Join(r.mWorldPath, EntitiesDir)
	case Nether:
		p = path.Join(r.mWorldPath, NetherDir, EntitiesDir)
	case TheEnd:
		p = path.Join(r.mWorldPath, TheEndDir, EntitiesDir)
	}
	return getFileNames(p)
}

// WorldPath get the world folder path.
func (r *RegionManager) WorldPath() string {
	return r.mWorldPath
}
