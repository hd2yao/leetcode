## day78

## 代码随想录算法训练营第七十八天| 动态规划 647 516

### 647 回文子串

题目链接：https://leetcode.cn/problems/palindromic-substrings/

文章讲解：https://programmercarl.com/0647.%E5%9B%9E%E6%96%87%E5%AD%90%E4%B8%B2.html

视频讲解：https://www.bilibili.com/video/BV17G4y1y7z9/

#### 思路

按照以往的动态规划的思路，题目要求什么，我们的 dp 数组就设定为什么

但是本题不太一样，如果直接将 dp[i] 定义为 下标 i 之前的子串的回文个数为 dp[i]

那么在进行推导的时候，我们只能判断出 s[0] 和 s[i] 是否相等，即子串的两头，但无法判断子串中间是否为回文

因此，我们重新定义 dp，

`dp[i][j]`：索引 `i` 到 `j` 的子串 `s[i:j+1]` 是否是回文子串

接下来，我们看一下状态转移：

- 如果 `s[i] == s[j]`，且 `j-i <= 1`（即子串长度为 1 或 2），则 `dp[i][j] = true`
- 如果 `s[i] == s[j]` 且 `dp[i+1][j-1]` 为 true（即去掉首尾字符的子串也是回文），则 `dp[i][j] = true`

因为我们 dp 并不是计数，所以每次 `dp[i][j] = true` 时，我们计数加一

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day78/0647_palindromic_substrings.go)
