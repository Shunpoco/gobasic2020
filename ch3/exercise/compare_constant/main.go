package main

import "fmt"

const (
	a = 1000000
	b = 100
	c = iota // 2になる
)

const (
	d = iota // 0になる
)

func main() {
	fmt.Println(a/10000 == b)
	fmt.Println(c)
	fmt.Println(d)

	var e = [3]int{}

	const f = len(e) // ok
	const g = cap(e) // ok
}
