package utils

import (
	"cart/spec"
)

// Add a product to the cart
func AddToCart(cart *spec.Cart, product spec.Product, quantity int) {
	cartItem := spec.CartItem{
		Product:  product,
		Quantity: quantity,
	}

	cart.Items = append(cart.Items, cartItem)
}

// Remove a product from the cart
func RemoveFromCart(cart *spec.Cart, productID int) {
	for i, item := range cart.Items {
		if item.Product.ID == productID {
			cart.Items = append(cart.Items[:i], cart.Items[i+1:]...)
			break
		}
	}
}
