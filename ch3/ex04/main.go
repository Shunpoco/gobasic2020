package main

import (
	"ch3/ex04/surface"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(fmt.Sprintf("localhost:%d", 8888), nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	w.Header().Set("Content-Type", "image/svg+xml")

	sgo := surface.NewSurfaceGraphOptions(
		fetchProperIntValue(r.Form, "width"),
		fetchProperIntValue(r.Form, "height"),
		fetchProperStringValue(r.Form, "color"),
	)
	surface.Surface(w, sgo)
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

func fetchProperStringValue(values url.Values, key string) (result string) {
	if val := values.Get(key); val != "" {
		result = val
	}
	return
}
