package nbt

// TagNodeShort ...
type TagNodeShort struct {
	Data int16
}

// NewTagNodeShort ...
func NewTagNodeShort(data int16) *TagNodeShort {
	return &TagNodeShort{data}
}

func (t TagNodeShort) Int() int { return int(t.Data) }
