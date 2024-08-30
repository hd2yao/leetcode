## day74

## 代码随想录算法训练营第七十四天| 动态规划 392 115

### 392 判断子序列

题目链接：https://leetcode.cn/problems/is-subsequence/

文章讲解：https://programmercarl.com/0392.%E5%88%A4%E6%96%AD%E5%AD%90%E5%BA%8F%E5%88%97.html

视频讲解：https://www.bilibili.com/video/BV1tv4y1B7ym/

#### 思路

`dp[i][j]` ：以下标 `i - 1` 为结尾的 s，和以下标 `j - 1` 为结尾的 t，最长公共子序列长度为 `dp[i][j]`

本题与前两天的题目思路是一致的，所以很容易写出来

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day74/0392_is_subsequence.go)

