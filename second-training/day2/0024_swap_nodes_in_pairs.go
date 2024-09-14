package day2

func swapPairs(head *ListNode) *ListNode {
    dummy := &ListNode{Next: head}
    pre := dummy
    for head != nil && head.Next != nil {
        pre.Next = head.Next
        head.Next = pre.Next.Next
        pre.Next.Next = head
        pre = head
        head = head.Next
    }
    return dummy.Next
}
