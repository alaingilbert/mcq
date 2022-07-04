package nbt

// TagNodeByte ...
type TagNodeByte struct {
	mData byte
}

// NewTagNodeByte ...
func NewTagNodeByte(data byte) *TagNodeByte {
	return &TagNodeByte{data}
}

func (t TagNodeByte) Byte() byte { return t.mData }

func (t TagNodeByte) Int() int { return int(t.mData) }
