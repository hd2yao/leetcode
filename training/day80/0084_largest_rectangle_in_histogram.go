package day80

func largestRectangleArea(heights []int) int {
    stack := make([]int, 0)
    area := 0

    for i := 0; i <= len(heights); i++ {
        curHeight := 0
        if i < len(heights) {
            curHeight = heights[i]
        }
        for len(stack) > 0 && curHeight < heights[stack[len(stack)-1]] {
            h := heights[stack[len(stack)-1]]
            stack = stack[:len(stack)-1]

            width := 0
            if len(stack) == 0 {
                width = 1
            } else {
                width = i - stack[len(stack)-1] - 1
            }

            area = max(area, h*width)
        }
        stack = append(stack, i)
    }

    return area
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}
