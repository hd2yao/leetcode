package day28

import (
	"container/list"

	"github.com/hd2yao/leetcode/training/structures"
)

type TreeNode = structures.TreeNode

// 迭代法
// 使用 list 包

// 前序遍历
func preorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	stack := list.New()
	stack.PushBack(root)

	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*TreeNode)
		result = append(result, node.Val)

		if node.Right != nil {
			stack.PushBack(node.Right)
		}

		if node.Left != nil {
			stack.PushBack(node.Left)
		}
	}

	return result
}

// 中序遍历
// 使用指针来不断沿着左子树移动
// 如果不使用指针，那么每次移动，都需要从栈里取出刚入栈的元素，在代码层面会变得复杂；使用指针会更加的简单直观
func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	stack := list.New()
	cur := root

	for cur != nil || stack.Len() > 0 {
		for cur != nil {
			stack.PushBack(cur)
			cur = cur.Left
		}
		cur = stack.Remove(stack.Back()).(*TreeNode)
		result = append(result, cur.Val)
		cur = cur.Right
	}

	return result
}

// 后序遍历
func postorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	stack := list.New()
	stack.PushBack(root)

	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*TreeNode)
		result = append(result, node.Val)

		if node.Left != nil {
			stack.PushBack(node.Left)
		}

		if node.Right != nil {
			stack.PushBack(node.Right)
		}
	}

	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]
	}
	return result
}

// 层序遍历
func levelOrderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	queue := list.New()
	queue.PushBack(root)

	if queue.Len() > 0 {
		size := queue.Len()

		for i := 0; i < size; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			result = append(result, node.Val)

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
