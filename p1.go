package leetcode

func twoSum(nums []int, target int) []int {
	existedNumbers := map[int]int{} // value index
	for currentIndex, currentNumber := range nums {
		if targetIndex, ok := existedNumbers[target-currentNumber]; ok {
			return []int{targetIndex, currentIndex}
		}
		existedNumbers[currentNumber] = currentIndex
	}
	return nil
}
