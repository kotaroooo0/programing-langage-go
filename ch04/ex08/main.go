package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	categories := make(map[string]int)
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		for c, t := range unicode.Categories {
			if unicode.Is(t, r) {
				categories[c]++
				continue
			}
		}
	}
	fmt.Printf("\ncategory\tcount\n")
	for c, n := range categories {
		fmt.Printf("%s\t%d\n", string(c), n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
