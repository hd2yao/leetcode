package day30

import (
    "container/list"
    "github.com/hd2yao/leetcode/training/structures"
)

type TreeNode = structures.TreeNode

func findBottomLeftValue(root *TreeNode) int {
    levelAllValue := make([][]int, 0)

    queue := list.New()
    queue.PushBack(root)

    for queue.Len() > 0 {
        size := queue.Len()
        levelValue := make([]int, 0)
        for i := 0; i < size; i++ {
            node := queue.Remove(queue.Front()).(*TreeNode)
            levelValue = append(levelValue, node.Val)
            if node.Left != nil {
                queue.PushBack(node.Left)
            }
            if node.Right != nil {
                queue.PushBack(node.Right)
            }
        }
        levelAllValue = append(levelAllValue, levelValue)
    }

    return levelAllValue[len(levelAllValue)-1][0]
}

// 改进
func findBottomLeftValueImprove(root *TreeNode) int {
    result := 0

    queue := list.New()
    queue.PushBack(root)

    for queue.Len() > 0 {
        size := queue.Len()

        for i := 0; i < size; i++ {
            node := queue.Remove(queue.Front()).(*TreeNode)
            if i == 0 {
                result = node.Val
            }
            if node.Left != nil {
                queue.PushBack(node.Left)
            }
            if node.Right != nil {
                queue.PushBack(node.Right)
            }
        }
    }

    return result
}
