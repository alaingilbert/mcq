package nbt

import "fmt"

// TagNodeInt ...
type TagNodeInt struct {
	mData int32
}

// NewTagNodeInt ...
func NewTagNodeInt(data int32) *TagNodeInt {
	return &TagNodeInt{data}
}

func (t *TagNodeInt) Int() int {
	return int(t.mData)
}

func (t *TagNodeInt) String() string {
	return fmt.Sprintf("%d", t.mData)
}
