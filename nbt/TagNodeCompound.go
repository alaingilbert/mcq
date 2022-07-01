package nbt

// TagNodeCompound ...
type TagNodeCompound struct {
	Entries map[string]ITagNode
}

// NewTagNodeCompound ...
// pEntries a map of name -> tag
// It returns a pointer to the tagNodeCompound.
func NewTagNodeCompound(pEntries map[string]ITagNode) *TagNodeCompound {
	return &TagNodeCompound{pEntries}
}
