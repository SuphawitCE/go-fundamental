package input

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Image       string  `json:"image"`
}

func GetCart() []Product {
	fmt.Println("GetCart")
	url := "https://fakestoreapi.com/products"

	// Send GET request
	response, err := http.Get(url)

	if err != nil {
		fmt.Println("Get product error", err)
		// return
	}

	defer response.Body.Close()

	// Parse JSON response
	var products []Product
	err = json.NewDecoder(response.Body).Decode(&products)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		// return
	}
	return products
}
