package main

import "fmt"

//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

type Coordinate struct {
	x, y int
}

type Rectangle struct {
	left  Coordinate
	right Coordinate
}

func width(rect Rectangle) int {
	return rect.left.x - rect.right.x
}

func length(rect Rectangle) int {
	return rect.right.y - rect.right.x
}

func area(rect Rectangle) int {
	return length(rect) * width(rect)
}

func perimeter(rect Rectangle) int {
	return (width(rect) * 2) + (length(rect) * 2)
}

func printInfo(rect Rectangle) {
	fmt.Println("Area is", area(rect))
	fmt.Println("Perimeter is", perimeter(rect))
}

func main() {
	rect := Rectangle{
		right: Coordinate{0, 7},
		left:  Coordinate{10, 0},
	}

	printInfo(rect)
	rect.right.y *= 2
	rect.left.x *= 2
	printInfo(rect)
}
