package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	fmt.Println(hoge)
}

func hoge(x, y []byte) int {
	sx := sha256.Sum256(x)
	sy := sha256.Sum256(y)
	sum := 0
	for i := 0; i < 32; i++ {
		xor := sx[i] ^ sy[i]
		for xor > 0 {
			xor = xor & (xor - 1)
			sum += 1
		}
	}
	return sum
}
