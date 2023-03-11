package kata

func Zeros(n int) int {
	sum := 0

	for i := 5; n/i >= 1; i *= 5 {
		sum += n / i
	}

	return sum
}
