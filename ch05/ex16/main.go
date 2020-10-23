package main

import (
	"fmt"
	"strings"
)

func main() {
	s := []string{"a", "b", "c"}
	fmt.Println(strings.Join(s, "!"))
	fmt.Println(OrenoJoin("!", "afasf", "asfa", "5352"))
}

func OrenoJoin(sep string, elems ...string) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return elems[0]
	}
	n := len(sep) * (len(elems) - 1)
	for i := 0; i < len(elems); i++ {
		n += len(elems[i])
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(elems[0])
	for _, s := range elems[1:] {
		b.WriteString(sep)
		b.WriteString(s)
	}
	return b.String()
}
