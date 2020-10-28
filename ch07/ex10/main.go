package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(IsPalindrome(Strings{"a", "b", "b", "a"}))
	fmt.Println(IsPalindrome(Strings{"a", "b", "b", "a"}))
	fmt.Println(IsPalindrome(Strings{"a", "b", "c", "b", "a"}))
	fmt.Println(IsPalindrome(Strings{"a", "b", "c", "b", "a", "a"}))
}

func IsPalindrome(s sort.Interface) bool {
	l := s.Len()
	half := l / 2
	for i := 0; i < half; i++ {
		j := l - i - 1
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}
	return true
}

type Strings []string

func (s Strings) Len() int {
	return len(s)
}

func (s Strings) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s Strings) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
