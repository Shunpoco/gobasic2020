package main

import (
	"fmt"
	"pop/popcount"
)

func main() {
	fmt.Println(popcount.PopCountBitShift(255))
	fmt.Println(popcount.PopCountBitShift(256))
}
