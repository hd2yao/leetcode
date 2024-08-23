## day67

## 代码随想录算法训练营第六十七天| 动态规划 121 122

### 121 买卖股票的最佳时机

题目链接：https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/

文章讲解：https://programmercarl.com/0121.%E4%B9%B0%E5%8D%96%E8%82%A1%E7%A5%A8%E7%9A%84%E6%9C%80%E4%BD%B3%E6%97%B6%E6%9C%BA.html

视频讲解：https://www.bilibili.com/video/BV1Xe4y1u77q

#### 思路

首先先定义一下 dp 数组，

- `dp[i][0]` 表示第 i 天**持有**股票所得最多现金
- `dp[i][1]` 表示第 i 天**不持有**股票所得最多现金

那么，

如果第 i 天持有股票即 `dp[i][0]`

- 第 i-1 天就持有股票，那么就保持现状，所得现金就是昨天持有股票的所得现金 即：`dp[i - 1][0]`
- 第 i 天买入股票，所得现金就是买入今天的股票后所得现金即：`-prices[i]`

如果第 i 天不持有股票即 `dp[i][1]`，

- 第 i-1 天就不持有股票，那么就保持现状，所得现金就是昨天不持有股票的所得现金 即：`dp[i - 1][1]`
- 第 i 天卖出股票，所得现金就是按照今天股票价格卖出后所得现金即：`prices[i] + dp[i - 1][0]`

因此，

`dp[i][0] = max(dp[i-1][0], -prices[i])`

`dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])`

以上就是二维 dp

同样的，我们可以发现本题也可以降维到一维，因为 dp[i] 只与 dp[i - 1] 有关，

`dp[0] = max(dp[0], -prices[i])`

`dp[1] = max(dp[1], dp[0]+prices[i])`

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day67/0121_best_time_to_buy_and_sell_stock.go)