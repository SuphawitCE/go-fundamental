package main

import "fmt"

func double(num int) int {
	return num + num
}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func greet() {
	fmt.Println("Hello from greet function")
}

func main() {
	greet()

	dozen := double(5)
	fmt.Println("double=", dozen)

	bakersDozen := add(dozen, 2)
	fmt.Println("A baker's dozen is", bakersDozen)

	anotherBakersDozen := add(double(5), 1)
	fmt.Println("another baker's dozen is=", anotherBakersDozen)
}
