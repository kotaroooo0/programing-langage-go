package main

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	fmt.Println(CountWordsAndImage("https://hub.docker.com/"))
}

func CountWordsAndImage(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImage(doc)
	return
}

func countWordsAndImage(n *html.Node) (int, int) {
	return visit(n, 0, 0)
}

func visit(n *html.Node, words, images int) (int, int) {
	if n == nil {
		return words, images
	}

	if n.Type == html.ElementNode && n.Data == "img" {
		fmt.Println("GAsgghlghaslugjhasilhglaghsilshu")
		images++
	}

	words += len(strings.Split(n.Data, " "))

	// 子とお隣を見ている
	words, images = visit(n.FirstChild, words, images)
	words, images = visit(n.NextSibling, words, images)

	return words, images
}
