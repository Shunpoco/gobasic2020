package main

import (
	"fmt"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organizations",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operationg systems"},
	"operationg systems":    {"data structures", "computer organizations"},
	"programming languages": {"data structures", "computer organizations"},
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var visitAllMap func(map[string][]string, string)
	visitAllMap = func(itemMap map[string][]string, s string) {
		if s != "" {
			visitAll(itemMap[s])
		} else {
			for k := range itemMap {
				if !seen[k] {
					seen[k] = true
					visitAllMap(itemMap, k)
					order = append(order, k)
				}
			}
		}

	}

	// var keys []string
	// for key := range m {
	// 	keys = append(keys, key)
	// }
	// sort.Strings(keys)
	// visitAll(keys)
	visitAllMap(m, "")
	return order
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}
