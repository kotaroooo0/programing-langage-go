package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(expand("I am $kotaroooo0 and $snowboard", func(s string) string {
		return s + "+"
	}))
}

// EX "I am $kotaroooo0" -> "I am f("kotaroooo0")"
func expand(s string, f func(string) string) string {
	splited := strings.Split(s, " ")
	for i, v := range splited {
		if strings.HasPrefix(v, "$") {
			splited[i] = f(v[1:])
		}
	}
	return strings.Join(splited, " ")
}
