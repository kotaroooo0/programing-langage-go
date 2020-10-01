package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/golang/go/src/log"
)

func main() {
	wordToCount := make(map[string]int)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		wordToCount[scanner.Text()]++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("word\tcount\n")
	for w, c := range wordToCount {
		fmt.Printf("%s\t%d\n", w, c)
	}
}
