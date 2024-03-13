package main

import "fmt"

type Product struct {
	price int
	name  string
}

func printStats(products [4]Product) {
	var cost, totalItem int

	for i := 0; i < len(products); i++ {
		product := products[i]
		cost += product.price

		if product.name != "" {
			totalItem += 1
		}
	}

	fmt.Println("cost: ", cost)
	fmt.Println("number of items: ", totalItem)
	fmt.Println("last  of item: ", products[totalItem-1])

}

func main() {

	shoppingList := [4]Product{
		{1, "Banana"},
		{4, "Apple"},
		//{6, "Orange"},
		{18, "Lemon"},
	}

	printStats(shoppingList)

	shoppingList[3] = Product{
		price: 50,
		name:  "Notebook",
	}
	fmt.Println("----------------")

	printStats(shoppingList)

}
