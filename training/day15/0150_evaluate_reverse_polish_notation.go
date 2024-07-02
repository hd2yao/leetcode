package day15

import "strconv"

func evalRPN(tokens []string) int {
    stack := make([]int, 0)
    for i := 0; i < len(tokens); i++ {
        if tokens[i] != "+" && tokens[i] != "-" && tokens[i] != "*" && tokens[i] != "/" {
            num, _ := strconv.Atoi(tokens[i])
            stack = append(stack, num)
            continue
        }
        num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
        stack = stack[:len(stack)-2]
        switch tokens[i] {
        case "+":
            stack = append(stack, num1+num2)
        case "-":
            stack = append(stack, num1-num2)
        case "*":
            stack = append(stack, num1*num2)
        case "/":
            stack = append(stack, num1/num2)
        }
    }
    return stack[0]
}
