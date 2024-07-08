package day17

// 暴力循环解法 -- 超时
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

// 单调队列

func maxSlidingWindowMonotonic(nums []int, k int) []int {
	queue := NewMyQueue()
	result := make([]int, 0)
	for j := 0; j < k; j++ {
		queue.Push(nums[j])
	}
	result = append(result, queue.queue[0])
	for i := k; i < len(nums); i++ {
		queue.Pop(nums[i-k])
		queue.Push(nums[i])
		result = append(result, queue.queue[0])
	}
	return result
}

type MyQueue struct {
	queue []int
}

func NewMyQueue() *MyQueue {
	return &MyQueue{
		queue: make([]int, 0),
	}
}

func (m *MyQueue) Front() int {
	return m.queue[0]
}

func (m *MyQueue) Back() int {
	return m.queue[len(m.queue)-1]
}

func (m *MyQueue) IsEmpty() bool {
	return len(m.queue) == 0
}

func (m *MyQueue) Pop(val int) {
	if !m.IsEmpty() && val == m.Front() {
		m.queue = m.queue[1:]
	}
}

func (m *MyQueue) Push(val int) {
	for !m.IsEmpty() && val > m.Back() {
		m.queue = m.queue[:len(m.queue)-1]
	}
	m.queue = append(m.queue, val)
}
