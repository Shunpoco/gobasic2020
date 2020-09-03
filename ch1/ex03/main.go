package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	inputs := os.Args[1:]
	fmt.Println(SlowEcho(inputs))
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

	start = time.Now()
	fmt.Println(FastEcho(inputs))
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func SlowEcho(inputs []string) (s string) {
	var sep string
	for _, val := range inputs {
		s += sep + val
		sep = " "
	}
	return
}

func FastEcho(inputs []string) string {
	return strings.Join(inputs, " ")
}
