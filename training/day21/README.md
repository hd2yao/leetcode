## day21

## 代码随想录算法训练营第二十一天| 栈与队列 347

### 347 前K个高频元素

题目链接：https://leetcode.cn/problems/top-k-frequent-elements/

文章讲解：https://programmercarl.com/0347.%E5%89%8DK%E4%B8%AA%E9%AB%98%E9%A2%91%E5%85%83%E7%B4%A0.html

视频讲解：https://www.bilibili.com/video/BV1Xg41167Lz

#### 思路
我只有一句话，完全不会

堆（heap）是一种满足特定条件的完全二叉树，分为大顶堆和小顶堆

堆通常用于实现优先队列，大顶堆相当于元素从大到小的顺序出队的优先队列

> 因为都是栈顶元素出堆，对于大顶堆来说，堆顶元素是最大的

Go 中 "container/heap" heap 包提供了对任意类型的堆操作
```go
type Interface interface {
    sort.Interface
    Push(x interface{}) // 向末尾添加元素
    Pop() interface{}   // 从末尾删除元素
}
```
只要实现了 sort 接口中的三个方法以及 Push 和 Pop 方法，就可以构造堆了

`func Init(h Interface)` 一个堆在使用任何堆操作之前应先初始化

```go
func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}
```
`IHeap` 是一个切片类型，而在上面的方法中使用的是指针，即 `h` 是一个指向 IHeap 的指针

想要对切片进行操作，则需要 `*h` 解引用 得到原始的 `[][2]int` 切片

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day21/0347_top_k_frequent_elements.go)