package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	rotate(s, 2)
	fmt.Println(s)
	// nを大きくすると何もしない
	rotate(s, 9)
	fmt.Println(s)
}

func rotate(s []int, n int) {
	if n > len(s) {
		return
	}
	temp := append(s, s[:n]...)
	copy(s, temp[n:])
}
