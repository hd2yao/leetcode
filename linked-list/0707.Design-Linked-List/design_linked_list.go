package _707_Design_Linked_List

type MyLinkedList struct {
    Len       int
    DummyHead *Node
}

type Node struct {
    Val  int
    Next *Node
}

func Constructor() MyLinkedList {
    return MyLinkedList{
        Len:       0,
        DummyHead: &Node{Next: nil},
    }
}

func (this *MyLinkedList) Get(index int) int {
    if index >= this.Len {
        return -1
    }
    count := 0
    cur := this.DummyHead
    for count < index {
        cur = cur.Next
        count++
    }
    return cur.Next.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
    if this.DummyHead.Next == nil {
        this.DummyHead.Next = &Node{Val: val}
    } else {
        cur := this.DummyHead.Next
        this.DummyHead.Next = &Node{Val: val, Next: cur}
    }
    this.Len++
}

func (this *MyLinkedList) AddAtTail(val int) {
    cur := this.DummyHead
    for cur.Next != nil {
        cur = cur.Next
    }
    cur.Next = &Node{Val: val, Next: nil}
    this.Len++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
    if this.Get(index) != -1 {
        cur := this.DummyHead
        count := 0
        for count < index {
            cur = cur.Next
            count++
        }
        tmp := cur.Next
        cur.Next = &Node{Val: val, Next: tmp}
        this.Len++
    } else if this.Len == index {
        this.AddAtTail(val)
    }
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
    if this.Get(index) != -1 {
        cur := this.DummyHead
        count := 0
        for count < index {
            cur = cur.Next
            count++
        }
        cur.Next = cur.Next.Next
        this.Len--
    }
}
