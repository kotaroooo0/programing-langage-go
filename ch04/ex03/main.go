package main

import (
	"fmt"
)

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	reverse(&a)
	fmt.Println(a)
}

func reverse(s *[10]int) {
	for i, j := 0, 9; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
