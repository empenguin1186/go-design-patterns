package iterator

import "fmt"

type Aggregate interface {
	Iterator() Iterator
}

type Iterator interface {
	HasNext() bool
	Next() (*Book, error)
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

type BookShelfIterator struct {
	bookShelf *BookShelf
	index     int
}

func NewBookShelfIterator(bookShelf *BookShelf) *BookShelfIterator {
	return &BookShelfIterator{bookShelf: bookShelf, index: 0}
}

func (b *BookShelfIterator) HasNext() bool {
	if b.index < len(b.bookShelf.books) {
		return true
	} else {
		return false
	}
}

func (b *BookShelfIterator) Next() (*Book, error) {
	if b.index < len(b.bookShelf.books) {
		book := b.bookShelf.books[b.index]
		b.index += 1
		return book, nil
	} else {
		return nil, fmt.Errorf("index out of bounds. index=%d", b.index)
	}
}
