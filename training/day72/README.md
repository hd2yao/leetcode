## day72

## 代码随想录算法训练营第七十二天| 动态规划 300 674 718

### 300 最长递增子序列

题目链接：https://leetcode.cn/problems/longest-increasing-subsequence/

文章讲解：https://programmercarl.com/0300.%E6%9C%80%E9%95%BF%E4%B8%8A%E5%8D%87%E5%AD%90%E5%BA%8F%E5%88%97.html

视频讲解：https://www.bilibili.com/video/BV1ng411J7xP

#### 思路

本题对 `dp[i]` 的定义很巧妙

`dp[i]`：以 `nums[i]` 为结尾的最长子序列的长度为 `dp[i]`

然后依次比较 `nums[j]` 和 `nums[i]` 的大小，j 表示从 0 到 i-1，

如果 `nums[j]` 小于 `nums[i]`，那么 `dp[i]` 至少为 `dp[j] + 1`

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day72/0300_longest_increasing_subsequence.go)
