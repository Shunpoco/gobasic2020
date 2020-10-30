package main

import (
	"fmt"
	"strings"
)

type tree struct {
	value       int
	left, right *tree
}

func (t *tree) String() string {
	values := appendValues([]int{}, t)

	var strValues []string
	for _, value := range values {
		strValues = append(strValues, fmt.Sprintf("%d", value))
	}
	return strings.Join(strValues, ", ")
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func main() {
	t := add(nil, 10)
	t = add(t, 12)
	t = add(t, 13)

	fmt.Println(t.String())
}
