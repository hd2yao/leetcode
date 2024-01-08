# [15. 3Sum](https://leetcode.com/problems/3sum/)

## 题目

给你一个整数数组 `nums` ，判断是否存在三元组 `[nums[i], nums[j], nums[k]]` 满足 `i != j`、`i != k` 且 `j != k` ，同时还满足 `nums[i] + nums[j] + nums[k] == 0` 。

请你返回所有和为 `0` 且 **不重复** 的三元组。

**注意**: 答案中不可以包含重复的三元组。


**示例 1:**
> **输入**: nums = [-1,0,1,2,-1,-4]
>
> **输出**: [[-1,-1,2],[-1,0,1]]
>
> **解释**：
> 
> nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
> 
> 
> nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
> 
> nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
> 
> 不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
> 
> 注意，输出的顺序和三元组的顺序并不重要。

**示例 2:**
> **输入**: nums = [0,1,1]
>
> **输出**: []
> 
> **解释**：唯一可能的三元组和不为 0 。

**示例 3:**
> **输入**: nums = [0,0,0]
>
> **输出**: [[0,0,0]]
>
> **解释**：唯一可能的三元组和为 0 。


## 解题思路

用 map 提前计算好任意 2 个数字之和，保存起来，可以将时间复杂度降到 O(n^2)。这一题比较麻烦的一点在于，最后输出解的时候，要求输出不重复的解。数组中同一个数字可能出现多次，同一个数字也可能使用多次，但是最后输出解的时候，不能重复。例如 [-1，-1，2] 和 [2, -1, -1]、[-1, 2, -1] 这 3 个解是重复的，即使 -1 可能出现 100 次，每次使用的 -1 的数组下标都是不同的。

这里就需要去重和排序了。map 记录每个数字出现的次数，然后对 map 的 key 数组进行排序，最后在这个排序以后的数组里面扫，找到另外 2 个数字能和自己组成 0 的组合。









