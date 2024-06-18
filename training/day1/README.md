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

#### 思路
由于数组在内存中的地址是连续的，所以无法直接删除其中某一个元素，只能通过覆盖当前元素实现删除元素的操作

题目要求 **原地** 修改， 因此想到使用 **双指针**，下面是两种使用双指针的方法

##### 双指针法
根据题目要求，返回的结果只看前 k 个不等于 val 的元素，因此想到将等于 val 的元素都移到数组的末尾

- 创建两个指针 front, rear 分别指向数组的头和尾
- 保留尾部等于 val 的元素，以及头部不等于 val 的元素
```go
if nums[front] != val {
    front++
}
if nums[rear] == val {
    rear--
}
```
- 将头部等于 val 的元素与尾部不等于 val 的元素交换
```go
if nums[front] == val && nums[rear] != val {
    nums[front], nums[rear] = nums[rear], nums[front]
    front++
    rear--
}
```
- 由于指针都是向中间移动，因此循环结束的条件就是两个指针相遇
- 此时，指针指向的元素需要进行判断
```go
if nums[front] == val {
    return front
} else {
    return front + 1
}
```

#### 快慢指针法
快慢指针的关键是要理清两个指针的含义
- fast: 获取新元素
- slow: 构建新数组

fast 在原数组中依次寻找，将符合条件的元素传给由 slow 构建的新数组中

```go
if nums[fast] != val {
    nums[slow], nums[fast] = nums[fast], nums[slow]
    fast++
    slow++
} else {
    fast++
}
```

#### 快慢指针法--改进
- 首先，从上面的代码可以看出，fast 指针其实是用来遍历整个数组的，因此可直接 `for-range` 代替 fast指针
- 其次，题目中有说 “nums 的其余元素和 nums 的大小并不重要“ ，因此可以省去上面代码交换值的操作，直接将等于 val 的值覆盖掉
```go
for _, num := range nums {
    if num != val {
        nums[slow] = num
        slow++
    }
}
```

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day1/0027_remove_element.go)