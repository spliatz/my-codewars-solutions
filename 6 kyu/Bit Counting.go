package kata

import "strconv"

func CountBits(number uint) int {
	sum := 0
	str := strconv.FormatUint(uint64(number), 2)
	for _, elem := range str {
		if elem == '1' {
			sum++
		}
	}
	return sum
}
