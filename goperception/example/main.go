package main

import (
	"fmt"

	"github.com/moritztng/perceptionOS/goperception"
)

func main() {
	capture := goperception.Camera("asdf")
	capture.SaveImage("asdf")
	fmt.Println("Hello, World!")
}
