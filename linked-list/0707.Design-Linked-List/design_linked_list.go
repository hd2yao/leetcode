package _707_Design_Linked_List

// 单链表

type MyLinkedList struct {
    Len       int
    DummyHead *SingleListNode
}

type SingleListNode struct {
    Val  int
    Next *SingleListNode
}

func Constructor() MyLinkedList {
    return MyLinkedList{
        Len:       0,
        DummyHead: &SingleListNode{},
    }
}

func (this *MyLinkedList) Get(index int) int {
    if index >= this.Len {
        return -1
    }
    cur := this.DummyHead
    for index > 0 {
        cur = cur.Next
        index--
    }
    return cur.Next.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
    if this.DummyHead.Next == nil {
        this.DummyHead.Next = &SingleListNode{Val: val}
    } else {
        newNode := &SingleListNode{Val: val, Next: this.DummyHead.Next}
        this.DummyHead.Next = newNode
    }
    this.Len++
}

func (this *MyLinkedList) AddAtTail(val int) {
    cur := this.DummyHead
    for cur.Next != nil {
        cur = cur.Next
    }
    cur.Next = &SingleListNode{Val: val, Next: nil}
    this.Len++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
    if this.Get(index) != -1 {
        cur := this.DummyHead
        for index > 0 {
            cur = cur.Next
            index--
        }
        newNode := &SingleListNode{Val: val, Next: cur.Next}
        cur.Next = newNode
        this.Len++
    } else if this.Len == index {
        this.AddAtTail(val)
    }
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
    if this.Get(index) != -1 {
        cur := this.DummyHead
        for index > 0 {
            cur = cur.Next
            index--
        }
        cur.Next = cur.Next.Next
        this.Len--
    }
}
