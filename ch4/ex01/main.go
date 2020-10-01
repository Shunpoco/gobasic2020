package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		log.Print("2 strings required...")
		os.Exit(1)
	}

	str1 := os.Args[1]
	str2 := os.Args[2]
	fmt.Println(str1)
	fmt.Println(str2)

	c1 := sha256.Sum256([]byte(str1))
	c2 := sha256.Sum256([]byte(str2))

	fmt.Printf("%x\n", c1)
	fmt.Printf("%x\n", c2)
	fmt.Println(Counter(c1, c2))
}

func Counter(c1, c2 [sha256.Size]byte) int {
	result := 0
	for i := 0; i < sha256.Size; i++ {
		result += int(c1[i] ^ c2[i])
	}
	return result
}
