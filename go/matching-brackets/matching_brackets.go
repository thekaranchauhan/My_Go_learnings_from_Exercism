package brackets

var pairs = map[rune]rune{')': '(', ']': '[', '}': '{'}

func Bracket(text string) bool {
	stack := make([]rune, 0, 4)
	for _, r := range text {
		if r == '(' || r == '{' || r == '[' {
			stack = append(stack, r)
		} else if opener, ok := pairs[r]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != opener {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0

}
