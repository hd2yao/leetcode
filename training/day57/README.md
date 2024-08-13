## day57

## 代码随想录算法训练营第五十七天| 动态规划 343 96

### 343 整数拆分

题目链接：https://leetcode.cn/problems/integer-break/

文章讲解：https://programmercarl.com/0343.%E6%95%B4%E6%95%B0%E6%8B%86%E5%88%86.html

视频讲解：https://www.bilibili.com/video/BV1Mg411q7YJ/

#### 思路
这道题，我动手写了一遍，大致搞清楚了逻辑

不过还是在看了讲解之后才更加的清晰

##### 1. 确定 dp 数组以及下标的含义

`dp[i]`：拆分数字 i，得到的最大乘积为 `dp[i]` 

##### 2. 确定递推公式

这里我们可以先动手计算一下，

dp[2] = 1 是显而易见的

而 dp[3] 是根据 dp[2] 推导出来的，dp[3] = 2，我们选择 3 = 1 + 2，dp[3] = 1 * 2 = 2

我们可以看到，对于 dp[3] 中计算拆分后的 2，我们并没有选择使用 dp[2] 而是直接使用了 2

因此，对于一个整数 i，dp[i] 是对 i 拆分后的最大乘积，未必会比 i 本身大

那么在后续计算过程中，我们应该选择 `max(i, dp[i])`

所以 `dp[i] = max(dp[i], max((i - j) * j, dp[i - j] * j))` 

`max((i - j) * j, dp[i - j] * j)` 是进行遍历，得出每一种组合的结果，然后再取最大值

##### 3. dp 数组初始化

如题意可得，dp[2] = 1

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day57/0343_integer_break.go)
