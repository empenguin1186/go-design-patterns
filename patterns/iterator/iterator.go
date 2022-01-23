package iterator

import "fmt"

type Aggregate interface {
	Iterator() Iterator
}

type Iterator interface {
	HasNext() bool
	Next() (*Item, error)
}

type Item struct {
	name string
}

func NewItem(name string) *Item {
	return &Item{name}
}

func (i *Item) GetName() string {
	return i.name
}

type ShoppingCart struct {
	items []*Item
}

func NewShoppingCart() *ShoppingCart {
	return &ShoppingCart{
		items: []*Item{},
	}
}

func (s *ShoppingCart) GetItemAt(index int) (*Item, error) {
	if len(s.items)-1 < index {
		return nil, fmt.Errorf("index out of bounds. index=%d", index)
	}
	return s.items[index], nil
}

func (s *ShoppingCart) AppendItem(item *Item) {
	s.items = append(s.items, item)
}

func (s *ShoppingCart) GetLength() int {
	return len(s.items)
}

func (s *ShoppingCart) Iterator() Iterator {
	return NewShoppingCartIterator(s)
}

type ShoppingCartIterator struct {
	shoppingCart *ShoppingCart
	index        int
}

func NewShoppingCartIterator(shoppingCart *ShoppingCart) *ShoppingCartIterator {
	return &ShoppingCartIterator{shoppingCart: shoppingCart, index: 0}
}

func (s *ShoppingCartIterator) HasNext() bool {
	if s.index < len(s.shoppingCart.items) {
		return true
	}
	return false
}

func (s *ShoppingCartIterator) Next() (*Item, error) {
	if s.index < len(s.shoppingCart.items) {
		item := s.shoppingCart.items[s.index]
		s.index += 1
		return item, nil
	}
	return nil, fmt.Errorf("index out of bounds. index=%d", s.index)
}
