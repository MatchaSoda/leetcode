package leetcode

func binarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (right-left)/2 + left
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return left
}

func sortArray(nums []int) []int {
	for index := 1; index < len(nums); index++ {
		currentValue := nums[index]
		position := binarySearch(nums[0:index-1], currentValue)
		for pendingIndex := index - 1; pendingIndex >= position; pendingIndex-- {
			nums[pendingIndex+1] = nums[pendingIndex]
		}
		nums[position] = currentValue
	}
	return nums
}
