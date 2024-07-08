package day21

import "container/heap"

func topKFrequent(nums []int, k int) []int {
    result := make([]int, 0)
    numsMap := make(map[int]int)
    for _, num := range nums {
        numsMap[num]++
    }
    h := &IHeap{}
    heap.Init(h)
    for num, count := range numsMap {
        heap.Push(h, [2]int{num, count})
        if h.Len() > k {
            heap.Pop(h)
        }
    }
    for i := 0; i < k; i++ {
        result = append(result, heap.Pop(h).([2]int)[0])
    }
    return result
}

// 构建小顶堆

// IHeap 中存放在是 [num, count] 的数组

type IHeap [][2]int

func (h IHeap) Len() int {
    return len(h)
}

func (h IHeap) Less(i, j int) bool {
    return h[i][1] < h[j][1]
}

func (h IHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *IHeap) Push(x interface{}) {
    *h = append(*h, x.([2]int))
}

func (h *IHeap) Pop() interface{} {
    x := (*h)[len(*h)-1]
    *h = (*h)[:len(*h)-1]
    return x
}
