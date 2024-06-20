package day3

// 双指针法-迭代法
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
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
func reverseListRecursive(head *ListNode) *ListNode {
    var prev *ListNode
    return reverse(prev, head)
}

func reverse(prev, rear *ListNode) *ListNode {
    if rear == nil {
        return prev
    }
    tmp := rear
    rear = rear.Next
    tmp.Next = prev
    return reverse(tmp, rear)
}
