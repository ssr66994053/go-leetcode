package P20

func isValid(s string) bool {
	arr := []byte(s)
	stack := []byte{}
	for i := 0; i < len(arr); i++ {
		t := arr[i]
		if len(stack) == 0 {
			stack = append(stack, t)
		} else if '(' == stack[len(stack)-1] && ')' == t {
			stack = stack[0 : len(stack)-1]
		} else if '[' == stack[len(stack)-1] && ']' == t {
			stack = stack[0 : len(stack)-1]
		} else if '{' == stack[len(stack)-1] && '}' == t {
			stack = stack[0 : len(stack)-1]
		} else {
			stack = append(stack, t)
		}
	}

	return len(stack) == 0
}
