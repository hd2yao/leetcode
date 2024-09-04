package day79

// 暴力法 超时
func dailyTemperatures(temperatures []int) []int {
    answer := make([]int, len(temperatures))
    for i := 0; i < len(temperatures); i++ {
        temperature := temperatures[i]
        j := i + 1
        for j < len(temperatures) && temperatures[j] <= temperature {
            j++
        }
        if j != len(temperatures) {
            answer[i] = j - i
        } else {
            answer[i] = 0
        }
    }
    return answer
}

// 单调栈
func dailyTemperaturesStack(temperatures []int) []int {
    stack := make([]int, 0)
    answer := make([]int, len(temperatures))

    for i := 0; i < len(temperatures); i++ {
        for len(stack) != 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
            index := stack[len(stack)-1]
            answer[index] = i - index
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, i)
    }

    return answer
}
