package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	var nodes []*html.Node
	nodes = find(nodes, doc, name)
	return nodes
}

func find(nodes []*html.Node, doc *html.Node, name []string) []*html.Node {
	if doc.Type == html.ElementNode {
		if contains(doc.Data, name) {
			nodes = append(nodes, doc)
		}
	}
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		nodes = find(nodes, c, name)
	}
	return nodes
}

func contains(data string, tags []string) bool {
	for _, tag := range tags {
		if data == tag {
			return true
		}
	}
	return false
}

func main() {
	url := os.Args[1]
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)

	nodes := ElementsByTagName(doc, "a", "img")

	for _, node := range nodes {
		fmt.Printf("<%s>\n", node.Data)
	}
}
