package day23

import "github.com/hd2yao/leetcode/training/structures"

type TreeNode = structures.TreeNode

func levelOrder(root *TreeNode) [][]int {
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

	return result
}
