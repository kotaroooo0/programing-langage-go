package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(max(1, 3, 4, 7))
	fmt.Println(min(1, 3, 4, 7))
	fmt.Println(max())
	fmt.Println(min())

	fmt.Println(max2(1, 3, 4, 7))
	fmt.Println(min2(1, 3, 4, 7))
	fmt.Println(max2())
	fmt.Println(min2())
}

func max(vals ...int) int {
	max := math.MinInt64
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}

func min(vals ...int) int {
	min := math.MaxInt64
	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return min
}

func max2(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("error: empty args")
	}
	max := math.MinInt64
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max, nil
}

func min2(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("error: empty args")
	}
	min := math.MaxInt64
	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return min, nil
}
