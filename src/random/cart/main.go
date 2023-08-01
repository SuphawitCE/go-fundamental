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
	productIds, productQuantity := input.ReadInputFromKeyBoard()
	if err != nil {
		fmt.Println("Error fetching product data:", err)
		return
	}

	// Create a cart
	cart := spec.Cart{}
	fmt.Println("pp", products[0])

	// Add some products to the cart
	for i := 0; i < len(productIds); i++ {
		utils.AddToCart(&cart, products[productIds[i]], productQuantity[i], productIds[i])
	}

	// Print the cart
	printCart(&cart)

	// Remove a product from the cart
	utils.RemoveFromCart(&cart, products[0].ID)
	printCart(&cart)
}
