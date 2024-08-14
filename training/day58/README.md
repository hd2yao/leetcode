## day58

## 代码随想录算法训练营第五十八天| 动态规划 背包01 416

### 01 背包理论基础

文章讲解：https://programmercarl.com/%E8%83%8C%E5%8C%85%E7%90%86%E8%AE%BA%E5%9F%BA%E7%A1%8001%E8%83%8C%E5%8C%85-1.html

视频讲解：https://www.bilibili.com/video/BV1cg411g7Y6/

#### 二维 dp 数组

##### 1. 确定 dp 数组以及下标的含义

`dp[i][j]`：从下标为 [0 - i] 的物品里任意取，放进容量为 j 的背包中，价值总和的最大值为 `dp[i][j]` 

##### 2. 确定递推公式

具体的过程可以看上面的文章讲解，这里不再赘述，

- **不放物品 i**：由 `dp[i - 1][j]` 推出，即背包容量为 j，里面不放物品 i 的最大价值，此时 `dp[i][j]` 就是 `dp[i - 1][j]`

- **放物品 i**：由 `dp[i - 1][j - weight[i]]` 推出，`dp[i - 1][j - weight[i]]` 为背包容量为 `j - weight[i]` 的时候不放物品 i 的最大价值，
那么 `dp[i - 1][j - weight[i]] + value[i]`，就是背包放物品 i 得到的最大价值

递归公式： `dp[i][j] = max(dp[i - 1][j], dp[i - 1][j - weight[i]] + value[i])`

##### 3. dp 数组初始化

- 背包容量为 0 时，总和一定都是 0
- 只放物品 0 的时，背包容量超过物品 0 重量后，总和都是物品 0 的价值

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day58/01_bag.go)
