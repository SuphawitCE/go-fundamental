package spec

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

type CartItem struct {
	Product
	Quantity int
}

type Cart struct {
	Items []CartItem
}
