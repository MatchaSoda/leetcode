package leetcode

func longestCommonPrefix(strs []string) string {
	for charIndex := 0; charIndex < len(strs[0]); charIndex++ {
		for stringIndex := 0; stringIndex < len(strs); stringIndex++ {
			if charIndex == len(strs[stringIndex]) || strs[stringIndex][charIndex] != strs[0][charIndex] {
				return strs[0][:charIndex]
			}
		}
	}
	return strs[0]
}
