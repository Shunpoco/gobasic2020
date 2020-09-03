// fetches URLs in parallel and reports their time
package main

import (
	"io"
	"net/http"
	"os"
	"time"

	"golang.org/x/exp/errors/fmt"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	file, err := os.Create("./result.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
	// 2 times
	for i := 0; i < 2; i++ {
		for _, url := range os.Args[1:] {
			go fetch(url, ch, file)
		}
		for range os.Args[1:] {
			fmt.Println(<-ch)
		}
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string, w io.Writer) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(w, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("While reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
