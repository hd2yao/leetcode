package _142_Linked_List_Cycle_II

import "github.com/hd2yao/leetcode/structures"

type ListNode = structures.ListNode

// 数学等式会推导
func detectCycle(head *ListNode) *ListNode {

    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        // 慢指针每次走一格，快指针每次走两格
        fast = fast.Next.Next
        slow = slow.Next
        // 若存在环，则快慢指针必定在环内相遇
        if fast == slow {
            // 此时，两个指针以相同的速度：一个从头结点出发，一个从环内相遇点出发
            // 相遇点即是入环的第一个结点
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
