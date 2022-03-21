package leetcode

func removeDuplicates(nums []int) int {
	currentIndex := 1
	targetIndex := 1
	for currentIndex < len(nums) {
		if nums[currentIndex] != nums[currentIndex-1] {
			nums[targetIndex] = nums[currentIndex]
			targetIndex++
			currentIndex++
		} else {
			currentIndex++
		}
	}
	return targetIndex
}
