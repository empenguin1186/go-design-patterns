package main

import (
	"fmt"
	"github.com/empenguin1186/go-design-patterns/patterns/iterator"
	"os"
)

func main() {
	shoppingCart := iterator.NewShoppingCart()
	shoppingCart.AppendItem(iterator.NewItem("HOGE"))
	shoppingCart.AppendItem(iterator.NewItem("FUGA"))
	shoppingCart.AppendItem(iterator.NewItem("PIYO"))

	i := shoppingCart.Iterator()
	for i.HasNext() {
		book, err := i.Next()
		if err != nil {
			fmt.Println("error occurred.")
			os.Exit(1)
		}
		fmt.Println(book.GetName())
	}
}
