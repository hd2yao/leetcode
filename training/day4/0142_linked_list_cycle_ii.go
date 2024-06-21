package day4

// 快慢指针+虚拟头节点
func detectCycleDummy(head *ListNode) *ListNode {
    dummyHead := &ListNode{Next: head}
    slow, fast := dummyHead, dummyHead

    for fast.Next != nil && fast.Next.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
        if slow == fast {
            slow = dummyHead
            for slow != fast {
                slow = slow.Next
                fast = fast.Next
            }
            return fast
        }
    }

    return nil
}

// 快慢指针
func detectCycle(head *ListNode) *ListNode {
    slow, fast := head, head

    for fast.Next != nil && fast.Next.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
        if slow == fast {
            slow = head
            for slow != fast {
                slow = slow.Next
                fast = fast.Next
            }
            return fast
        }
    }

    return nil
}
