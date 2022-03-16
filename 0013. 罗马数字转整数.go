package leetcode

var romanCharValues = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func romanToInt(s string) int {
	result := 0
	for charIndex := range s {
		currentValue := romanCharValues[s[charIndex]]
		if charIndex < len(s)-1 && currentValue < romanCharValues[s[charIndex+1]] {
			result -= currentValue
		} else {
			result += currentValue
		}
	}
	return result
}
