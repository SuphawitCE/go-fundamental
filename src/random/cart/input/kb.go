package input

import (
	"fmt"
	"log"
)

func ReadInputFromKeyBoard() []int {
	fmt.Print("Number of product in cart: ")
	var product int
	_, err := fmt.Scan(&product)

	if err != nil {
		log.Fatal(err)
	}

	var productIds = make([]int, product)

	for i := 0; i < len(productIds); i++ {
		_, err := fmt.Scan(&productIds[i])

		if err != nil {
			log.Fatal(err)
		}
	}

	return productIds
}
