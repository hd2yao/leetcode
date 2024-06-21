## day4

## 代码随想录算法训练营第四天| 链表 24 19 160 142

### 24 两两交换链表中的节点

题目链接：https://leetcode.cn/problems/swap-nodes-in-pairs/

文章讲解：https://programmercarl.com/0024.%E4%B8%A4%E4%B8%A4%E4%BA%A4%E6%8D%A2%E9%93%BE%E8%A1%A8%E4%B8%AD%E7%9A%84%E8%8A%82%E7%82%B9.html

视频讲解：https://www.bilibili.com/video/BV1YT411g7br

#### 思路
这道题使用双指针没有问题，至于这双指针应该指向哪两个节点，我觉得不需要死记硬背，手动画一画模拟一下交换的过程就明白了

例如：1 -> 2 -> 3 -> 4 -> nil

我开始选则将双指针分别指向要交换的两个节点上，即 `1` 和 `2`，在一次交换过后，链表为 2 -> 1 -> 3 -> 4 -> nil

此时双指针指向了下一组要交换的节点上，即 `3` 和 `4` 上，这时候问题就出现了，没有指向 `1` 的指针，无法进行交换

如果说 `head` 是指向 `1` 的指针，但是当链表更长呢，`head` 指针无法总是指向要进行交换的两个节点的前一个节点，除非说 `head` 指针需要一起跟着移动，那就有三个指针了，
同时，我们还需要一个指针来指向新的头节点

综上所述，上面的过程说明了直接指向交换的两个节点存在两个问题：

- 新的头节点要如何查找
- 无法找到交换节点前的一个节点

针对上面两个问题，比较容易相处解决方案：

- 使用虚拟头节点
- 双指针分别指向交换前一个节点以及要交换的第一个节点

##### 1. 定义虚拟头节点以及初始双指针
```go
dummyHead := &ListNode{Next: head}
front := dummyHead
```
另一个指针可以直接使用 `head`

##### 2. 交换过程
手动画一画，交换的过程还是比较清晰的，此处的代码不唯一：
```go
// 相邻节点进行交换
tmp := head.Next
head.Next = tmp.Next
tmp.Next = head
front.Next = tmp

// 指针向后移动
front = head
head = head.Next
```

##### 3. 循环条件
这里是可以分为两种情况：链表节点个数为奇数和偶数
```go
for head != nil && head.Next != nil {
    // ...
}
```
> - 要先判断 head != nil 再判断 head.Next != nil; 如果反过来 当 head == nil, head.Next 就会发生空指针异常
> - 不能使用 ||, 当链表为奇数个时, head 指向最后一个元素, head.Next == nil 但是 head != nil, 同样再次循环, 即 要同时满足奇数个和偶数个的条件, 因此需要使用 &&

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day4/0024_swap_nodes_in_pairs.go)


### 19 删除链表的倒数第N个节点

题目链接：https://leetcode.cn/problems/remove-nth-node-from-end-of-list/

文章讲解：https://programmercarl.com/0019.%E5%88%A0%E9%99%A4%E9%93%BE%E8%A1%A8%E7%9A%84%E5%80%92%E6%95%B0%E7%AC%ACN%E4%B8%AA%E8%8A%82%E7%82%B9.html

视频讲解：https://www.bilibili.com/video/BV1vW4y1U7Gf

#### 思路
经典的使用快慢指针

其实现在也可以总结出经验了，只要是返回头节点，一定是要使用虚拟头节点，或者说所有的链表题都可以先使用虚拟头节点

这道题，我给出了两段代码，第一段快慢指针都是从 head 出发，这样就会出现当要删除头节点时，快指针会先移动到 nil，因此在循环条件上还有删除操作，
都多出了一种考虑的情况

而从虚拟头节点出发就不会有这种情况

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day4/0019_remove_nth_node_from_end_of_list.go)


### 160 相交链表

题目链接：https://leetcode.cn/problems/intersection-of-two-linked-lists/

文章讲解：https://programmercarl.com/%E9%9D%A2%E8%AF%95%E9%A2%9802.07.%E9%93%BE%E8%A1%A8%E7%9B%B8%E4%BA%A4.html

#### 思路
将两个链表拼接，然后使用两个指针去同步循环，直至指向同一个节点即为相交的节点；若是两个链表不想交，最终两个指针都会同时指向 nil

这里先给出我第一次写的代码，只通过了部分实例
```go
for tmpA != tmpB {
    if tmpA == nil {
        tmpA = headB
    }
    if tmpB == nil {
        tmpB = headA
    }
    tmpA = tmpA.Next
    tmpB = tmpB.Next
}
```
上面的代码跟正确的代码的区别就在于，把写在 else 里的内容都放在了外面，就这会导致在指针指向 nil 后还有以下两个步骤，

- 第一步，将指针重新指向 headA 或是 headB
- 第二部，向后移动一步

而正确的代码，在指针指向 nil 后只会发生上述第一步的过程

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day4/0160_intersection_of_two_linked_list.go)

### 142 环形链表II

题目链接：https://leetcode.cn/problems/linked-list-cycle-ii/

文章讲解：https://programmercarl.com/0142.%E7%8E%AF%E5%BD%A2%E9%93%BE%E8%A1%A8II.html

视频讲解：https://www.bilibili.com/video/BV1if4y1d7ob
