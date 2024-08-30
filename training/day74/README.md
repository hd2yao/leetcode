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

### 115 不同的子序列

题目链接：https://leetcode.cn/problems/distinct-subsequences/

文章讲解：https://programmercarl.com/0115.%E4%B8%8D%E5%90%8C%E7%9A%84%E5%AD%90%E5%BA%8F%E5%88%97.html

视频讲解：https://www.bilibili.com/video/BV1fG4y1m75Q/

#### 思路

`dp[i][j]`：表示字符串 t 的前 j-1 个字符在字符串 s 的前 i-1 个字符中作为子序列出现的次数

1. 当 `t[j-1] == s[i-1]` 时，`dp[i][j]` 可以从两个部分得到：

- 不使用 `s[i-1]`：即 `dp[i][j] = dp[i-1][j]`
- 使用 `s[i-1]`：即 `dp[i][j] += dp[i-1][j-1]`

这意味着我们可以选择忽略 `s[i-1]`，也可以将 `s[i-1]` 作为 `t[j-1]`，这两种方案的总和就是 `dp[i][j]`

2. 当 `t[j-1] != s[i-1]` 时，`dp[i][j]` 只能从 `dp[i-1][j]` 得到，因为我们只能忽略 `s[i-1]`

在根据 dp[i][j] 的定义，当 j = 0 时，`dp[i][0] = 1`

因为空字符串是任何字符串的子序列

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day74/0115_distinct_subsequence.go)