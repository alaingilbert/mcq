package mc

import (
	"github.com/alaingilbert/mcq/nbt"
)

type Lectern struct {
	blockEntity
	bookItem *BookItem
}

func LecternFromNbt(node *nbt.TagNodeCompound) *Lectern {
	l := new(Lectern)
	l.blockEntity = *BlockEntityFromNbt(node)

	if book, ok := node.Entries["Book"].(*nbt.TagNodeCompound); ok {
		bookID := ID(book.Entries["id"].(*nbt.TagNodeString).String())
		switch bookID {
		case WrittenBookID:
			l.bookItem = WrittenBookItemFromNbt(book)
		case WritableBookID:
			l.bookItem = WritableBookItemFromNbt(book)
		}
	}
	return l
}

func (l *Lectern) HasBook() bool {
	return l.bookItem != nil
}

func (l *Lectern) BookItem() *BookItem {
	return l.bookItem
}

func (l *Lectern) Book() IBook {
	if l.HasBook() {
		return l.bookItem.Book
	}
	return nil
}
