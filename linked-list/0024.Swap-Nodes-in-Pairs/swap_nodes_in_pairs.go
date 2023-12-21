package _024_Swap_Nodes_in_Pairs

import "github.com/hd2yao/leetcode/structures"

type ListNode = structures.ListNode

func swapPairs(head *ListNode) *ListNode {
    dummyHead := &ListNode{Next: head}
    cur := dummyHead
    // 要先判断 head != nil 再判断 head.Next != nil
    // 如果反过来 当 head == nil, head.Next 就会发生空指针异常
    // cur.Next != nil && cur.Next.Next != nil
    // 不能使用 ||, 当链表为奇数个时, head 指向最后一个元素, head.Next == nil 但是 head != nil, 同样再次循环
    // 即 要同时满足奇数个和偶数个的条件, 因此需要使用 &&
    for head != nil && head.Next != nil {
        cur.Next = head.Next
        head.Next = head.Next.Next
        cur.Next.Next = head

        cur = head
        head = head.Next
    }

    return dummyHead.Next
}

// 递归

func swapPairsRecursive(head *ListNode) *ListNode {
    dummyHead := &ListNode{Next: head}
    cur := dummyHead
    swap(cur, head)
    return dummyHead.Next
}

func swap(cur, head *ListNode) {
    if cur.Next != nil && cur.Next.Next != nil {
        cur.Next = head.Next
        head.Next = head.Next.Next
        cur.Next.Next = head
        swap(head, head.Next)
    }
}
