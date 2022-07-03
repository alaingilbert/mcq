package mc

type Block struct {
	id ID
	Coordinate
}

func NewBlock(id ID, coord Coordinate) *Block {
	return &Block{id: id, Coordinate: coord}
}

func (b Block) ID() ID { return b.id }
