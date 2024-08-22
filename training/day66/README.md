## day66

## 代码随想录算法训练营第六十六天| 动态规划 139

### 139 单词拆分

题目链接：https://leetcode.cn/problems/word-break/

文章讲解：https://programmercarl.com/0139.%E5%8D%95%E8%AF%8D%E6%8B%86%E5%88%86.html

视频讲解：https://www.bilibili.com/video/BV1pd4y147Rh/

#### 思路
同样从二维 dp 出发，来推导一维并理解

##### 二维 dp

`dp[i][j]`：使用前 i 个单词是否可以组成长度为 j 的子串 s[:j+1] 

本题有一个关键点，本质上是排列的完全背包问题，因此我们需要先背包再物品！

`s = "applepenapple", wordDict = ["apple", "pen"] `

"apple", "pen" 是物品，那么我们要求 物品的组合一定是 "apple" + "pen" + "apple" 才能组成 "applepenapple"

"apple" + "apple" + "pen" 或者 "pen" + "apple" + "apple" 是不可以的，那么我们就是强调物品之间顺序！

对于每一个 `dp[i][j]`，我们有两个选择：

- 不使用第 i 个单词，那么 `dp[i][j] = dp[i-1][j]`
- 使用第 i 个单词，那么当前子串又被分为两个部分：除去当前单词长度的前半部分以及单词长度的后半部分，只有两个部分都符合才可以，即
如果 j >= len(wordDict[i-1])，那么 `wordMap[s[j-len(wordDict[i-1]):j]] && dp[len(wordDict)][j-len(wordDict[i-1])]`

##### 一维 dp

其实可以直接考虑一维，反而还会更容易理解一些，

`dp[j]` : 字符串长度为 j 的话，`dp[j]` 为 true，表示可以拆分为一个或多个在字典中出现的单词

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day66/0139_word_break.go)