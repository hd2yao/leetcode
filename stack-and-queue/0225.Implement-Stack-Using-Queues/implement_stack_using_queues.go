package _225_Implement_Stack_Using_Queues

type MyStack struct {
    // 创建两个队列
    queue1 []int
    queue2 []int
}

func Constructor() MyStack {
    return MyStack{
        queue1: make([]int, 0),
        queue2: make([]int, 0),
    }
}

func (this *MyStack) Push(x int) {
    // 先将数据存在 queue2 中
    this.queue2 = append(this.queue2, x)
    // 将 queue1 中所有元素移到 queue2 中，再将两个队列进行交换
    this.Move()
}

func (this *MyStack) Move() {
    if len(this.queue1) == 0 {
        // 交换，queue1 置为 queue2，queue2 置为空
        this.queue1, this.queue2 = this.queue2, this.queue1
    } else {
        // queue1 元素从头开始一个一个追加到 queue2 中
        this.queue2 = append(this.queue2, this.queue1[0])
        this.queue1 = this.queue1[1:] // 去除第一个元素
        this.Move()                   // 重复
    }
}

func (this *MyStack) Pop() int {
    val := this.queue1[0]
    this.queue1 = this.queue1[1:] // 去除第一个元素
    return val
}

func (this *MyStack) Top() int {
    return this.queue1[0] // 直接返回
}

func (this *MyStack) Empty() bool {
    return len(this.queue1) == 0
}
