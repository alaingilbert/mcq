package mc

type Coordinate struct {
	dim     Dimension
	x, y, z int
}

func NewCoord(dim Dimension, x, y, z int) *Coordinate {
	return &Coordinate{dim: dim, x: x, y: y, z: z}
}

func NewCoordO(x, y, z int) *Coordinate { return NewCoord(Overworld, x, y, z) }
func NewCoordN(x, y, z int) *Coordinate { return NewCoord(Nether, x, y, z) }
func NewCoordE(x, y, z int) *Coordinate { return NewCoord(TheEnd, x, y, z) }

func (c Coordinate) Dim() Dimension                     { return c.dim }
func (c Coordinate) X() int                             { return c.x }
func (c Coordinate) Y() int                             { return c.y }
func (c Coordinate) Z() int                             { return c.z }
func (c Coordinate) Unpack() (Dimension, int, int, int) { return c.dim, c.x, c.y, c.z }
