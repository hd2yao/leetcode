package _206_Reverse_Linked_List

import "github.com/hd2yao/leetcode/structures"

type ListNode = structures.ListNode

// 虚拟头结点，头插法
func reverseListWithDummy(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    dummyHead := &ListNode{Next: head}
    cur := dummyHead.Next
    for cur.Next != nil {
        dummyHead.Next = cur.Next
        cur.Next = cur.Next.Next
        dummyHead.Next.Next = head
        head = dummyHead.Next
    }
    return dummyHead.Next
}

// 双指针法(迭代法)
func reverseListTwoPointer(head *ListNode) *ListNode {
    var pre *ListNode
    for head != nil {
        tmp := head.Next
        head.Next = pre
        pre = head
        head = tmp
    }
    return pre
}

// 递归法：根据双指针的思路再来写递归
func reverseListRecursive(head *ListNode) *ListNode {
    return reverse(head, nil)
}

func reverse(cur, pre *ListNode) *ListNode {
    if cur == nil {
        return pre
    }
    tmp := cur.Next
    cur.Next = pre
    return reverse(tmp, cur)
}
