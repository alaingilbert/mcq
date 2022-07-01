package nbt

// TagNodeByteArray ...
type TagNodeByteArray struct {
	mData []byte
}

// NewTagNodeByteArray ...
func NewTagNodeByteArray(data []byte) *TagNodeByteArray {
	return &TagNodeByteArray{data}
}

// Data ...
func (t *TagNodeByteArray) Data() []byte {
	return t.mData
}
