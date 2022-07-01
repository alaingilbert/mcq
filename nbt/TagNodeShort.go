package nbt

// TagNodeShort ...
type TagNodeShort struct {
	Data int16
}

// NewTagNodeShort ...
func NewTagNodeShort(data int16) *TagNodeShort {
	return &TagNodeShort{data}
}
