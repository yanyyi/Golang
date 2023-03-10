package main

func isValid(s string) bool {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}
			push := string(stack[len(stack)-1])
			stack = stack[0 : len(stack)-1]
			if s[i] == ')' && push == "(" || s[i] == ']' && push == "[" || s[i] == '}' && push == "{" {
				continue
			} else {
				return false
			}
		}
	}
	//判断栈是否为空
	if len(stack) != 0 {
		return false
	}
	return true

}
