package day15

import "strings"

func removeDuplicates(s string) string {
    ss := []byte(s)
    result := make([]string, 0)
    for _, alpha := range ss {
        // 栈不空 且 与栈顶元素不等
        if len(result) != 0 && result[len(result)-1] == string(alpha) {
            result = result[:len(result)-1]
        } else {
            // 入栈
            result = append(result, string(alpha))
        }
    }

    return strings.Join(result, "")
}
