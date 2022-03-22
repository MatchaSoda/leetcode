package leetcode

func sortArray(nums []int) []int {
	lastSwapIndex := len(nums)
	for lastSwapIndex > 0 {
		currentSwapIndex := 0
		for index := 1; index < lastSwapIndex; index++ {
			if nums[index] < nums[index-1] {
				nums[index], nums[index-1] = nums[index-1], nums[index]
				currentSwapIndex = index
			}
		}
		lastSwapIndex = currentSwapIndex
	}
	return nums
}
