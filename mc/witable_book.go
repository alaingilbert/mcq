package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type WritableBook struct {
	BaseBook
}

func NewWritableBook() *WritableBook {
	b := new(WritableBook)
	b.pages = make([]string, 0)
	return b
}

func WritableBookFromNbt(node *nbt.TagNodeCompound) *WritableBook {
	b := NewWritableBook()
	pages := node.Entries["pages"].(*nbt.TagNodeList)
	pages.Each(func(page nbt.ITagNode) {
		b.pages = append(b.pages, page.(*nbt.TagNodeString).String())
	})
	return b
}
