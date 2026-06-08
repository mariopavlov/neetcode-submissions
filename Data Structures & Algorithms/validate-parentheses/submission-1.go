func isValid(s string) bool {
	closeBrackets := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	var stack []rune

	for _, char := range s {
		val, ok := closeBrackets[char]
		if ok {
			if len(stack) == 0 {
				return false
			}

			topElement := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if topElement != val {
				return false
			}
		} else {
			stack = append(stack, char)
		}
	}

	if len(stack) == 0 {
		return true
	}

	return false
}
