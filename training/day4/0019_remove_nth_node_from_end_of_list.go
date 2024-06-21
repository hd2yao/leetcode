package day4

// 快慢指针
func removeNthFromEndOne(head *ListNode, n int) *ListNode {
    dummyHead := &ListNode{Next: head}
    fast := head
    for i := 0; i < n; i++ {
        fast = fast.Next
    }
    for fast != nil && fast.Next != nil {
        fast = fast.Next
        head = head.Next
    }
    if fast == nil {
        dummyHead.Next = head.Next
    } else {
        head.Next = head.Next.Next
    }
    return dummyHead.Next
}

// 快慢指针 + 虚拟头节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummyHead := &ListNode{Next: head}
    fast, head := dummyHead, dummyHead
    for i := 0; i < n; i++ {
        fast = fast.Next
    }
    for fast.Next != nil {
        fast = fast.Next
        head = head.Next
    }

    head.Next = head.Next.Next
    return dummyHead.Next
}
