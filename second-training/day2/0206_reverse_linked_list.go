package day2

func reverseList(head *ListNode) *ListNode {
    var pre *ListNode
    for head != nil {
        cur := head.Next
        head.Next = pre
        pre = head
        head = cur
    }
    return pre
}
