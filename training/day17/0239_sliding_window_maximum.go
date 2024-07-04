package day17

func maxSlidingWindow(nums []int, k int) []int {
    queue := NewQueue()
    result := make([]int, 0)
    for j := 0; j < k; j++ {
        queue.Push(nums[j])
    }
    result = append(result, queue.Max())
    for i := k; i < len(nums); i++ {
        queue.Pop()
        queue.Push(nums[i])
        result = append(result, queue.Max())
    }
    return result
}

type Queue struct {
    queue []int
}

func NewQueue() *Queue {
    return &Queue{
        queue: make([]int, 0),
    }
}

func (q *Queue) Push(val int) {
    q.queue = append(q.queue, val)
}

func (q *Queue) Pop() {
    if len(q.queue) != 1 {
        q.queue = q.queue[1:]
    } else {
        q.queue = []int{}
    }
}

func (q *Queue) Max() int {
    max := q.queue[0]
    for _, val := range q.queue {
        if val >= max {
            max = val
        }
    }
    return max
}
