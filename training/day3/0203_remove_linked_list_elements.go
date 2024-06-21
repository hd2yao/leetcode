package day3

import "github.com/hd2yao/leetcode/training/structures"

// 虚拟头节点
func removeElementsDummy(head *structures.ListNode, val int) *structures.ListNode {
    dummyHead := &structures.ListNode{Next: head}
    head = dummyHead
    for head.Next != nil {
        if head.Next.Val == val {
            head.Next = head.Next.Next
        } else {
            head = head.Next
        }
    }
    return dummyHead.Next
}

// 临时节点
func removeElementsTmp(head *structures.ListNode, val int) *structures.ListNode {
    tmp := head
    for tmp != nil && tmp.Next != nil {
        if tmp.Next.Val == val {
            tmp.Next = tmp.Next.Next
        } else {
            tmp = tmp.Next
        }
    }
    if head != nil && head.Val == val {
        return head.Next
    }
    return head
}
