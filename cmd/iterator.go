package main

import (
	"fmt"
	"github.com/empenguin1186/go-design-patterns/patterns/iterator"
	"os"
)

func main() {
	bookShelf := iterator.NewBookShelf()
	bookShelf.AppendBook(iterator.NewBook("HOGE"))
	bookShelf.AppendBook(iterator.NewBook("FUGA"))
	bookShelf.AppendBook(iterator.NewBook("PIYO"))

	i := bookShelf.Iterator()
	for i.HasNext() {
		book, err := i.Next()
		if err != nil {
			fmt.Println("error occurred.")
			os.Exit(1)
		}
		fmt.Println(book.GetName())
	}
}
