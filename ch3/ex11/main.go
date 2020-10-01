package main

import (
	"fmt"
)

func main() {
	f := 123456789.0123456
	fmt.Println(f)
	s := fmt.Sprintf("%f", f)
	fmt.Println(comma(s))
	fmt.Println(commaWithDot(s))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func commaWithDot(s string) string {
	var dotIdx int
	n := len(s)
	for i := 0; i < n; i++ {
		if s[i] == '.' {
			dotIdx = i
		}
	}
	return comma(s[:dotIdx]) + s[dotIdx:]
}
