package main

import (
	"cart/input"
	"fmt"
)

/*
Implement cart functionality
Get products from https://fakestoreapi.com/products
User could purchase any product based on product id
The cart could be able to handle CRUD (CREATE, READ, UPDATE, UPDATE) product
The cart could be able to display both specific and all product and cost calculation with US concurrency
The cart should not display a product if an attribute is not define
*/

func main() {
	fmt.Println("Cart")

	// Receive input from user

	ids := input.ReadInputFromKeyBoard()

	fmt.Println("ProductIds:", ids)

	// fmt.Println("input:", input.ReadInputFromKeyBoard())
	// fmt.Println("gg", testpkg.greet())

	// var name string

	// fmt.Print("Enter your name: ")

	// _, err := fmt.Scan(&name)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("Hello %s!\n", name)
}
