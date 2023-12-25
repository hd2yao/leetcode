# [142. Linked List Cycle II](https://leetcode.com/problems/linked-list-cycle-ii/)

## 题目

给定一个链表的头节点 `head` ，返回链表开始入环的第一个节点。 如果链表无环，则返回 `null`。

如果链表中有某个节点，可以通过连续跟踪 `next` 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 `pos` 来表示链表尾连接到链表中的位置（**索引从 0 开始**）。
如果 `pos` 是 `-1`，则在该链表中没有环。
**注意：`pos` 不作为参数进行传递**，仅仅是为了标识链表的实际情况。

**不允许修改** 链表。


**示例 1:**

![](https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist.png)

> **输入**: head = [3,2,0,-4], pos = 1
>
> **输出**: 返回索引为 1 的链表节点
>
> **解释解释**：链表中有一个环，其尾部连接到第二个节点。

**示例 2:**

![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/07/circularlinkedlist_test2.png)

> **输入**: head = [1,2], pos = 0
>
> **输出**: 返回索引为 0 的链表节点
>
> **解释**：链表中有一个环，其尾部连接到第一个节点。

**示例 3:**

![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/07/circularlinkedlist_test3.png)

> **输入**: head = [1], pos = -1
>
> **输出**: 返回 null
>
> **解释**：链表中没有环。


## 解题思路

这道题是第 141 题的加强版。在判断是否有环的基础上，还需要输出环的第一个点。

分析一下判断环的原理。fast 指针一次都 2 步，slow 指针一次走 1 步。令链表 head 到环的一个点需要 x1 步，从环的第一个点到相遇点需要 x2 步，从环中相遇点回到环的第一个点需要 x3 步。那么环的总长度是 x2 + x3 步。

fast 和 slow 会相遇，说明他们走的时间是相同的，可以知道他们走的路程有以下的关系：

```c
fast 的 t = (x1 + x2 + x3 + x2) / 2
slow 的 t = (x1 + x2) / 1

x1 + x2 + x3 + x2 = 2 * (x1 + x2)

所以 x1 = x3
```

所以 2 个指针相遇以后，如果 slow 继续往前走，fast 指针回到起点 head，两者都每次走一步，那么必定会在环的起点相遇，相遇以后输出这个点即是结果。



