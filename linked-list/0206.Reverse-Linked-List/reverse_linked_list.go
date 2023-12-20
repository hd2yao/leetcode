package _206_Reverse_Linked_List

import "github.com/hd2yao/leetcode/structures"

type ListNode = structures.ListNode

// 虚拟头结点，头插法
func reverseListWithDummy(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    cur := head
    dummyHead := &ListNode{Next: head}
    for cur.Next != nil {
        dummyHead.Next = cur.Next
        cur.Next = cur.Next.Next
        dummyHead.Next.Next = head
        head = dummyHead.Next
    }

    return head
}
