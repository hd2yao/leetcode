package day23

func averageOfLevels(root *TreeNode) []float64 {
    if root == nil {
        return []float64{}
    }

    avg := make([]float64, 0)
    queue := []*TreeNode{root}
    for len(queue) > 0 {
        size := len(queue)
        sum := 0.0
        for i := 0; i < size; i++ {
            sum += float64(queue[0].Val)
            if queue[0].Left != nil {
                queue = append(queue, queue[0].Left)
            }
            if queue[0].Right != nil {
                queue = append(queue, queue[0].Right)
            }
            queue = queue[1:]
        }
        avg = append(avg, sum/float64(size))
    }
    return avg
}
