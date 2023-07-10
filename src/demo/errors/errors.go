package main

import (
	"fmt"
)

type Stuff struct {
	values []int
}

func (s *Stuff) Get(index int) (int, error) {
	if index > len(s.values) {
		return 0, fmt.Errorf("no element at index %v", index)
		// errors.New(fmt.Sprintf("No element at index %v", index))
	} else {
		return s.values[index], nil
	}
}

func main() {
	stuff := Stuff{}
	value, err := stuff.Get(1)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Value is", value)
	}
}
