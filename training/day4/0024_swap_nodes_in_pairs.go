package day4

import "github.com/hd2yao/leetcode/training/structures"

type ListNode = structures.ListNode

// 双指针-迭代法
func swapPairsIteration(head *ListNode) *ListNode {
    dummyHead := &ListNode{Next: head}
    front := dummyHead
    for head != nil && head.Next != nil {
        // 相邻节点进行交换
        tmp := head.Next
        head.Next = tmp.Next
        tmp.Next = head
        front.Next = tmp

        // 指针向后移动
        front = head
        head = head.Next
    }
    return dummyHead.Next
}

func swapPairs(head *ListNode) *ListNode {
    dummyHead := &ListNode{Next: head}
    front := dummyHead
    swapPairsRecursive(front, head)
    return dummyHead.Next
}

func swapPairsRecursive(front, head *ListNode) {
    if head != nil && head.Next != nil {
        tmp := head.Next
        head.Next = tmp.Next
        tmp.Next = head
        front.Next = tmp

        swapPairsRecursive(head, head.Next)
    }
}
