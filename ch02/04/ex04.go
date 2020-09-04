package popcount

func PopCount2(x uint64) int {
	sum := 0
	for i := 0; i < 64; i++ {
		sum += int(x>>uint(i)) & 1
	}
	return sum
}
