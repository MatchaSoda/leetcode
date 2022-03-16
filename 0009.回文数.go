package leetcode

func isPalindrome(input int) bool {
	if input < 0 {
		return false
	}
	originalInput := input
	reversedInput := 0
	for input > 0 {
		reversedInput = reversedInput*10 + input%10
		input /= 10
	}
	return originalInput == reversedInput
}
