package main

import (
	"log"
	"net/http"
	"net/url"
	"strconv"

	"ch3/ex09/mondelbrot"
)

func main() {
	http.HandleFunc("/", handler)

	if err := http.ListenAndServe("localhost:8888", nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	w.Header().Set("Content-Type", "image/png")
	mo := mondelbrot.NewMondelbrotOptions(
		fetchProperIntValue(r.Form, "width"),
		fetchProperIntValue(r.Form, "height"),
		fetchProperIntValue(r.Form, "scale"),
	)
	mondelbrot.Mondelbrot(w, mo)
}

func fetchProperIntValue(values url.Values, key string) (result int) {
	if val := values.Get(key); val != "" {
		value, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return
		}
		result = int(value)
	}
	return
}
