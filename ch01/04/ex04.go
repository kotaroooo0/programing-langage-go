package main

import (
	"bufio"
	"os"
	"strings"

	"golang.org/x/exp/errors/fmt"
)

// 考えたこと
// 1. stringで足して行って、splitする -> 重複の処理が面倒くさそう
// 2. string -> []stringのmapを作る -> できない?!
// 3. 構造体でやる -> 初期化面倒

// 課題
// Setの出力がランダムになる
// 入力順にソートした方がいい

type Data struct {
	Count int
	Names map[string]struct{}
}

func main() {
	words := make(map[string]*Data)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, words)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, words)
			f.Close()
		}
	}
	for line, w := range words {
		if w.Count > 1 {
			fmt.Printf("%d\t%s\t", w.Count, line)

			// SetをPrint
			keys := make([]string, 0, len(w.Names))
			for w := range w.Names {
				keys = append(keys, w)
			}
			fmt.Println(strings.Join(keys, ","))
		}
	}
}

func countLines(f *os.File, words map[string]*Data) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		t := input.Text()
		if words[t] == nil {
			words[t] = &Data{
				Count: 0,
				Names: make(map[string]struct{}),
			}
		}
		words[t].Count++
		words[t].Names[f.Name()] = struct{}{}
	}
}
