package popcount

func PopCount3(x uint64) int {
	sum := 0
	for x > 0 {
		x = x & (x - 1)
		sum += 1
	}
	return sum
}
