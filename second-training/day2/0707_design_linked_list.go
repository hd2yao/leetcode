package day2

type MyLinkedList struct {
    Val  int
    Next *MyLinkedList
    Len  int
}

func Constructor() MyLinkedList {
    return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
    if index >= this.Len {
        return -1
    }
    cur := this
    for i := 0; i <= index; i++ {
        cur = cur.Next
    }
    return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
    head := &MyLinkedList{Val: val, Next: this.Next}
    this.Next = head
    this.Len++
}

func (this *MyLinkedList) AddAtTail(val int) {
    cur := this
    for cur.Next != nil {
        cur = cur.Next
    }
    cur.Next = &MyLinkedList{Val: val}
    this.Len++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
    if index == this.Len {
        this.AddAtTail(val)
    } else if index == 0 {
        this.AddAtHead(val)
    } else if index < this.Len {
        cur := this
        for i := 0; i < index; i++ {
            cur = cur.Next
        }
        newNode := &MyLinkedList{Val: val, Next: cur.Next}
        cur.Next = newNode
        this.Len++
    }
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
    if index < this.Len {
        cur := this
        for i := 0; i < index; i++ {
            cur = cur.Next
        }
        cur.Next = cur.Next.Next
        this.Len--
    }
}
