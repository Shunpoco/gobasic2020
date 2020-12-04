package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func main() {
	args := os.Args[1:]
	for _, val := range args {
		loc := strings.Split(val, "=")
		if len(loc) == 2 {
			go dial(loc[1], loc[0])
		}
	}
	for {
	}
}

func dial(url string, name string) {
	conn, err := net.Dial("tcp", url)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	fmt.Fprintf(os.Stdout, "%s\n", name)
	mustCopy(os.Stdout, conn)
}
