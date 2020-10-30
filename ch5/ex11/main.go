package main

import (
	"fmt"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operationg systems"},
	"operationg systems":    {"data structures", "computer organizations"},
	"programming languages": {"data structures", "computer organizations"},
	"linear algebra":        {"calculus"},
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)

	var checkChild func([]string, string)
	checkChild = func(items []string, s string) {
		for _, item := range items {
			if item == s {
				fmt.Printf("cyclic!: %s\n", item)
				return
			}
			checkChild(m[item], s)
		}
	}

	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			} else {
				checkChild(m[item], item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)
	return order
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}
