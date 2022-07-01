package nbt

// TagNodeLong ...
type TagNodeLong struct {
	Data int64
}

// NewTagNodeLong ...
func NewTagNodeLong(data int64) *TagNodeLong {
	return &TagNodeLong{data}
}
