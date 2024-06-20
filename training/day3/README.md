## day3

## 代码随想录算法训练营第三天| 链表 203 707 206

### 203 移除链表元素

题目链接：https://leetcode.cn/problems/remove-linked-list-elements/

文章讲解：https://programmercarl.com/0203.%E7%A7%BB%E9%99%A4%E9%93%BE%E8%A1%A8%E5%85%83%E7%B4%A0.html

视频讲解：https://www.bilibili.com/video/BV18B4y1s7R9

#### 思路
这道题并不难，可以使用临时节点方法，也可以使用虚拟头节点的方法：

- 临时节点需要单独考虑头节点的情况
- 增加虚拟头节点，使得原链表中的头节点 head 也成为了与其他节点相当的节点

#### 临时节点
在这种情况下，链表中的节点可分为：**头节点** 和 **剩余节点**

1. 先考虑 剩余节点

即使不考虑 head.Val == val 的情况，也需要考虑 head == nil 的情况，因此循环条件中需要先判断 `tmp != nil`
```go
tmp := head
for tmp != nil && tmp.Next != nil {
    if tmp.Next.Val == val {
        tmp.Next = tmp.Next.Next
    } else {
        tmp = tmp.Next
    }
}
```
此时，链表中除头节点外已经不在有 val 的节点，接下来考虑头节点的情况
```go
if head != nil && head.Val == val {
    return head.Next
}
```

2. 先考虑 头节点

那么首先要将头部节点值等于 val 的节点都删除
```go
for head != nil && head.Val == val {
    head = head.Next
}
```
然后跟上面方法一样，删除剩余节点中的目标节点

#### 虚拟头节点
使用虚拟头节点的好处就是，无需单独考虑头节点的情况，这里不在赘述

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day3/0203_remove_linked_list_elements.go)


### 707 设计链表

题目链接：https://leetcode.cn/problems/design-linked-list/description/

文章讲解：https://programmercarl.com/0707.%E8%AE%BE%E8%AE%A1%E9%93%BE%E8%A1%A8.html

视频讲解：https://www.bilibili.com/video/BV1FU4y1X7WD

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day3/707_design_linked_list.go)
