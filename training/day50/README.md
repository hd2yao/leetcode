## day50

## 代码随想录算法训练营第五十天| 贪心法 860 406 452

### 860 柠檬水找零

题目链接：https://leetcode.cn/problems/lemonade-change/

文章讲解：https://programmercarl.com/0860.%E6%9F%A0%E6%AA%AC%E6%B0%B4%E6%89%BE%E9%9B%B6.html

视频讲解：https://www.bilibili.com/video/BV12x4y1j7DD

#### 思路

这道题直接模拟过程就好了

不过看了讲解才知道贪心在哪里：

对于 20 的情况，要优先找零 10 美元和 5 美元，而不是先 3 个 5 美元

因为 5 美元是万能的

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day50/0860_lemonade_change.go)

### 406 根据身高重建队列

题目链接：https://leetcode.cn/problems/queue-reconstruction-by-height/

文章讲解：https://programmercarl.com/0406.%E6%A0%B9%E6%8D%AE%E8%BA%AB%E9%AB%98%E9%87%8D%E5%BB%BA%E9%98%9F%E5%88%97.html

视频讲解：https://www.bilibili.com/video/BV1EA411675Y

#### 思路
这道题还真的没有想法

不过第一步也算是误打误撞的写出来了，在第一次排序之后，就不知道该如何继续推进了

这里有两种方法，
```go
for i := 0; i < len(people); i++ {
    index := people[i][1]
    queue = append(queue[:index], append([][]int{people[i]}, queue[index:]...)...)
}
```

以及，

```go
for i, person := range people {
    copy(people[person[1]+1:i+1], people[person[1]:i+1])
    people[person[1]] = person
}
```

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day50/0406_queue_reconstruction_by_height.go)
