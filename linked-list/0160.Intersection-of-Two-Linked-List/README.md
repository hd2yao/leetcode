# [160. Intersection of Two Linked Lists](https://leetcode.com/problems/intersection-of-two-linked-lists/)

## 题目

给你两个单链表的头节点 `headA` 和 `headB` ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 `null` 。

图示两个链表在节点 `c1` 开始相交：

![](https://assets.leetcode.com/uploads/2018/12/13/160_statement.png)


**示例 1:**

![](https://assets.leetcode.com/uploads/2018/12/13/160_example_1.png)

> **输入**: intersectVal = 8, listA = [4,1,8,4,5], listB = [5,6,1,8,4,5], skipA = 2, skipB = 3
>
> **输出**: Intersected at '8'

**示例 2:**

![](https://assets.leetcode.com/uploads/2018/12/13/160_example_2.png)

> **输入**: intersectVal = 2, listA = [1,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
>
> **输出**: Intersected at '2'

**示例 3:**

![](https://assets.leetcode.com/uploads/2018/12/13/160_example_3.png)

> **输入**: intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
>
> **输出**: null

## 解题思路

这道题的思路其实类似链表找环。

给定的 2 个链表的长度如果一样长，都从头往后扫即可。如果不一样长，需要先“拼成”一样长。把 B 拼接到 A 后面，把 A 拼接到 B 后面。这样 2 个链表的长度都是 A + B。再依次扫描比较 2 个链表的结点是否相同。


