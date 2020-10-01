package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	str := "+314124141.5131252"
	fmt.Println(commaBytes(str))
}
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func commaBytes(s string) string {
	var hugo = ""
	if strings.HasPrefix(s, "-") || strings.HasPrefix(s, "+") {
		hugo = s[:1]
		s = s[1:]
	}
	var buf bytes.Buffer
	l := len(s)
	intLength := strings.Index(s, ".")
	if intLength == -1 {
		intLength = l
	}
	tmp := intLength % 3
	buf.WriteString(s[:tmp])
	for tmp < intLength {
		if tmp != 0 {
			buf.WriteString(",")
		}
		buf.WriteString(s[tmp : tmp+3])
		tmp += 3
	}
	if l == intLength {
		return hugo + buf.String()
	}
	buf.WriteString(".")
	return hugo + buf.String() + s[l-intLength+1:]
}
