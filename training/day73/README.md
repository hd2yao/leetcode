## day73

## 代码随想录算法训练营第七十三天| 动态规划 1143 1035 53

### 1143 最长公共子序列

题目链接：https://leetcode.cn/problems/longest-common-subsequence/

文章讲解：https://programmercarl.com/1143.%E6%9C%80%E9%95%BF%E5%85%AC%E5%85%B1%E5%AD%90%E5%BA%8F%E5%88%97.html

视频讲解：https://www.bilibili.com/video/BV1ye4y1L7CQ

#### 思路

`dp[i][j]` ：以下标 `i - 1` 为结尾的 nums1，和以下标 `j - 1` 为结尾的 nums2，最长公共子序列长度为 `dp[i][j]`

跟 **718 最长重复子数组** 思路一致

这里有一点要注意，本题 i 和 j 的定义要理解！

不同的是，本题不要求连续，因此在 text1[i - 1] 与 text2[j - 1] 不相同时，还需要向前去看最大值

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day73/1143_longest_common_subsequence.go)

