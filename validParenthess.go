package main

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']',
// determine if the input string is valid.

// Input:	s = "()"
// Output:	true

// Input:	s = "()[]{}"
// Output:	true

// Input:	s = "(]"
// Output:	false

func main() {
	s := "()"
	println(isValid(s))
}

// 0ms
// 1.9MB
func isValid(s string) bool {
	stack := make([]byte, len(s))
	head := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(', '[', '{':
			stack[head] = s[i]
			head++
		case ')':
			if head == 0 {
				return false
			}
			head--
			if stack[head] != '(' {
				return false
			}
		case ']':
			if head == 0 {
				return false
			}
			head--
			if stack[head] != '[' {
				return false
			}
		case '}':
			if head == 0 {
				return false
			}
			head--
			if stack[head] != '{' {
				return false
			}
		}
	}
	return head == 0
}