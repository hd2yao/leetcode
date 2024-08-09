## day53

## 代码随想录算法训练营第五十三天| 动态规划 509 70 746

### 动态规划五部曲

1. 确定 dp 数组以及下标的含义
2. 确定递推公式
3. dp 数组如何初始化
4. 确定遍历顺序
5. 举例推导 dp 数组

### 509 斐波那契数

题目链接：https://leetcode.cn/problems/fibonacci-number/

文章讲解：https://programmercarl.com/0509.%E6%96%90%E6%B3%A2%E9%82%A3%E5%A5%91%E6%95%B0.html

视频讲解：https://www.bilibili.com/video/BV1f5411K7mo

#### 思路
很简单的一道题，我想所有人学写代码的一个都是 "Hello World!"

那一次了解递归，肯定都是 斐波那契数 了

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day53/0509_fibonacci_number.go)

### 70 爬楼梯

题目链接：https://leetcode.cn/problems/climbing-stairs/

文章讲解：https://programmercarl.com/0070.%E7%88%AC%E6%A5%BC%E6%A2%AF.html

视频讲解：https://www.bilibili.com/video/BV17h411h7UH

#### 思路
真正意义上的第一题，不出意外地出了意外

没有想到 `dp[i]` 的含义，下面就用五部曲来分析一下：

##### 1. 确定 dp 数组以及下标的含义

`dp[i]`：爬到第 i 层楼梯，有 dp[i] 种方法

很巧妙！

我本来还想着第 i 步走到第几层...

##### 2. 确定递推公式

从 `dp[i]` 的定义可以看出，`dp[i]` 可以有两个方向推出来：

- `dp[i - 1]`，上 i-1 层楼梯，有 `dp[i - 1]` 种方法，那么再一步跳一个台阶就是 `dp[i]` 了
- `dp[i - 2]`，上 i-2 层楼梯，有 `dp[i - 2]` 种方法，那么再一步跳两个台阶就是 `dp[i]` 了

所以 `dp[i] = dp[i - 1] + dp[i - 2]` !

##### 3. dp 数组初始化

有了上面的数组含义以及递推公式，这里就比较简单了

剩下的代码也比较容易写出来了

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day53/0070_climbing_stairs.go)

### 746 使用最小花费爬楼梯

题目链接：https://leetcode.cn/problems/min-cost-climbing-stairs/

文章讲解：https://programmercarl.com/0746.%E4%BD%BF%E7%94%A8%E6%9C%80%E5%B0%8F%E8%8A%B1%E8%B4%B9%E7%88%AC%E6%A5%BC%E6%A2%AF.html

视频讲解：https://www.bilibili.com/video/BV16G411c7yZ/

#### 思路
我们还是按照五部曲来，整体跟上面一题很类似的：

##### 1. 确定 dp 数组和下标含义

`dp[i]`：到达第 i 台阶所花费的最少体力为 `dp[i]`

##### 2. 确定递推公式

从 `dp[i]` 的定义可以看出，`dp[i]` 同样有两个方向推出来：

- `dp[i - 1]` 上 1 层楼梯，花费 `dp[i - 1] + cost[i - 1]`
- `dp[i - 2]` 上 2 层楼梯，花费 `dp[i - 2] + cost[i - 2]`

所以 `dp[i] = min( dp[i - 2] + cost[i - 2], dp[i - 1] + cost[i - 1] )` !

##### 3. dp 数组初始化

根据题意，我们可以选择从 0 或 1 级台阶开始，因此

dp[0] = 0, 

dp[1] = 0

##### 4. 确定遍历顺序

这道题有一点不太一样，因为我们要计算的是到顶层所花费的体力，因此我们需要计算到 dp[n]

##### 附：
我开始的时候想的是 

dp[i] 是从哪一层到达第 i 层的

反而更加的复杂，不仅想法复杂而且代码复杂

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day53/0746_min_cost_climbing_stairs.go)
