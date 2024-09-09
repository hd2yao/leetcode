package _181_Merge_Nodes_in_between_Zeros

import "github.com/hd2yao/leetcode/structures"

type ListNode = structures.ListNode

func mergeNodes(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	slow := dummy

	for head.Next != nil {
		if head.Val == 0 {
			slow = slow.Next
			slow.Val = 0
		} else {
			slow.Val += head.Val
		}
		head = head.Next
	}
	slow.Next = nil

	return dummy.Next
}
