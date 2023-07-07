package main

import (
	"go-fundamental/src/demo/pkg/display"
	"go-fundamental/src/demo/pkg/msg"
)

func main() {
	msg.Hi()
	display.Display("Hello from display main.go")
	msg.Exciting("An exciting message")
}
