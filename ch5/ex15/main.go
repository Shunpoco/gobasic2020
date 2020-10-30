package main

import "fmt"

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func main() {
	fmt.Println(sum())
	fmt.Println(sum(1, 2, 3, 4, 5))

	fmt.Println(max())
	fmt.Println(max(100, 10, -1))

	fmt.Println(min())
	fmt.Println(min(100, 10, -1))

	fmt.Println(maxRequiredVals(100, 10, -1))
	fmt.Println(maxRequiredVals())
}

func max(vals ...int) int {
	max := 0
	for _, val := range vals {
		if max < val {
			max = val
		}
	}
	return max
}

func min(vals ...int) int {
	min := 0
	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return min
}

func maxRequiredVals(vals ...int) (max int) {
	if len(vals) < 1 {
		panic("maxRequiredVals requires at least one variable...")
	}

	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return
}
