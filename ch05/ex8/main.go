package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	resp, err := http.Get("http://www.gopl.io/")
	if err != nil {
		return
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}

	fmt.Println(ElementByID(doc, "toc"))
}

func forEachNode(n *html.Node, pre, post func(*html.Node) bool) {
	if pre != nil {
		if ok := pre(n); !ok {
			return
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		if ok := post(n); !ok {
			return
		}
	}
}

func ElementByID(doc *html.Node, id string) *html.Node {
	var node *html.Node

	forEachNode(doc, func(n *html.Node) bool {
		if n.Type != html.ElementNode {
			return true
		}

		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				node = n
				return false
			}
		}
		return true
	}, nil)

	return node
}
