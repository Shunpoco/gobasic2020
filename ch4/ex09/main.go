package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		if err := input.Err(); err == io.EOF {
			break
		}
		counts[input.Text()]++
	}
	printCount(counts)
}

func printCount(count map[string]int) {
	fmt.Printf("word\tcount\n")
	for s, n := range count {
		fmt.Printf("%s\t%d\n", s, n)
	}
	fmt.Print("\nlen\tcount\n")
}
