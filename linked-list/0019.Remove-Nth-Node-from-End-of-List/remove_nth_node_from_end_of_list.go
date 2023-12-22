package _019_Remove_Nth_Node_from_End_of_List

import (
	"github.com/hd2yao/leetcode/structures"
)

type ListNode = structures.ListNode

// 快慢指针+虚拟头结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}
	slow, fast := dummyHead, dummyHead
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next

	// 可使用 head 代替 fast 指针
	//slow := dummyHead
	//for i := 0; i < n-1; i++ {
	//	head = head.Next
	//}
	//for head.Next != nil {
	//	head = head.Next
	//	slow = slow.Next
	//}
	//slow.Next = slow.Next.Next
	return dummyHead.Next
}
