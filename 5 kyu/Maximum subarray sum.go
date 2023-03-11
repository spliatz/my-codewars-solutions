package kata

func MaximumSubarraySum(numbers []int) int {
	var cursum int
	var maxsum int

	for end := 0; end < len(numbers); end++ {
		cursum += numbers[end]
		if cursum > maxsum {
			maxsum = cursum
		}
		if cursum < 0 {
			cursum = 0
		}
	}
	return maxsum
}
