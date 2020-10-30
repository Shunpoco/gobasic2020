package main

import (
	"ch5/ex13/links"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	netURL "net/url"
	"os"
	"strings"
)

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		// fmt.Println(worklist)
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	rootURL, err := netURL.Parse(url)
	hostName := rootURL.Hostname()
	if err != nil {
		log.Print(err)
	}
	if _, err := os.Stat(hostName); os.IsNotExist(err) {
		os.Mkdir(rootURL.Hostname(), os.ModePerm)
	}
	save(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	breadthFirst(crawl, os.Args[1:])
}

func save(url string) {
	rootURL, err := netURL.Parse(url)
	path := rootURL.Hostname() + rootURL.Path
	if !strings.HasSuffix(path, "/") {
		path = path + "/"
	}
	fmt.Println(path)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}
	resp, err := http.Get(url)
	if err != nil {
		log.Print(err)
	}
	dump, _ := httputil.DumpResponse(resp, true)

	err = ioutil.WriteFile(path+"index.html", dump, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
