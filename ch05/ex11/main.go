package main

import (
	"fmt"
	"log"
)

var prereqs = map[string]map[string]bool{
	"algorithms":     {"data structures": true},
	"calculus":       {"linear algebra": true},
	"linear algebra": {"calculus": true},

	"compilers": {
		"data structures":       true,
		"formal languages":      true,
		"computer organization": true,
	},

	"data structures":       {"discrete math": true},
	"databases":             {"data structures": true},
	"discrete math":         {"intro to programming": true},
	"formal languages":      {"discrete math": true},
	"networks":              {"operating systems": true},
	"operating systems":     {"data structures": true, "computer organization": true},
	"programming languages": {"data structures": true, "computer organization": true},
}

func main() {
	sorted, err := TopoSort(prereqs)
	if err != nil {
		log.Fatalln(err)
	}
	for i, course := range sorted {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func TopoSort(m map[string]map[string]bool) ([]string, error) {
	var order []string
	seen := make(map[string]bool)
	// visited := make(map[string]bool)
	var visitAll func(items map[string]bool) error

	visitAll = func(items map[string]bool) error {
		// fmt.Println(visited)
		for item := range items {

			// if visited[item] {
			// 	return fmt.Errorf("error: closed!")
			// }
			if !seen[item] {
				// visited[item] = true
				seen[item] = true
				if err := visitAll(m[item]); err != nil {
					return err
				}
				order = append(order, item)
			}
		}
		return nil
	}

	keys := make(map[string]bool)
	for key := range m {
		keys[key] = true
	}

	if err := visitAll(keys); err != nil {
		return nil, err
	}
	return order, nil
}
