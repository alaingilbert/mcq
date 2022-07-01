package nbt

// TagNodeString ...
type TagNodeString struct {
	mData string
}

// NewTagNodeString ...
func NewTagNodeString(data string) *TagNodeString {
	return &TagNodeString{data}
}

// String ...
func (t *TagNodeString) String() string {
	return t.mData
}
