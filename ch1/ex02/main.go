package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for idx, val := range os.Args {
		fmt.Println(strconv.Itoa(idx) + " " + val)
	}
}
