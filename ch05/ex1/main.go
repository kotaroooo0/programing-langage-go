package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// go run ch01/08/ex08.go https://qiita.com/ | go run ch05/ex1/main.go
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "find: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	// 子とお隣を見ている
	links = visit(links, n.FirstChild)
	links = visit(links, n.NextSibling)

	return links

}
