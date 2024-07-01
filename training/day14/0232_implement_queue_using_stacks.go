package day14

type MyQueue struct {
    StackIn  []int // 输入栈
    StackOut []int // 输出栈
}

func Constructor() MyQueue {
    return MyQueue{
        StackIn:  make([]int, 0),
        StackOut: make([]int, 0),
    }
}

func (this *MyQueue) Push(x int) {
    this.StackIn = append(this.StackIn, x)
}

func (this *MyQueue) Pop() int {
    if len(this.StackOut) == 0 {
        for i := len(this.StackIn) - 1; i >= 0; i-- {
            this.StackOut = append(this.StackOut, this.StackIn[i])
        }
        this.StackIn = []int{}
    }
    value := this.StackOut[len(this.StackOut)-1]
    this.StackOut = this.StackOut[:len(this.StackOut)-1]
    return value
}

func (this *MyQueue) Peek() int {
    if len(this.StackOut) != 0 {
        return this.StackOut[len(this.StackOut)-1]
    }
    return this.StackIn[0]
}

func (this *MyQueue) Empty() bool {
    if len(this.StackIn) == 0 && len(this.StackOut) == 0 {
        return true
    }
    return false
}
