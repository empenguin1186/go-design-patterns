package main

import (
	"fmt"
	"github.com/empenguin1186/go-design-patterns/patterns/iterator"
)

func main() {
	shoppingCart := iterator.NewShoppingCart()
	shoppingCart.AppendItem(iterator.NewItem("HOGE"))
	shoppingCart.AppendItem(iterator.NewItem("FUGA"))
	shoppingCart.AppendItem(iterator.NewItem("PIYO"))

	i := shoppingCart.Iterator()
	for i.HasNext() {
		item, _ := i.Next()
		fmt.Println(item.GetName())
	}
}
