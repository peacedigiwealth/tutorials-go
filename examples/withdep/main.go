package main

import (
	"github.com/peace/color"
)

func main() {
	// print a colored message using a tiny third-party package
	color.New(color.FgGreen).Println("Hello with color from peace/color")
}
