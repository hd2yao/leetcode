package day23

func rightSideView(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }

    result := make([]int, 0)
    queue := []*TreeNode{root}
    for len(queue) > 0 {
        size := len(queue)
        result = append(result, queue[len(queue)-1].Val)
        for i := 0; i < size; i++ {
            if queue[0].Left != nil {
                queue = append(queue, queue[0].Left)
            }
            if queue[0].Right != nil {
                queue = append(queue, queue[0].Right)
            }
            queue = queue[1:]
        }
    }

    return result
}
