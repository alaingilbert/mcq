package nbt

// TagNodeFloat ...
type TagNodeFloat struct {
	mData float32
}

// NewTagNodeFloat ...
func NewTagNodeFloat(data float32) *TagNodeFloat {
	return &TagNodeFloat{data}
}

func (n *TagNodeFloat) Float32() float32 { return n.mData }
