package main

import (
	"errors"
	"fmt"
)

// Product represents the product entity.
type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
}

// ProductRepository defines the interface for product data access.
type ProductRepository interface {
	GetAll() ([]Product, error)
	GetByID(id int) (*Product, error)
	Create(product *Product) error
	Update(product *Product) error
	Delete(id int) error
}

// inMemoryProductRepository is an in-memory implementation of ProductRepository.
type inMemoryProductRepository struct {
	products map[int]*Product
}

// NewInMemoryProductRepository creates a new instance of inMemoryProductRepository.
func NewInMemoryProductRepository() ProductRepository {
	return &inMemoryProductRepository{
		products: make(map[int]*Product),
	}
}

func (r *inMemoryProductRepository) GetAll() ([]Product, error) {
	products := make([]Product, 0, len(r.products))
	for _, p := range r.products {
		products = append(products, *p)
	}
	return products, nil
}

func (r *inMemoryProductRepository) GetByID(id int) (*Product, error) {
	product, found := r.products[id]
	if !found {
		return nil, errors.New("product not found")
	}
	return product, nil
}

func (r *inMemoryProductRepository) Create(product *Product) error {
	if _, found := r.products[product.ID]; found {
		return errors.New("product with the same ID already exists")
	}
	r.products[product.ID] = product
	return nil
}

func (r *inMemoryProductRepository) Update(product *Product) error {
	if _, found := r.products[product.ID]; !found {
		return errors.New("product not found")
	}
	r.products[product.ID] = product
	return nil
}

func (r *inMemoryProductRepository) Delete(id int) error {
	if _, found := r.products[id]; !found {
		return errors.New("product not found")
	}
	delete(r.products, id)
	return nil
}

func main() {
	repo := NewInMemoryProductRepository()

	// Create a new product
	newProduct := &Product{
		ID:          1,
		Name:        "Sample Product",
		Price:       19.99,
		Description: "This is a sample product.",
	}
	err := repo.Create(newProduct)
	if err != nil {
		fmt.Println("Error creating product:", err)
		return
	}

	// Get all products
	products, err := repo.GetAll()
	if err != nil {
		fmt.Println("Error getting products:", err)
		return
	}
	fmt.Println("All Products:")
	fmt.Println(products)

	// Get product by ID
	productByID, err := repo.GetByID(1)
	if err != nil {
		fmt.Println("Error getting product by ID:", err)
		return
	}
	fmt.Println("Product with ID 1:")
	fmt.Println(productByID)

	// Update product
	productToUpdate := &Product{
		ID:          1,
		Name:        "Updated Product",
		Price:       24.99,
		Description: "This product has been updated.",
	}
	err = repo.Update(productToUpdate)
	if err != nil {
		fmt.Println("Error updating product:", err)
		return
	}

	// Delete product
	err = repo.Delete(1)
	if err != nil {
		fmt.Println("Error deleting product:", err)
		return
	}

	fmt.Println("Product with ID 1 has been deleted.")
}
