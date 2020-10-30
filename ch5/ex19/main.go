package main

import "fmt"

func panicRecover(p string) (s string) {
	defer func() {
		p := recover()
		s = fmt.Sprintf("%s", p)
	}()

	panic(p)
}

func main() {
	s := panicRecover("panic")
	fmt.Println(s)
}
