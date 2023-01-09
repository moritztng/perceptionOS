package main

import (
    "github.com/moritztng/perceptionOS/goperception"
	"fmt"
)

func main() {
	capture := goperception.camera("asdf")
    capture.SaveImage("asdf")
    fmt.Println("Hello, World!")
}
