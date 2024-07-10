## day23

## 代码随想录算法训练营第二十三天| 二叉树 102 107 199 637 429

文章讲解：https://programmercarl.com/0102.%E4%BA%8C%E5%8F%89%E6%A0%91%E7%9A%84%E5%B1%82%E5%BA%8F%E9%81%8D%E5%8E%86.html

视频讲解：https://www.bilibili.com/video/BV1GY4y1u7b2

### 102 二叉树的层序遍历

题目链接：https://leetcode.cn/problems/binary-tree-level-order-traversal/

#### 思路
看完讲解之后发现，层序遍历的逻辑非常容易理解，就着手去写代码，可能是陷入了前面递归的思维方式，一时间用递归也没能写出来，就尝试使用迭代来写

层次遍历的关键就在于两点：

- 弹出一个节点，就将这个节点的左右孩子节点入队
- 如何区分当前队列中的节点是同一层

上面的两点，我们可以通过一层一层来解决，我们在每一层循环中，通过当前层生成出下一层，这样就能确保在下一次循环开始时，队列里的节点都是同一层

##### 1.从 root 开始
```go
queue := []*TreeNode{root}
```
首先第一层就是 root，然后要通过 root 生成下一层，也就是将 root.Left root.Right 放入队列，然后把 root 出队就可以得到只有第二层节点

那如何从包含 N 层和 N + 1 层的队列中，将第 N 层元素出队，因为左右孩子节点不确定是否都存在？

按照上面的描述，每一次循环开始时，队列里只有当前层，此时我们可以先得到当前层的元素，这样在后续中就可以控制出队元素个数
```go
size := len(queue)
for i := 0; i < size; i++ {
    // ...
}
```
##### 2.入队和出队
在上面的循环中，每出队一个元素，就将左右孩子入队，然后记录当前节点的值
```go
currentVal := make([]int, 0)
for i := 0; i < size; i++ {
    currentVal = append(currentVal, queue[0].Val)
    if queue[0].Left != nil {
        queue = append(queue, queue[0].Left)
    }
    if queue[0].Right != nil {
        queue = append(queue, queue[0].Right)
    }
    queue = queue[1:]
}
result = append(result, currentVal)
```

##### 注意点
开始的时候，我是使用一个 level 来记录当前层，
```go
level := 0 
for len(queue) > 0 {
    size := len(queue)
    for i := 0; i < size; i++ {
        result[level] = append(result[level], queue[0].Val)
        if queue[0].Left != nil {
            queue = append(queue, queue[0].Left)
        }
        if queue[0].Right != nil {
            queue = append(queue, queue[0].Right)
        }
        queue = queue[1:]
    }
    level++
}
```
代码的逻辑跟上面是相同的，就是在 `result[level] = append(result[level], queue[0].Val)` 会下标越界，因此这里记录一下用于提醒自己

### 107 二叉树的层序遍历 II

题目链接：https://leetcode.cn/problems/binary-tree-level-order-traversal-ii/

### 199 二叉树的右视图

题目链接：https://leetcode.cn/problems/binary-tree-right-side-view/

### 637 二叉树的层平均值

题目链接：https://leetcode.cn/problems/average-of-levels-in-binary-tree/

### 429 N 叉树的层序遍历

题目链接：https://leetcode.cn/problems/n-ary-tree-level-order-traversal/

---

