package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	fileMemory := make(map[string][]string)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, fileMemory, "")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, fileMemory, arg)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t %s\t %v\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int, memory map[string][]string, name string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		memory[input.Text()] = append(memory[input.Text()], name)
	}
}
