package nbt

// TagNodeIntArray ...
type TagNodeIntArray struct {
	data []int32
}

// NewTagNodeIntArray ...
func NewTagNodeIntArray(data []int32) *TagNodeIntArray {
	return &TagNodeIntArray{data}
}

// Data ...
func (t *TagNodeIntArray) Data() []int32 {
	return t.data
}
