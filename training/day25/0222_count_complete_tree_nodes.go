package day25

func countNodes(root *TreeNode) int {
    if root == nil {
        return 0
    }

    queue := []*TreeNode{root}
    count := 0
    for len(queue) > 0 {
        size := len(queue)
        count += size
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
    return count
}
