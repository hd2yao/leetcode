package _142_Linked_List_Cycle_II

import "github.com/hd2yao/leetcode/structures"

type ListNode = structures.ListNode

// 快慢指针
func detectCycle(head *ListNode) *ListNode {
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
