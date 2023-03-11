package kata

func TwoSum(nums []int, target int) [2]int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == target-nums[i] {
				return [2]int{i, j}
			}
		}
	}
	return [2]int{}
}
