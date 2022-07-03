package mcq

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"fmt"
	"github.com/alaingilbert/mcq/mc"
	"github.com/alaingilbert/mcq/nbt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path"
)

// Region represent a minecraft region.
type Region struct {
	x, z          int
	dim           mc.Dimension
	regionManager *RegionManager
}

// NewRegion instantiate a Region.
// regionManager pointer to the region manager who is calling the function.
// regionX region X axis.
// regionZ region Z axis.
// It returns a pointer to the region.
func NewRegion(regionManager *RegionManager, dim mc.Dimension, regionX, regionZ int) *Region {
	region := new(Region)
	region.regionManager = regionManager
	region.x = regionX
	region.z = regionZ
	region.dim = dim
	return region
}

// NewRegionFromXYZ ...
func NewRegionFromXYZ(regionManager *RegionManager, dim mc.Dimension, x, y, z int) *Region {
	regionX, regionZ := regionCoordinatesFromXYZ(x, y, z)
	return NewRegion(regionManager, dim, regionX, regionZ)
}

// RegionCoordinatesFromWorldXZ right shift 4 means divided by 16
func RegionCoordinatesFromWorldXZ(worldX, worldZ int) (int, int) {
	chunkX := worldX >> 4
	chunkZ := worldZ >> 4
	return regionCoordinatesFromChunkXZ(chunkX, chunkZ)
}

// regionCoordinatesFromChunkXZ right shift 5 means divided by 32
func regionCoordinatesFromChunkXZ(chunkX, chunkZ int) (int, int) {
	regionX := chunkX >> 5
	regionZ := chunkZ >> 5
	return regionX, regionZ
}

// regionCoordinatesFromXYZ ...
func regionCoordinatesFromXYZ(x, y, z int) (int, int) {
	var regionX = int(math.Floor(float64(x) / (math.Pow(2, float64(z)))))
	var regionZ = int(math.Floor(float64(y) / (math.Pow(2, float64(z)))))
	return regionX, regionZ
}

func (r *Region) Each(clb func(chunk *Chunk)) {
	var p string
	switch r.dim {
	case mc.Overworld:
		p = path.Join(r.regionManager.WorldPath(), RegionDir, r.FileName())
	case mc.Nether:
		p = path.Join(r.regionManager.WorldPath(), NetherDir, RegionDir, r.FileName())
	case mc.TheEnd:
		p = path.Join(r.regionManager.WorldPath(), TheEndDir, RegionDir, r.FileName())
	}

	f, err := os.Open(p)
	if err != nil {
		return
	}
	defer f.Close()

	by, _ := ioutil.ReadAll(f)
	b := bytes.NewReader(by)

	for cx := 0; cx < 32; cx++ {
		for cz := 0; cz < 32; cz++ {
			chunk := getChunkFromReadSeeker(b, r.x, r.z, cx, cz)
			if chunk == nil {
				continue
			}
			clb(chunk)
		}
	}
}

func (r *Region) EachEntities(clb func(chunk *Chunk)) {
	var p string
	switch r.dim {
	case mc.Overworld:
		p = path.Join(r.regionManager.WorldPath(), EntitiesDir, r.FileName())
	case mc.Nether:
		p = path.Join(r.regionManager.WorldPath(), NetherDir, EntitiesDir, r.FileName())
	case mc.TheEnd:
		p = path.Join(r.regionManager.WorldPath(), TheEndDir, EntitiesDir, r.FileName())
	}

	f, err := os.Open(p)
	if err != nil {
		return
	}
	defer f.Close()

	by, _ := ioutil.ReadAll(f)
	b := bytes.NewReader(by)

	for cx := 0; cx < 32; cx++ {
		for cz := 0; cz < 32; cz++ {
			chunk := getChunkFromReadSeeker(b, r.x, r.z, cx, cz)
			if chunk == nil {
				continue
			}
			clb(chunk)
		}
	}
}

// FileName get the file name for the region.
// It returns the file name for the region.
func (r *Region) FileName() string {
	return fmt.Sprintf("r.%d.%d.mca", r.x, r.z)
}

// FilePath get the file path.
// It returns the file path.
func (r *Region) FilePath() string {
	return path.Join(r.regionManager.WorldPath(), RegionDir)
}

// Exists ...
func (r *Region) Exists() bool {
	fPath := path.Join(r.FilePath(), r.FileName())
	_, err := os.Stat(fPath)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func (r Region) GetX() int {
	return r.x
}

func (r Region) GetZ() int {
	return r.z
}

const (
	gzipCompression = 1
	zlibCompression = 2
	uncompressed    = 3
)

// https://minecraft.fandom.com/wiki/Region_file_format
func processRMCAFile(p string, regionX, regionZ, chunkX, chunkZ int) *Chunk {
	f, err := os.Open(p)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer f.Close()

	return getChunkFromReadSeeker(f, regionX, regionZ, chunkX, chunkZ)
}

func getChunkFromReadSeeker(f io.ReadSeeker, regionX, regionZ, chunkX, chunkZ int) *Chunk {
	location := chunkHeaderOffset(chunkX, chunkZ)

	// Move to chunk header location
	_, _ = f.Seek(int64(location), 0)

	// First 3 bytes is the offset where the chunk data is in the file
	b := make([]byte, 3)
	_, _ = f.Read(b)
	offset := int64(b[2]) | int64(b[1])<<8 | int64(b[0])<<16

	// No offset means the chunk does not exist
	if offset == 0 {
		return nil
	}

	// Move to chunk data location
	_, _ = f.Seek(offset*4096, 0)

	// Get length in bytes
	lengthBytes := make([]byte, 4)
	_, _ = f.Read(lengthBytes)
	length := int64(binary.BigEndian.Uint32(lengthBytes))

	// Get compression type (1: GZip, 2: Zlib, 3: uncompressed)
	compressionByte := make([]byte, 1)
	_, _ = f.Read(compressionByte)
	compression := compressionByte[0]

	// Get compressed data
	compressedData := make([]byte, length-1)
	_, _ = f.Read(compressedData)

	if compression == zlibCompression {
		reader, err := zlib.NewReader(bytes.NewReader(compressedData))
		if err != nil {
			log.Fatal(err)
		}
		defer reader.Close()

		buf := new(bytes.Buffer)
		_, _ = buf.ReadFrom(reader)
		tree := nbt.NewNbtTree(buf)
		chunk := NewChunk(regionX, regionZ, chunkX, chunkZ)
		chunk.SetData(tree)
		return chunk
	}
	return nil
}

func (r *Region) GetDimension() mc.Dimension {
	return r.dim
}

func (r *Region) GetChunkFromWorldXZ(worldX, worldZ int) *Chunk {
	cx := worldX >> 4
	cz := worldZ >> 4
	return r.GetChunk(r.dim, cx, cz)
}

func (r *Region) GetEntities(localX, localZ int) *Chunk {
	var p string
	switch r.dim {
	case mc.Overworld:
		p = path.Join(r.regionManager.WorldPath(), EntitiesDir, r.FileName())
	case mc.Nether:
		p = path.Join(r.regionManager.WorldPath(), NetherDir, EntitiesDir, r.FileName())
	case mc.TheEnd:
		p = path.Join(r.regionManager.WorldPath(), TheEndDir, EntitiesDir, r.FileName())
	}
	return processRMCAFile(p, r.x, r.z, localX, localZ)
}

// GetChunk get the information for a specific chunk.
// localX X position of the chunk in the region.
// localZ Z position of the chunk in the region.
// It returns a pointer to the chunk.
func (r *Region) GetChunk(dimension mc.Dimension, localX, localZ int) *Chunk {
	var p string
	switch dimension {
	case mc.Overworld:
		p = path.Join(r.regionManager.WorldPath(), RegionDir, r.FileName())
	case mc.Nether:
		p = path.Join(r.regionManager.WorldPath(), NetherDir, RegionDir, r.FileName())
	case mc.TheEnd:
		p = path.Join(r.regionManager.WorldPath(), TheEndDir, RegionDir, r.FileName())
	}
	return processRMCAFile(p, r.x, r.z, localX, localZ)
}

// chunkHeaderOffset get the offset of the chunk information in the file header.
// It returns the offset in bytes.
// & 0b1_1111 ensure the value is always in the range [0-31]
func chunkHeaderOffset(chunkX, chunkZ int) int {
	return ((chunkX & 0b1_1111) + (chunkZ&0b1_1111)*32) * 4
}
