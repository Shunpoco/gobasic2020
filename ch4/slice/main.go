package main

import "fmt"

func main() {
	var a [3]int
	fmt.Printf("%T\n", a)
	a[0] = 1
	fmt.Println(a)
}
