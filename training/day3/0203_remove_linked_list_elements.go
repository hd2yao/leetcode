package day3

// 虚拟头节点
func removeElementsDummy(head *ListNode, val int) *ListNode {
    dummyHead := &ListNode{Next: head}
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
func removeElementsTmp(head *ListNode, val int) *ListNode {
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
