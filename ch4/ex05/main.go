package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Print("You must pass more than one of words.")
	}
	s := []string{}
	s = append(s, os.Args[1:]...)
	fmt.Println(s)
	s = removeDup(s)
	fmt.Println(s)
}

func removeDup(s []string) []string {
	var left string
	result := []int{}
	for i, val := range s {
		if val == left {
			result = append(result, i)
			continue
		}
		left = val
	}
	for i, v := range result {
		s = remove(s, v-i)
	}
	return s
}

func remove(s []string, i int) []string {
	copy(s[i:], s[i+1:])
	fmt.Println(s)
	return s[:len(s)-1]
}
