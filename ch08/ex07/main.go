package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kotaroooo0/programing-language-go/ch08/ex06/links"
)

// TODO: やってない
func main() {
	worklist := make(chan []string)
	unseenLinks := make(chan string)

	links := make([]string, len(os.Args)-1)
	go func() {
		worklist <- links
	}()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				urls := crawl(link)
				go func() {
					worklist <- urls
				}()
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}

func crawl(link string) []string {
	fmt.Println(link)
	list, err := links.Extract(link)
	if err != nil {
		log.Print(err)
	}
	return list
}
