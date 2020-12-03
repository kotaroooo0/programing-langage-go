package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/kotaroooo0/programing-language-go/ch08/ex06/links"
)

// TODO: 正常に終了しない?
func main() {
	depth := flag.Int64("depth", 3, "depth limit of crawling")
	flag.Parse()

	worklist := make(chan []Link)
	unseenLinks := make(chan Link)

	links := make([]Link, len(os.Args)-1)
	for i, url := range os.Args[1:] {
		links[i] = Link{
			Depth: 0,
			URL:   url,
		}
	}

	go func() {
		worklist <- links
	}()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				if link.Depth >= *depth {
					return
				}
				urls := crawl(link)
				foundLinks := make([]Link, len(urls))
				for j, url := range urls {
					foundLinks[j] = Link{
						Depth: link.Depth + 1,
						URL:   url,
					}
				}
				go func() {
					worklist <- foundLinks
				}()
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link.URL] {
				seen[link.URL] = true
				unseenLinks <- link
			}
		}
	}
}

func crawl(link Link) []string {
	fmt.Println(link.URL)
	fmt.Println(link.Depth)
	list, err := links.Extract(link.URL)
	if err != nil {
		log.Print(err)
	}
	return list
}

type Link struct {
	Depth int64
	URL   string
}
