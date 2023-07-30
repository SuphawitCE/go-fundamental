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

/*
Output:
title, price, totalPrice
*/

func main() {
	fmt.Println("Cart")
	// Receive input from user
	ids := input.ReadInputFromKeyBoard()
	fmt.Println("ProductIds:", ids)

	product := input.GetCart()
	fmt.Println("Product:", product[0])
	var temp []input.Product
	for _, item := range ids {
		fmt.Println("item", item)
		product[item-1].Description = ""
		temp = append(temp, product[item-1])
	}
	fmt.Println("temp:", temp)
}
