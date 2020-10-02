package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
)

var url = "http://www.omdbapi.com/?apikey=6927ea3c"

type OmdbapiResult struct {
	Title  string
	Poster string
}

func main() {
	if len(os.Args) < 2 {
		log.Println("You must add movie's title...")
		os.Exit(1)
	}
	movie := os.Args[1]
	res, err := fetch(url + "&t=" + movie)
	if err != nil {
		log.Fatal(err)
	}
	err = download(res.Poster)
	if err != nil {
		log.Fatal(err)
	}
}

func download(url string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return errors.New("not 200")
	}

	_, err = io.Copy(os.Stdout, res.Body)
	if err != nil {
		return err
	}

	return nil
}

func fetch(url string) (*OmdbapiResult, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.New("not 200")
	}

	var result OmdbapiResult
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		res.Body.Close()
		return nil, err
	}
	return &result, nil
}
