package day3

import "github.com/hd2yao/leetcode/training/structures"

// 双指针法-迭代法
func reverseList(head *structures.ListNode) *structures.ListNode {
    var prev *structures.ListNode
    rear := head
    for rear != nil {
        tmp := rear
        rear = rear.Next
        tmp.Next = prev
        prev = tmp
    }
    return prev
}

// 双指针法-递归法
func reverseListRecursive(head *structures.ListNode) *structures.ListNode {
    var prev *structures.ListNode
    return reverse(prev, head)
}

func reverse(prev, rear *structures.ListNode) *structures.ListNode {
    if rear == nil {
        return prev
    }
    tmp := rear
    rear = rear.Next
    tmp.Next = prev
    return reverse(tmp, rear)
}
