package main

import "fmt"

func hoge() (fuga int) {
	defer func() {
		recover()
	}()
	fuga = 423432 + 532532
	panic(0)
}

func main() {
	fmt.Println(hoge())
}
