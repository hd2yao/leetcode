package day80

func trap(height []int) int {
    var stack []int // 用来存储柱子的索引
    water := 0      // 存储总的雨水量

    for i := 0; i < len(height); i++ {
        // 当前柱子的高度比栈顶柱子的高度大，形成凹槽
        for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
            top := stack[len(stack)-1]   // 栈顶元素的索引，凹槽底部
            stack = stack[:len(stack)-1] // 弹出栈顶元素

            // 如果栈为空，说明此时没有左边界，跳出循环
            if len(stack) == 0 {
                break
            }

            left := stack[len(stack)-1]                                 // 当前凹槽的左边界柱子的索引
            width := i - left - 1                                       // 凹槽的宽度
            boundedHeight := min(height[left], height[i]) - height[top] // 水的高度
            water += width * boundedHeight                              // 水的体积
        }

        // 当前柱子的索引入栈
        // 作为新的左边界，等待匹配
        stack = append(stack, i)
    }

    return water
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
