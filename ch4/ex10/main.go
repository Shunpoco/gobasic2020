package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"ch4/ex10/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	now := time.Now()
	aMonthAgo := now.AddDate(0, -1, 0)
	aYearAgo := now.AddDate(-1, 0, 0)
	var lessThanAMonth []*github.Issue
	var lessThanAYear []*github.Issue
	var moreThanAYear []*github.Issue

	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		if aMonthAgo.Sub(item.CreatedAt) < 0 {
			lessThanAMonth = append(lessThanAMonth, item)
			continue
		}
		if aYearAgo.Sub(item.CreatedAt) < 0 {
			lessThanAYear = append(lessThanAYear, item)
			continue
		}
		moreThanAYear = append(moreThanAYear, item)
	}
	printer(lessThanAMonth, "less than a month")
	printer(lessThanAYear, "less than a year")
	printer(moreThanAYear, "more than a year")
}

func printer(issues []*github.Issue, name string) {
	fmt.Printf("issues of %s\n", name)
	fmt.Print("-------------------------------------\n")
	for _, item := range issues {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
	fmt.Print("-------------------------------------\n")
}
