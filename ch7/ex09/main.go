package main

import (
	"html/template"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"
)

var tempHTML = template.Must(template.ParseFiles("./index.html"))

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

// 先に渡したcolumn名を優先としてソートします
func DoubleSort(tracks []*Track, columns []string) sort.Interface {
	return customSort{
		tracks,
		func(x, y *Track) bool {
			for _, column := range columns {
				if column == "Title" && x.Title != y.Title {
					return x.Title < y.Title
				} else if column == "Artist" && x.Artist != y.Artist {
					return x.Artist < y.Artist
				} else if column == "Album" && x.Album != y.Album {
					return x.Album < y.Album
				} else if column == "Year" && x.Year != y.Year {
					return x.Year < y.Year
				} else if column == "Length" && x.Length != y.Length {
					return x.Length < y.Length
				}
			}
			return true
		},
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:10000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	var ss string
	for key, val := range query {
		log.Printf("%s: %s", key, val)
		if key == "sort" && len(val) > 0 {
			columns := strings.Split(val[0], ",")
			sort.Sort(DoubleSort(tracks, columns))
			ss = columns[0]
		}
	}
	op := Options{tracks, ss}
	err := tempHTML.ExecuteTemplate(w, "index.html", op)
	if err != nil {
		log.Fatal(err)
	}
}

type Options struct {
	Tracks     []*Track
	SecondSort string
}
