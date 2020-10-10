package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var m = make(map[string]int)

// go run ch01/08/ex08.go https://qiita.com/ | go run ch05/ex2/main.go
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "find: %v\n", err)
		os.Exit(1)
	}
	m := tagToCount(make(map[string]int), doc)
	fmt.Println(m)
}

func tagToCount(m map[string]int, n *html.Node) map[string]int {
	if n == nil {
		return m
	}

	if n.Type == html.ElementNode {
		_, ok := m[n.Data]
		if !ok {
			m[n.Data] = 1
		} else {
			m[n.Data]++
		}
	}

	// 子とお隣を見ている
	m = tagToCount(m, n.FirstChild)
	m = tagToCount(m, n.NextSibling)

	return m
}
