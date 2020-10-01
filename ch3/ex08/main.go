package main

import (
	_ "ch3/ex08/mondelbrot128"
	"ch3/ex08/mondelbrot64"
)

func main() {
	mondelbrot64.Mondelbrot()
	// mondelbrot128.Mondelbrot()
}
