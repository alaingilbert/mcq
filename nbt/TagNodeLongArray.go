package nbt

// TagNodeLongArray ...
type TagNodeLongArray struct {
	data []int64
}

// NewTagNodeLongArray ...
func NewTagNodeLongArray(data []int64) *TagNodeLongArray {
	return &TagNodeLongArray{data}
}

// Data ...
func (t *TagNodeLongArray) Data() []int64 {
	return t.data
}
