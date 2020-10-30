package count

import (
	"bufio"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.TextNode {
		text := n.Data
		scanner := bufio.NewScanner(strings.NewReader(text))
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			words++
		}
	}
	if n.Type == html.ElementNode && n.Data == "img" {
		for _, a := range n.Attr {
			if a.Key == "src" && a.Val != "" {
				images++
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		tWords, tImages := countWordsAndImages(c)
		words = words + tWords
		images = images + tImages
	}
	return
}
