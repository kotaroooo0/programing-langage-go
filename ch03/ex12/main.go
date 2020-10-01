package main

import "fmt"

func main() {
	fmt.Println(IsAnagram("xyz", "xyz"))
	fmt.Println(IsAnagram("xyz", "xyaz"))
	fmt.Println(IsAnagram("xyz", "zxy"))
	fmt.Println(IsAnagram("kjkj", "kjkjkj"))
}

func IsAnagram(x, y string) bool {
	if len(x) != len(y) {
		return false
	}

	var xMap = make(map[byte]int)
	var yMap = make(map[byte]int)
	for i := 0; i < len(x); i++ {
		xMap[x[i]]++
		yMap[y[i]]++
	}

	if len(xMap) != len(yMap) {
		return false
	}
	for k, _ := range xMap {
		if xMap[k] != yMap[k] {
			return false
		}
	}
	return true
}
