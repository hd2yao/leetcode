package day15

func isValid(s string) bool {
    // 字符串长度等于 0 或为奇数比不满足题意
    if len(s) == 0 || len(s)%2 == 1 {
        return false
    }
    // 使用切片来模拟栈
    stack := make([]byte, 0)

    for _, p := range s {
        // 如果是左括号就直接放入
        if p == '(' || p == '[' || p == '{' {
            stack = append(stack, byte(p))
            continue
        }
        // 如果是右括号，需要依次匹配
        switch p {
        case ')':
            // len(stack) == 0 是 ')(' 的情况
            // 因为括号匹配要以正确的顺序闭合，则栈顶元素一定是与之匹配的左括号
            if len(stack) == 0 || stack[len(stack)-1] != '(' {
                return false
            }
            // 匹配成功则出栈
            stack = stack[:len(stack)-1]
        case ']':
            if len(stack) == 0 || stack[len(stack)-1] != '[' {
                return false
            }
            stack = stack[:len(stack)-1]
        case '}':
            if len(stack) == 0 || stack[len(stack)-1] != '{' {
                return false
            }
            stack = stack[:len(stack)-1]
        }
    }

    // 判断是否为全部左括号的情况
    if len(stack) != 0 {
        return false
    }
    return true
}
