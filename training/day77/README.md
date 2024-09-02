## day77

## 代码随想录算法训练营第七十七天| 动态规划 583 72

### 583 两个字符串的删除操作

题目链接：https://leetcode.cn/problems/delete-operation-for-two-strings/

文章讲解：https://programmercarl.com/0583.%E4%B8%A4%E4%B8%AA%E5%AD%97%E7%AC%A6%E4%B8%B2%E7%9A%84%E5%88%A0%E9%99%A4%E6%93%8D%E4%BD%9C.html

视频讲解：https://www.bilibili.com/video/BV1we4y157wB/

#### 思路

本题是要删除两个字符串中的字符已达两个字符相等

删除后的字符串其实也就是这两个字符串的最长公共子序列

因此，本题又变成了求最长公共子序列的题目

不过最后的结果需要计算一下，并非 dp 的值

```go
return len(word1) + len(word2) - 2*dp[len(word1)][len(word2)]
```

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day77/0583_delete_operation_for_two_strings.go)
