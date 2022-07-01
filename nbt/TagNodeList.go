package nbt

// TagNodeList ...
type TagNodeList struct {
	mType   TagType
	mLength int32
	mList   []ITagNode
}

func NewTagNodeList(pType TagType, pLength int32, pList []ITagNode) *TagNodeList {
	return &TagNodeList{pType, pLength, pList}
}

// Length ...
func (t *TagNodeList) Length() int32 {
	return t.mLength
}

// Get ...
func (t *TagNodeList) Get(i int) ITagNode {
	return t.mList[i]
}

// Add ...
func (t *TagNodeList) Add(item ITagNode, i int) {
	t.mList[i] = item
}

// ValueType ...
func (t *TagNodeList) ValueType() TagType {
	return t.mType
}

// Each ...
func (t *TagNodeList) Each(clb func(node ITagNode)) {
	for i := 0; i < int(t.Length()); i++ {
		clb(t.Get(i))
	}
}
