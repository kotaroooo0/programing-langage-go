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

	forEachNode(doc, startElement, endElement)
}

// io.Writerを引数に渡すと幸せそう
func forEachNode(n *html.Node, pre, post func(*html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s/ %s>\n", depth*2, "", n.Data, n.Attr)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
	}
}
