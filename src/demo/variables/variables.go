package main

import "fmt"

func main() {
	var myName = "Bank"
	fmt.Println("My name is", myName, myName);

	var name string = "cat"
	fmt.Println("name = ", name)

	userName := "admin"
	fmt.Println("userName: ", userName)

	var sum int
	fmt.Println("The sum is", sum)

	part1, other := 1, 5
	fmt.Println("part 1 is", part1, "other is", other)

	part2, other := 2, 0
	fmt.Println("part 2 is", part2, "other is", other)

	sum = part1 + part2
	fmt.Println("sum =", sum)

	var(
		lessonName = "Variables"
		lessonType =  "Demo"
	)
	fmt.Println("lessonName=", lessonName)
	fmt.Println("lessonType=", lessonType)

	word1, word2, _ := "Hello", "World", "!"
	fmt.Println(word1, word2)
}
