func isValid(s string) bool {

	closeBrackets := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	var stack []rune

	for _, char := range s {
		item, ok := closeBrackets[char]
		if ok {
			if len(stack) == 0 || pop(&stack) != item {
				return false
			}
		} else {
			push(char, &stack)
		}

	}

	return len(stack) == 0
}

func pop(stack *[]rune) rune {
	elem := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]

	return elem
}

func push(char rune, stack *[]rune) {
	*stack = append(*stack, char)
}
