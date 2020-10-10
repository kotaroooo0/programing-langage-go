package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var m = make(map[string]int)

// go run ch01/08/ex08.go https://qiita.com/ | go run ch05/ex3/main.go
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "find: %v\n", err)
		os.Exit(1)
	}
	texts := visit(nil, doc)
	for i := 0; i < len(texts); i++ {
		fmt.Println(texts[i])
	}
}

func visit(texts []string, n *html.Node) []string {
	if n == nil {
		return texts
	}

	if n.Type == html.TextNode {
		texts = append(texts, n.Data)
	}

	// 子とお隣を見ている
	texts = visit(texts, n.FirstChild)
	texts = visit(texts, n.NextSibling)

	return texts
}
