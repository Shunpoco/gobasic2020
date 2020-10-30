package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var elements = map[string]bool{"a": true, "img": true, "style": true, "script": true}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && elements[n.Data] {
		for _, a := range n.Attr {
			if a.Key == "href" || a.Key == "src" {
				links = append(links, n.Data+":"+a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}
