# [1. Two Sum](https://leetcode.com/problems/two-sum/)

## 题目

给定一个整数数组 `nums` 和一个整数目标值 `target`，请你在该数组中找出 **和为目标值** `target`  的那 **两个** 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。


**示例 1:**
> **输入**: nums = [2,7,11,15], target = 9
>
> **输出**: [0,1]
>
> **解释**：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 

**示例 2:**
> **输入**: nums = [3,2,4], target = 6
>
> **输出**: [1,2]

**示例 3:**
> **输入**: nums = [3,3], target = 6
>
> **输出**: [0,1]

## 解题思路

这道题最优的做法时间复杂度是 O(n)。

顺序扫描数组，对每一个元素，在 map 中找能组合给定值的另一半数字，如果找到了，直接返回 2 个数字的下标即可。如果找不到，就把这个数字存入 map 中，等待扫到“另一半”数字的时候，再取出来返回结果。










