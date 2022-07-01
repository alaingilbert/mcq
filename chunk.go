package mcq

import (
	"github.com/alaingilbert/mcq/nbt"
)

const NbSection int = 16
const SectionHeight int = 16
const XDim int = 16
const YDim int = 256
const ZDim int = 16

// Chunk ...
type Chunk struct {
	localX, localZ int
	data           *nbt.NbtTree
}

// NewChunk ...
func NewChunk(localX, localZ int) *Chunk {
	chunk := new(Chunk)
	chunk.localX = localX
	chunk.localZ = localZ
	return chunk
}

func (c *Chunk) GetX() int {
	return c.localX
}

func (c *Chunk) GetZ() int {
	return c.localZ
}

func (c *Chunk) GetData() (data *nbt.NbtTree) {
	return c.data
}

// SetData ...
func (c *Chunk) SetData(data *nbt.NbtTree) {
	c.data = data
}
