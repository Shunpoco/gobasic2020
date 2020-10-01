package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var s = flag.String("crypto", "sha256", "set cryptographic hash functions")

func main() {
	flag.Parse()
	fmt.Println(*s)
	if len(os.Args) < 2 {
		fmt.Println("文字列を渡してください")
		os.Exit(1)
	}
	str := os.Args[1]

	switch *s {
	case "sha256":
		fmt.Printf("%x\n", sha256.Sum256([]byte(str)))
	case "sha384":
		fmt.Printf("%x\n", sha512.Sum384([]byte(str)))
	case "sha512":
		fmt.Printf("%x\n", sha512.Sum512([]byte(str)))
	default:
		fmt.Println("invalid option, crypto by sha256...")
		fmt.Printf("%x\n", sha256.Sum256([]byte(str)))
	}
}
