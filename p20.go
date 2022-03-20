package leetcode

func isValid(s string) bool {
	leftParenthesis := make([]byte, 0)
	for _, currentChar := range []byte(s) {
		if currentChar == '(' || currentChar == '[' || currentChar == '{' {
			leftParenthesis = append(leftParenthesis, currentChar)
		} else if len(leftParenthesis) == 0 {
			return false
		} else {
			matchingChar := leftParenthesis[len(leftParenthesis)-1]
			if currentChar == ')' && matchingChar != '(' {
				return false
			} else if currentChar == ']' && matchingChar != '[' {
				return false
			} else if currentChar == '}' && matchingChar != '{' {
				return false
			}
			leftParenthesis = leftParenthesis[:len(leftParenthesis)-1]
		}
	}
	return len(leftParenthesis) == 0
}
