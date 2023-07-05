//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:

package main

import "fmt"

func greetPerson(name string) {
	fmt.Println("Hello", name)
}

func hiThere() string {
	return "Hi There"
}

func twoTwos() (int, int) {
	return 2, 4
}

func five() int {
	return 5
}

func addThree(a, b, c int) int {
	return a + b + c
}

func main() {
	//* Write a function that accepts a person's name as a function
	//  parameter and displays a greeting to that person.
	greetPerson("Bank")

	//* Write a function that returns any message, and call it from within
	//  fmt.Println()
	fmt.Println(hiThere())

	//* Write a function to add 3 numbers together, supplied as arguments, and
	//  return the answer
	//* Write a function that returns any number
	//* Write a function that returns any two numbers
	a, b := twoTwos()

	//* Add three numbers together using any combination of the existing functions.
	answer := addThree(five(), a, b)

	//  * Print the result
	//* Call every function at least once

	fmt.Println("Cal=", answer)
}
