package _020_Valid_Parentheses

func isValid(s string) bool {
    if len(s) == 0 {
        return true
    }
    stack := make([]byte, 0)
    for i := 0; i < len(s); i++ {
        if s[i] == '(' || s[i] == '{' || s[i] == '[' {
            stack = append(stack, s[i])
        } else if (len(stack) > 0 && s[i] == ')' && stack[len(stack)-1] == '(') ||
            (len(stack) > 0 && s[i] == '}' && stack[len(stack)-1] == '{') ||
            (len(stack) > 0 && s[i] == ']' && stack[len(stack)-1] == '[') {
            stack = stack[:len(stack)-1]
        } else {
            return false
        }
    }
    return len(stack) == 0
}
