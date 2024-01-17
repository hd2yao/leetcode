package _019_Remove_Nth_Node_from_End_of_List

import (
    "github.com/hd2yao/leetcode/structures"
)

type ListNode = structures.ListNode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummyHead := &ListNode{Next: head}
    slow, fast := dummyHead, dummyHead

    for n > 0 {
        fast = fast.Next
        n--
    }

    for fast.Next != nil {
        fast = fast.Next
        slow = slow.Next
    }
    slow.Next = slow.Next.Next
    return dummyHead.Next
}
