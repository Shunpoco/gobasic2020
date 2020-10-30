package main

import (
	"ch5/ex05/count"
	"fmt"
	"os"
)

func main() {
	url := os.Args[1]
	words, images, err := count.CountWordsAndImages(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("words: %d and images: %d", words, images)
}
