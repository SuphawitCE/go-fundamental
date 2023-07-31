package input

import (
	"cart/spec"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Fetch products from the API
func FetchProducts() ([]spec.Product, error) {
	url := "https://fakestoreapi.com/products"

	// Send GET request
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// Parse JSON response
	var products []spec.Product
	err = json.Unmarshal(body, &products)
	if err != nil {
		return nil, err
	}

	return products, nil
}
