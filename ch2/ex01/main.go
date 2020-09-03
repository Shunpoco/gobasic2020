package main

import (
	"ch2/ex01/tempconv"
	"fmt"
)

func main() {
	c := tempconv.Celsius(0)
	fmt.Printf("%s\n", c.String())
	fmt.Printf("Brrrr %v\n", tempconv.AbsoluteZeroC)
	fmt.Println(tempconv.CToF(c))
	k := tempconv.AbsoluteZeroK
	fmt.Println(k.String())
	k2 := tempconv.Kelvin(100)
	fmt.Printf("%s is %s", k2.String(), tempconv.KToC(k2))
}
