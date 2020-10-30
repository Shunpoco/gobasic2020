package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}

	counter := make(map[string]int)
	counter = elementCount(counter, doc)

	for k, v := range counter {
		fmt.Printf("%s: %d\n", k, v)
	}
}

func elementCount(counter map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		counter[n.Data]++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		counter = elementCount(counter, c)
	}
	return counter
}
