//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	name  string
	price int
}

func printStats(products [4]Product) {
	var cost, totalItem int

	for i := 0; i < len(products); i++ {
		item := products[i]
		cost += item.price

		if item.name != "" {
			totalItem++
		}
	}

	fmt.Println("Last item:", products[totalItem-1])
	fmt.Println("Total items:", totalItem)
	fmt.Println("Total cost:", cost)
}

func main() {
	shoppingList := [4]Product{
		{name: "Milk", price: 2},
		{name: "Bread", price: 3},
		{name: "Eggs", price: 4},
	}

	printStats(shoppingList)

	shoppingList[3].name = "Butter"
	shoppingList[3].price = 5

	printStats(shoppingList)
}
