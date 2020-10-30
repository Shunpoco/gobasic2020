package main

import "fmt"

func Join(elems ...string) (joined string) {
	if len(elems) < 2 {
		return
	}
	sep := elems[len(elems)-1]
	elems = elems[:len(elems)-1]
	for idx, val := range elems {
		joined += val
		if idx < len(elems)-1 {
			joined += sep
		}
	}
	return
}

func main() {
	fmt.Println(Join("hoge", "fuga", "fizz", "buzz", " "))
}
