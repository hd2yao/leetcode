package _142_Linked_List_Cycle_II

import "github.com/hd2yao/leetcode/structures"

type ListNode = structures.ListNode

// 快慢指针
// fast 速度为 2，slow 速度为 1，相对速度是 fast 以速度 1 去接近 slow
// 如果 fast 速度为 3，slow 速度为 1，相对速度是 2，fast 就有可能跳过 slow
func detectCycleDummy(head *ListNode) *ListNode {
    dummyHead := &ListNode{Next: head}
    slow, fast := dummyHead, dummyHead
    for fast.Next != nil && fast.Next.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
        if fast == slow {
            fast = dummyHead
            for fast != slow {
                fast = fast.Next
                slow = slow.Next
            }
            return fast
        }
    }
    return nil
}

func detectCycle(head *ListNode) *ListNode {
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
        if fast == slow {
            fast = head
            for fast != slow {
                fast = fast.Next
                slow = slow.Next
            }
            return fast
        }
    }
    return nil
}
