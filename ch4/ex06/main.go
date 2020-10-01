package main

import (
	"bytes"
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	b := []byte("hoge  　ho  g e")
	fmt.Println(string(removeDuplicatedSpace(b)))
}

func removeDuplicatedSpace(b []byte) []byte {
	var buf bytes.Buffer
	for len(b) > 0 {
		r, size := utf8.DecodeRune(b)
		if unicode.IsSpace(r) {
			nextR, _ := utf8.DecodeRune(b[size:])
			if !unicode.IsSpace(nextR) {
				buf.WriteRune(' ')
			}
		} else {
			buf.WriteRune(r)
		}
		b = b[size:]
	}
	return buf.Bytes()
}
