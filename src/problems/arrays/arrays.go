package main

import "fmt"

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

type Product struct {
	price int
	name  string
}

func printStats(list [4]Product) {

	var cost, totalItems int
	for i := 0; i < len(list); i++ {
		item := list[i]
		cost += item.price

		if item.name != "" {
			totalItems += 1
		}
	}

	//* Print to the terminal:
	//  - The last item on the list
	fmt.Println("Last item on the list:", list[totalItems-1])
	//  - The total number of items
	fmt.Println("Total items", totalItems)
	//  - The total cost of the items
	fmt.Println("Total cost", cost)

}

func main() {

	//* Using an array, create a shopping list with enough room
	//  for 4 products
	//  - Products must include the price and the name
	shoppingList := [4]Product{
		{1, "Banana"},
		{6, "Meat"},
		{3, "Salad"},
	}

	printStats(shoppingList)
	shoppingList[3] = Product{4, "Bread"}
	printStats(shoppingList)

}
