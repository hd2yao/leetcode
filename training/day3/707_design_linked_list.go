package day3

type MyLinkedList struct {
    Val    int
    Next   *MyLinkedList
    Length int
}

func Constructor() MyLinkedList {
    return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
    if index > this.Length {
        return -1
    }
    tmp := this
    for i := 0; i <= index; i++ {
        tmp = tmp.Next
    }
    return tmp.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
    node := &MyLinkedList{
        Val:  val,
        Next: this.Next,
    }
    this.Next = node
    this.Length++
}

func (this *MyLinkedList) AddAtTail(val int) {
    node := &MyLinkedList{
        Val: val,
    }
    tmp := this
    for i := 0; i < this.Length; i++ {
        tmp = tmp.Next
    }
    tmp.Next = node
    this.Length++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
    if index == this.Length {
        this.AddAtTail(val)
    } else if index < this.Length {
        node := &MyLinkedList{
            Val: val,
        }
        tmp := this
        for i := 0; i < index; i++ {
            tmp = tmp.Next
        }
        node.Next = tmp.Next
        tmp.Next = node
        this.Length++
    }
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
    if index < this.Length {
        tmp := this
        for i := 0; i < index; i++ {
            tmp = tmp.Next
        }
        tmp.Next = tmp.Next.Next
        this.Length--
    }
}
