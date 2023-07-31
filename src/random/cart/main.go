package main

import (
	"cart/input"
	"cart/spec"
	"cart/utils"
	"fmt"
)

// Print the cart items
func printCart(cart *spec.Cart) {
	fmt.Println("Cart:")
	for _, item := range cart.Items {
		fmt.Printf("Product ID: %d, Title: %s, Quantity: %d\n", item.Product.ID, item.Product.Title, item.Quantity)
	}
	fmt.Println("-----------------------")
}

func main() {
	// Fetch mock product data from the API
	products, err := input.FetchProducts()
	if err != nil {
		fmt.Println("Error fetching product data:", err)
		return
	}

	// Create a cart
	cart := spec.Cart{}

	// Add some products to the cart
	utils.AddToCart(&cart, products[0], 2)
	utils.AddToCart(&cart, products[1], 1)

	// Print the cart
	printCart(&cart)

	// Remove a product from the cart
	utils.RemoveFromCart(&cart, products[0].ID)
	printCart(&cart)
}
