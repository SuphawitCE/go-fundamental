package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
Terminate go terminal
Windows: Control+C
Mac and Linux: Control+D
*/

func main() {
	read := bufio.NewReader(os.Stdin)
	sum := 0
	fmt.Println("Insert your input:")
	for {
		input, inputError := read.ReadString(' ')
		space := strings.TrimSpace(input)
		if space == "" {
			continue
		}
		num, convertError := strconv.Atoi(space)
		if convertError != nil {
			fmt.Println(convertError)
		} else {
			sum += num
		}

		if inputError == io.EOF {
			fmt.Println("End Of File")
			break
		}

		if inputError != nil {
			fmt.Println("Error reading Stdio:", inputError)
		}
	}
	fmt.Printf("sum: %v\n", sum)
}
