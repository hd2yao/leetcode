## day64

## 代码随想录算法训练营第六十四天| 动态规划 70 322 279

### 70 爬楼梯（卡码网）

题目链接：https://kamacoder.com/problempage.php?pid=1067

文章讲解：https://programmercarl.com/0070.%E7%88%AC%E6%A5%BC%E6%A2%AF%E5%AE%8C%E5%85%A8%E8%83%8C%E5%8C%85%E7%89%88%E6%9C%AC.html

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day64/0070_climb_ladder.go)

### 322 零钱兑换

题目链接：https://leetcode.cn/problems/coin-change/

文章讲解：https://programmercarl.com/0322.%E9%9B%B6%E9%92%B1%E5%85%91%E6%8D%A2.html

视频讲解：https://www.bilibili.com/video/BV14K411R7yv/

#### 思路
同样从二维 dp 出发，来推导一维并理解

##### 二维 dp

如何定义 dp？

从题目出发，所求即所得

本题要我们求出最少需要的硬币数，那么 dp 就应该是最少需要的硬币数，因此

`dp[i][j]`：使用前 i 种硬币凑成金额 j 所需的最少硬币数量

对于每一个 `dp[i][j]`，我们有两个选择：

- 不使用第 i 种硬币，那么 `dp[i][j] = dp[i-1][j]`
- 使用第 i 种硬币，如果 j >= coins[i-1]，那么 `dp[i][j] = min(dp[i-1][j], dp[i][j - coins[i-1]] + 1)`

> 这里同样使用的是 `dp[i][j - coins[i-1]] + 1)` 而不是 `dp[i-1][j - coins[i-1]] + 1)`
> 具体可看昨天的讲解

最后在初始化中需要注意的是，如果不使用任何硬币，所有非 0 金额都无法凑成，因此 `dp[0][j] = \infty`（用某个大数表示）

##### 一维 dp

同理可直接推导出一维 dp

`dp[j] = min(dp[j], dp[j-coins[i]]+1)`

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day64/0322_coin_change.go)