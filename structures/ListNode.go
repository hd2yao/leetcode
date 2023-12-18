package structures

import "fmt"

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// List2Ints convert List to []int
func List2Ints(head *ListNode) []int {
	// 链表深度限制，超出次限制，会 panic
	limit := 100
	times := 0

	var result []int
	for head != nil {
		times++
		if times > limit {
			msg := fmt.Sprintf("链表深度超过 %d，可能出现环状链表。请检查错误，或者放宽 l2s 函数中 limit 的限制。", limit)
			panic(msg)
		}
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

// Ints2List convert []int to List
func Ints2List(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{}
	cur := head
	for _, val := range nums {
		cur.Next = &ListNode{Val: val}
		cur = cur.Next
	}
	return head.Next
}
