package main

import (
	"fmt"
	"pop/popcount"
)

func main() {
	fmt.Println(popcount.PopCountExclusive(10))
	fmt.Println(popcount.PopCountExclusive(255))
}
