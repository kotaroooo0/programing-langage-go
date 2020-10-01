package main

import "fmt"

func main() {
	s := []string{"a", "b", "c", "c", "c", "b", "b"}
	fmt.Println(hoge(s))
	fmt.Println(s)
}

func hoge(strings []string) []string {
	i := 1
	l := len(strings)
	for j := 0; j < l; j++ {
		if j == 0 {
			continue
		}
		if strings[i] == strings[i-1] {
			for k := i; k < l-1; k++ {
				strings[k] = strings[k+1]
			}
		} else {
			i++
		}
	}
	return strings[:i]
}
