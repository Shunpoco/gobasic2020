package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var wc WordCounter
	fmt.Fprintf(&wc, "Hello, world!\nHogehoge fugafuga hoge")
	fmt.Println(wc)

	var lc LineCounter
	fmt.Fprintf(&lc, "Hello, world!\nHogehoge fugafuga hoge")
	fmt.Println(lc)
}

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	count := 0

	for scanner.Scan() {
		count++
	}
	*c += WordCounter(count)
	return len(p), nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	count := 0
	for scanner.Scan() {
		count++
	}
	*c += LineCounter(count)
	return len(p), nil
}
