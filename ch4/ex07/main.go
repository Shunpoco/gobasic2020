package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	b := []byte("hogehoge fuga 世界")
	fmt.Println(string(b))
	fmt.Println(string(reverse(b)))
}

func reverse(b []byte) []byte {
	// 各rune単位でreverseする。
	// size = 1のruneは何も影響はない
	for i := 0; i < len(b); {
		_, size := utf8.DecodeRune(b[i:])
		reverseByte(b[i : i+size])
		i += size
	}
	// 全体をreverse
	reverseByte(b)
	return b
}

func reverseByte(b []byte) []byte {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return b
}
