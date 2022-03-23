package leetcode

func longestCommonPrefix(strs []string) string {
	for charIndex, currentChar := range []byte(strs[0]) {
		for _, currentString := range strs {
			if charIndex == len(currentString) {
				return currentString
			}
			if currentString[charIndex] != currentChar {
				return currentString[:charIndex]
			}
		}
	}
	return strs[0]
}
