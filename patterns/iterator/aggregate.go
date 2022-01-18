package iterator

import "fmt"

type Iterator interface {
	HasNext() bool
	Next() Book
}

type Book struct {
	name string
}

func NewBook(name string) *Book {
	return &Book{name}
}

func (b Book) GetName() string {
	return b.name
}

type BookShelf struct {
	books []*Book
}

func NewBookShelf() *BookShelf {
	return &BookShelf{
		books: []*Book{},
	}
}

func (b *BookShelf) GetBookAt(index int) (*Book, error) {
	if len(b.books)-1 < index {
		return nil, fmt.Errorf("index out of bounds. index=%d", index)
	}
	return b.books[index], nil
}

func (b *BookShelf) AppendBook(book *Book) {
	b.books = append(b.books, book)
}

func (b *BookShelf) GetLength() int {
	return len(b.books)
}

func (b *BookShelf) Iterator() Iterator {
	return NewBookShelfIterator(b)
}
