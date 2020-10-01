package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println(comma("hogehogehogehoge"))
		return
	}
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(comma(os.Args[i]))
	}
}

func comma(s string) string {
	var buf bytes.Buffer
	n := len(s)
	// commaを打つ点を初期化する
	i := (3 - n%3) % 3
	for _, r := range s {
		if i == 3 {
			buf.WriteByte(',')
			i = 0
		}
		buf.WriteRune(r)
		i++
	}
	return buf.String()
}
