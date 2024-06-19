## day2

## 代码随想录算法训练营第二天| 数组 977 209 59

### 977 有序数组的平方

题目链接：https://leetcode.cn/problems/squares-of-a-sorted-array/

文章讲解：https://programmercarl.com/0977.%E6%9C%89%E5%BA%8F%E6%95%B0%E7%BB%84%E7%9A%84%E5%B9%B3%E6%96%B9.html

视频讲解： https://www.bilibili.com/video/BV1QB4y1D7ep

#### 思路
数组是有序的，且存在负数，此时数组元素平方后的最大值**只会在数组的左右两端**出现，因此考虑使用**双指针法**

知道上面的关键后，这道题的代码并不难写出来，这里想要说明两点注意

##### 需要新建一个数组
开始的时候，我选择在原数组上直接进行修改：
1.  只使用一个指针指 `rear` 向末尾元素 `nums[rear]`
2.  与头元素 `nums[0]` 比较， 若是末尾元素 `nums[rear]` 绝对值大，直接计算平方并重新覆盖到末尾；若是 nums[0] 绝对值大，先交换两个元素，再计算平方覆盖
```go
rear := len(nums) - 1
for rear >= 0 {
    if abs(nums[0]) > abs(nums[rear]) {
        nums[0], nums[rear] = nums[rear], nums[0]    
    }
    nums[rear] = square(nums[rear])
    rear--
}
```
上面的代码在 [-5, -3, -2, -1] 的输入下，输出的结果为 [1, 9, 4, 25]

这是因为，上面的代码只能保证将头尾的较大值放入尾部，但是不能保证每次交换后，头尾元素的绝对值在整个数组中是最大的两个，因为元素交换后，数组可能不再有序

##### 使用平方值比较
代码中可以直接使用 **平方值** 代替 **绝对值** 去比较

虽然 **绝对值** 是第一反应，不过最终结果是计算平方，这一点可以作为考虑

##### 使用 index
若是不使用 index 标记新数组并将平方值直接放入对应位置，只能是使用 append 结果是平方值降序排序，最后还需要对 result 做一个逆序的操作

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day2/0977_squares_of_a_sorted_array.go)

### 209 长度最小的子数组
题目链接：https://leetcode.cn/problems/minimum-size-subarray-sum/

文章讲解：https://programmercarl.com/0209.%E9%95%BF%E5%BA%A6%E6%9C%80%E5%B0%8F%E7%9A%84%E5%AD%90%E6%95%B0%E7%BB%84.html

视频讲解：https://www.bilibili.com/video/BV1tZ4y1q7XE