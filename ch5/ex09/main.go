package main

import (
	"fmt"
	"regexp"
)

func echo(s string) string {
	return s
}

var re = regexp.MustCompile(`\$[^\s]+`)

func expand(s string, f func(string) string) string {
	return re.ReplaceAllStringFunc(s, func(s string) string {
		return f(s[1:])
	})
}

func main() {
	s := "ほげほげお $foo ばー"
	fmt.Println(expand(s, echo))
}
