package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price.String())
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price.String())
}

func (db database) add(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	if item == "" || price == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "add requires both item and price\n")
		return
	}
	_, ok := db[item]
	if ok {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "item: %q already exist\n", item)
		return
	}
	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "failed to add new item\n")
		return
	}
	db[item] = dollars(p)
	w.WriteHeader(http.StatusAccepted)
}

func (db database) read(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%q: %s\n", item, price.String())
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	if item == "" || price == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "update requires both item and price\n")
		return
	}
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "item: %q is not exist\n", item)
		return
	}
	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "failed to update price\n")
		return
	}
	db[item] = dollars(p)
	w.WriteHeader(http.StatusAccepted)

}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "delete requires item\n")
		return
	}
	delete(db, item)
	_, ok := db[item]
	if ok {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "failed to delete item: %q\n", item)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
	mux.Handle("/add", http.HandlerFunc(db.add))
	mux.Handle("/read", http.HandlerFunc(db.read))
	mux.Handle("/update", http.HandlerFunc(db.update))
	mux.Handle("/delete", http.HandlerFunc(db.delete))
	log.Fatal(http.ListenAndServe("localhost:10000", mux))
}
