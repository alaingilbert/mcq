package mc

import (
	"encoding/json"
	"github.com/alaingilbert/mcq/nbt"
)

type BaseBook struct {
	pages []string
}

func BaseBookFromNbt(node *nbt.TagNodeCompound) *BaseBook {
	b := new(BaseBook)
	return b
}

func (b BaseBook) Pages() []string { return b.pages }

type WrittenBook struct {
	BaseBook
	Author string
	Title  string
}

func WrittenBookFromNbt(node *nbt.TagNodeCompound) *WrittenBook {
	type txt struct {
		Text string `json:"text"`
	}
	b := new(WrittenBook)
	b.Author = node.Entries["author"].(*nbt.TagNodeString).String()
	b.Title = node.Entries["title"].(*nbt.TagNodeString).String()
	b.pages = make([]string, 0)
	pages := node.Entries["pages"].(*nbt.TagNodeList)
	pages.Each(func(page nbt.ITagNode) {
		var t txt
		_ = json.Unmarshal([]byte(page.(*nbt.TagNodeString).String()), &t)
		b.pages = append(b.pages, t.Text)
	})
	return b
}

type IBook interface {
	Pages() []string
}

type BookItem struct {
	BaseItem
	Book IBook
}

func WritableBookItemFromNbt(node *nbt.TagNodeCompound) *BookItem {
	i := new(BookItem)
	i.BaseItem = *baseItemFromNbt(node)
	if tag, ok := node.Entries["tag"].(*nbt.TagNodeCompound); ok {
		i.Book = WritableBookFromNbt(tag)
	} else {
		i.Book = NewWritableBook()
	}
	return i
}

func WrittenBookItemFromNbt(node *nbt.TagNodeCompound) *BookItem {
	i := new(BookItem)
	i.BaseItem = *baseItemFromNbt(node)
	if tag, ok := node.Entries["tag"].(*nbt.TagNodeCompound); ok {
		i.Book = WrittenBookFromNbt(tag)
	}
	return i
}
