## day1

## 代码随想录算法训练营第一天| 704. 二分查找、27. 移除元素

### 704 二分查找

题目链接：https://leetcode.cn/problems/binary-search/

文章讲解：https://programmercarl.com/0704.%E4%BA%8C%E5%88%86%E6%9F%A5%E6%89%BE.html

视频讲解：https://www.bilibili.com/video/BV1fA4y1o715

#### 思路
二分查找的关键：确定好每次查找的左右区间

二分查找的区间一般有两种：**左闭右开** `[left, right)` 和 **左闭右闭** `[left, right]`

##### 左闭右开

当前每次查找的区间为 [left, right)，在这种情况下是**取不到右边界的值**

首先，left == right 是没有意义的，由此可写出循环条件：
```go
for left < right {
    // 具体代码
}
```

接下来判断 中间值 `nums[middle]` 与 目标值 `target` 的大小：

- nums[middle] > target, 说明 target 可能在左区间，当前的 middle 已经不满足要求，可以排除掉，因此 `right = middle`
- nums[middle] < target, 说明 target 可能在右区间，当前的 middle 已经不满足要求，可以排除掉，因此 `left = middle + 1`

##### 左闭右闭

当前每次查找的区间为 [left, right]，在这种情况下左右边界值都是可以取到的

首先，这种情况下的 left == right 是有意义的，由此可写出循环条件：
```go
for left <= right {
    // 具体代码
}
```

接下来判断 中间值 `nums[middle]` 与 目标值 `target` 的大小：

- nums[middle] > target, 说明 target 可能在左区间，当前的 middle 已经不满足要求，可以排除掉，因此 `right = middle - 1`
- nums[middle] < target, 说明 target 可能在右区间，当前的 middle 已经不满足要求，可以排除掉，因此 `left = middle + 1`

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day1/0704_binary_search.go)

### 27 移除元素
    
题目链接：https://leetcode.cn/problems/remove-element/

文章讲解：https://programmercarl.com/0027.%E7%A7%BB%E9%99%A4%E5%85%83%E7%B4%A0.html

视频讲解：https://www.bilibili.com/video/BV12A4y1Z7LP

