package mc

type Dimension int

func (d Dimension) String() string {
	switch d {
	case Overworld:
		return "Overworld"
	case Nether:
		return "Nether"
	case TheEnd:
		return "TheEnd"
	default:
		return "n/a"
	}
}

const (
	Overworld Dimension = iota + 1
	Nether
	TheEnd
)
