package main

import "fmt"

func main() {
	s := []string{"a", "b", "c", "d", "e"}
	fmt.Println(rotate(s, 1))
	fmt.Println(rotate(s, 2))
}

// バブルソートみたいにダッーとswapしていけばできそうだなーと考えたが実装は諦めた
func rotate(s []string, n int) []string {
	k := len(s) - n%len(s)
	s = append(s[k:], s[:k]...)
	return s
}
