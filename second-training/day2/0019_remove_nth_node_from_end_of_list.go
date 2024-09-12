package day2

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{Next: head}
    cur := dummy
    for i := 0; i < n; i++ {
        head = head.Next
    }
    for head != nil {
        head = head.Next
        cur = cur.Next
    }
    cur.Next = cur.Next.Next
    return dummy.Next
}
