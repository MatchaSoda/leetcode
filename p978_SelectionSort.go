package leetcode

func sortArray(nums []int) []int {
	for pendingIndex, _ := range nums[:len(nums)-1] {
		minIndex := pendingIndex
		for currentIndex, currentValue := range nums[pendingIndex+1:] {
			if nums[minIndex] > currentValue {
				minIndex = currentIndex + pendingIndex + 1
			}
		}
		nums[pendingIndex], nums[minIndex] = nums[minIndex], nums[pendingIndex]
	}
	return nums
}
