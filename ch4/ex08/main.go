// charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	countsL := make(map[rune]int)
	countsC := make(map[rune]int)
	countsM := make(map[rune]int)
	countsN := make(map[rune]int)
	countsP := make(map[rune]int)
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		// EOFを与えるには、Ctrl+dで行う
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsLetter(r) {
			countsL[r]++
		}
		if unicode.IsControl(r) {
			countsC[r]++
		}
		if unicode.IsMark(r) {
			countsM[r]++
		}
		if unicode.IsNumber(r) {
			countsN[r]++
		}
		if unicode.IsPunct(r) {
			countsP[r]++
		}
	}
	printCount(countsL, "letters")
	printCount(countsC, "controls")
	printCount(countsM, "marks")
	printCount(countsN, "numbers")
	printCount(countsP, "punctuations")

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

func printCount(count map[rune]int, category string) {
	fmt.Printf("rune\tcount of %s\n", category)
	for c, n := range count {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
}
