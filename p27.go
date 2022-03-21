package leetcode

func removeElement(nums []int, val int) int {
	currentIndex := 0
	targetIndex := 0
	for currentIndex < len(nums) {
		if nums[currentIndex] != val {
			nums[targetIndex] = nums[currentIndex]
			targetIndex++
			currentIndex++
		} else {
			currentIndex++
		}
	}
	return targetIndex
}
