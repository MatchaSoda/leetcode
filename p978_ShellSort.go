package leetcode

func sortArray(nums []int) []int {
	for gap := len(nums) >> 1; gap > 0; gap >>= 1 {
		for index := gap; index < len(nums); index++ {
			currentValue := nums[index]
			pendingIndex := index - gap
			for pendingIndex >= 0 && nums[pendingIndex] > currentValue {
				nums[pendingIndex+gap] = nums[pendingIndex]
				pendingIndex -= gap
			}
			nums[pendingIndex+gap] = currentValue
		}
	}
	return nums
}
