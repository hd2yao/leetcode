package day23

func levelOrderBottom(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }

    result := make([][]int, 0)
    queue := []*TreeNode{root}
    for len(queue) > 0 {
        size := len(queue)
        currentVal := make([]int, 0)
        for i := 0; i < size; i++ {
            currentVal = append(currentVal, queue[0].Val)
            if queue[0].Left != nil {
                queue = append(queue, queue[0].Left)
            }
            if queue[0].Right != nil {
                queue = append(queue, queue[0].Right)
            }
            queue = queue[1:]
        }
        result = append(result, currentVal)
    }

    // 反转结果
    for i := 0; i < len(result)/2; i++ {
        result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]
    }

    return result
}
