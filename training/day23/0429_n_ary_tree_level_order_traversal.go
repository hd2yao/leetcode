package day23

type Node struct {
    Val      int
    Children []*Node
}

func levelOrderN(root *Node) [][]int {
    if root == nil {
        return [][]int{}
    }

    result := make([][]int, 0)
    queue := []*Node{root}
    for len(queue) > 0 {
        size := len(queue)
        currentVal := make([]int, 0)
        for i := 0; i < size; i++ {
            currentVal = append(currentVal, queue[0].Val)
            for _, child := range queue[0].Children {
                queue = append(queue, child)
            }
            queue = queue[1:]
        }
        result = append(result, currentVal)
    }

    return result
}
