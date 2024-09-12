package day2

import "github.com/hd2yao/leetcode/structures"

type ListNode = structures.ListNode

func removeElementsDummy(head *ListNode, val int) *ListNode {
    dummy := &ListNode{Next: head}
    head = dummy
    for head.Next != nil {
        if head.Next.Val == val {
            head.Next = head.Next.Next
        } else {
            head = head.Next
        }
    }
    return dummy.Next
}
