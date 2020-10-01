package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("２つの文字列を渡してください")
		os.Exit(1)
	}
	str1 := os.Args[1]
	str2 := os.Args[2]
	fmt.Println(str1)
	fmt.Println(str2)

	fmt.Println(isAnagram(str1, str2))
}

func isAnagram(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	mapStr1 := make(map[rune]int)
	mapstr2 := make(map[rune]int)

	for _, r := range str1 {
		mapStr1[r]++
	}
	for _, r := range str2 {
		mapstr2[r]++
	}

	for r, val := range mapStr1 {
		if mapstr2[r] != val {
			return false
		}
	}

	return true
}
