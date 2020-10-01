package main

import (
	"bytes"
	"fmt"
)

func main() {
	str := "hogefugapiyo"
	fmt.Println(str[:3])
	fmt.Println(str[4:])
	fmt.Println(commaBytes(str))
	fmt.Println(comma(str))
}
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func commaBytes(s string) string {
	var buf bytes.Buffer
	l := len(s)
	tmp := l % 3
	buf.WriteString(s[:tmp])
	for tmp < l {
		if tmp != 0 {
			buf.WriteString(",")
		}
		buf.WriteString(s[tmp : tmp+3])
		tmp += 3
	}
	return buf.String()
}
