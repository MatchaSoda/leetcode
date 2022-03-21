package leetcode

func searchInsert(nums []int, target int) int {
	//Binary Search
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (right-left)/2 + left
		if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			return mid
		}
	}
	return left
}
