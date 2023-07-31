# Cart

## [Draft] Problem Statement

Implement E-Commerce cart

Using mock product data from https://fakestoreapi.com/products assume it is internal Database.
The cart should be able to perform basic functionality to collect products.

## Implementation Details

1. Define a struct to represent product data structure.
2. Fetch `GET` product from `https://fakestoreapi.com/products`, then assign to product data structure.
3. Define a struct to represent cart data structure.
4. Get product id input from keyboard.
5. Insert matching product id to cart.
6. Calculate price and count the total of price.