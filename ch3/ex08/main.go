package main

import (
	_ "ch3/ex08/mondelbrot128"
	_ "ch3/ex08/mondelbrot64"
	_ "ch3/ex08/mondelbrotBigFloat"
	"ch3/ex08/mondelbrotBigRat"
)

func main() {
	// mondelbrot64.Mondelbrot()
	// mondelbrot128.Mondelbrot()
	// mondelbrotBigFloat.Mondelbrot()
	mondelbrotBigRat.Mondelbrot()
}
