package mcq

import (
	"github.com/alaingilbert/mcq/mc"
	"github.com/alaingilbert/mcq/nbt"
	"math/bits"
)

const NbSection = 16
const SectionHeight = 16
const ChunkXDim = 16
const ChunkYDim = 256 // 16 sections of 16 blocks
const ChunkZDim = 16
const RegionWidth = 512 // 32 chunks of 16 blocks

// Chunk ...
type Chunk struct {
	dim              mc.Dimension
	regionX, regionZ int
	localX, localZ   int
	data             *nbt.NbtTree
}

// NewChunk ...
func NewChunk(dim mc.Dimension, regionX, regionZ, localX, localZ int) *Chunk {
	chunk := new(Chunk)
	chunk.localX = localX
	chunk.localZ = localZ
	chunk.dim = dim
	chunk.regionX = regionX
	chunk.regionZ = regionZ
	return chunk
}

// GetWorldX relative to world
func (c *Chunk) GetWorldX() int { return c.localX<<4 + c.regionX<<9 }

// GetWorldZ relative to world
func (c *Chunk) GetWorldZ() int { return c.localZ<<4 + c.regionZ<<9 }

// GetX relative to the region (0 to 31)
func (c *Chunk) GetX() int { return c.localX }

// GetZ relative to the region (0 to 31)
func (c *Chunk) GetZ() int { return c.localZ }

func (c *Chunk) GetData() (data *nbt.NbtTree) {
	return c.data
}

// SetData ...
func (c *Chunk) SetData(data *nbt.NbtTree) {
	c.data = data
}

// Each iterates all blocks in a chunk
func (c *Chunk) Each(clb func(block mc.Block)) {
	coordFromPos := func(section, blockPos int) (x int, y int, z int) {
		y = blockPos / ChunkYDim
		z = blockPos % ChunkYDim / ChunkXDim
		x = blockPos % ChunkXDim
		y += section * SectionHeight
		y -= 64
		x += c.GetWorldX()
		z += c.GetWorldZ()
		return
	}
	sections := c.GetData().Root().Entries["sections"].(*nbt.TagNodeList)
	for s := 0; s < NbSection; s++ {
		section := sections.Get(s).(*nbt.TagNodeCompound)
		blockStates := section.Entries["block_states"].(*nbt.TagNodeCompound)
		palette := blockStates.Entries["palette"].(*nbt.TagNodeList)
		if palette.Length() == 1 {
			blockID := mc.ID(palette.Get(0).(*nbt.TagNodeCompound).Entries["Name"].(*nbt.TagNodeString).String())
			for blockPos := 0; blockPos < ChunkXDim*ChunkZDim*SectionHeight; blockPos++ {
				x, y, z := coordFromPos(s, blockPos)
				coord := *mc.NewCoord(c.dim, x, y, z)
				block := *mc.NewBlock(blockID, coord)
				clb(block)
			}
			continue
		}
		data := blockStates.Entries["data"].(*nbt.TagNodeLongArray)
		mask := uint8(0b1111)
		if palette.Length() > 64 {
			mask = 0b111_1111
		} else if palette.Length() > 32 {
			mask = 0b11_1111
		} else if palette.Length() > 16 {
			mask = 0b1_1111
		}
		ones := bits.OnesCount8(mask)
		for blockPos := 0; blockPos < ChunkXDim*ChunkZDim*SectionHeight; blockPos++ {
			blockLngIdx := blockPos / (64 / ones)
			lng := data.Data()[blockLngIdx]
			indexRemaining := blockPos % (64 / ones)
			blockPaletteIndex := int(uint8(lng>>(indexRemaining*ones)) & mask)
			blockID := mc.ID(palette.Get(blockPaletteIndex).(*nbt.TagNodeCompound).Entries["Name"].(*nbt.TagNodeString).String())
			x, y, z := coordFromPos(s, blockPos)
			coord := *mc.NewCoord(c.dim, x, y, z)
			block := *mc.NewBlock(blockID, coord)
			clb(block)
		}
	}
}
