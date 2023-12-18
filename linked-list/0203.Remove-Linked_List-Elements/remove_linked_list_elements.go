package _203_Remove_Linked_List_Elements

import "github.com/hd2yao/leetcode/structures"

// ListNode Definition for singly-linked list.
type ListNode = structures.ListNode

// 原链表删除元素
func removeElements(head *ListNode, val int) *ListNode {
    for head != nil && head.Val == val {
        head = head.Next
    }
    tmp := head
    for tmp != nil && tmp.Next != nil {
        if tmp.Next.Val == val {
            tmp.Next = tmp.Next.Next
        } else {
            tmp = tmp.Next
        }
    }
    return head
}

// 虚拟头结点删除元素
func removeElementsDummy(head *ListNode, val int) *ListNode {
    dummyHead := &ListNode{
        Next: head,
    }
    head = dummyHead
    for head.Next != nil {
        if head.Next.Val == val {
            head.Next = head.Next.Next
        } else {
            head = head.Next
        }
    }
    // 不能返回 head 是因为
    // 1.上面直接使用 head 来进行遍历
    // 2.若使用一个新的临时指针来遍历，也不能 return head，因为当 head.Val == val, head 就会被删掉
    return dummyHead.Next
}
