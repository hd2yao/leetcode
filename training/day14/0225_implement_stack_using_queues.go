package day14

type MyStack struct {
    Queue1 []int
    Queue2 []int
}

func ConstructorMyStack() MyStack {
    return MyStack{
        Queue1: make([]int, 0),
        Queue2: make([]int, 0),
    }
}

func (this *MyStack) Push(x int) {
    this.Queue1 = append(this.Queue1, x)
}

func (this *MyStack) Pop() int {
    for i := 0; i < len(this.Queue1)-1; i++ {
        this.Queue2 = append(this.Queue2, this.Queue1[i])
    }

    value := this.Queue1[len(this.Queue1)-1]
    this.Queue1 = []int{}

    //for i := 0; i < len(this.Queue2); i++ {
    //	this.Queue1 = append(this.Queue1, this.Queue2[i])
    //}
    //this.Queue2 = []int{}
    // 上面的步骤可用下面的代码代替
    this.Queue1, this.Queue2 = this.Queue2, this.Queue1
    return value
}

func (this *MyStack) Top() int {
    return this.Queue1[len(this.Queue1)-1]
}

func (this *MyStack) Empty() bool {
    if len(this.Queue1) == 0 {
        return true
    }
    return false
}
