func isValid(s string) bool {
	match := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	openBrackets := map[rune]struct{}{
		'(': {},
		'{': {},
		'[': {},
	}

	var stack []rune

	for _, char := range s {
		if _, isOpen := openBrackets[char]; isOpen {
			stack = append(stack, char)
			continue
		}

		if expected, isClose := match[char]; isClose {
			if len(stack) == 0 || stack[len(stack)-1] != expected {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
