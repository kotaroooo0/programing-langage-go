package main

import (
	"fmt"
	"strings"
)

func echoFor(args []string) {
	var s, sep string
	for i := 0; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echoJoin(args []string) {
	fmt.Println(strings.Join(args, " "))
}
