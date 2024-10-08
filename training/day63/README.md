## day63

## 代码随想录算法训练营第六十三天| 动态规划 518 377

### 518 零钱兑换2

题目链接：https://leetcode.cn/problems/coin-change-ii/

文章讲解：https://programmercarl.com/0518.%E9%9B%B6%E9%92%B1%E5%85%91%E6%8D%A2II.html

视频讲解：https://www.bilibili.com/video/BV1KM411k75j/

#### 思路
从本题开始，就是背包问题中的完全背包了

一句话来说明与 01背包 的不同：完全背包中物品的数量是无限的，即可以选择一个物品重复放入背包

##### 二维 dp

同样的，我还是选择先从二维 dp 开始做起，

`dp[i][j]`：表示 `[0,i]` 前 i-1 个物品中任取将 j 装满一共有 `dp[i][j]` 种方法

这里需要注意的一点是，前面一题我有说到过，我们可以类似想链表中的虚拟头节点将全部的节点统一的方法

本题也同样使用这种方法，因此在上面 dp 的定义中我才说是前 i-1，

然后下面就是核心的计算过程，这个过程不清楚的可以先手动算一遍：

- 不使用当前硬币，`dp[i][j] = dp[i-1][j]`
- 如果还可以使用当前硬币，`dp[i][j] += dp[i][j - coins[i-1]]`

可以看到上面有一点不太一样，如果可以使用当前硬币的时，我们增加的是 `dp[i][j - coins[i-1]]` 而不是 `dp[i - 1][j - coins[i-1]]`

如果一开始不能理解这一点，可以先来使用一下笨方法，列举一下情况，

coins = [1,2] amount = 5，共有以下三种方法，

- 1,1,1,1,1
- 1,1,1,2
- 1,2,2

在手动计算的过程中，我们会发现，有些元素需要多次使用，而如何计算可以使用几次呢，当然就需要根据容量 j 来判断，

```go
for i := 1; i < len(coins)+1; i++ {
    for j := 0; j <= amount; j++ {
        count := j / coins[i-1]
        for k := 0; k <= count; k++ {
            dp[i][j] += dp[i-1][j-k*coins[i-1]]
        }
    }
}
```

上面的代码，我们每次都会先计算一下当前容量最多可以存放几个当前的物品，

例如，j = 5 时，我们计算 coins = 1 的情况，count 此时是等于 5 的，那么应该是以下的情况，

- j = 0 时，再放入 5 个硬币 1
- j = 1 时，再放入 4 个硬币 1
- j = 2 时，再放入 3 个硬币 1
- j = 3 时，再放入 2 个硬币 1
- j = 4 时，再放入 1 个硬币 1
- j = 5 时，再放入 0 个硬币 1

也就是上面代码中最内层循环所计算的过程

同理，如果使用了硬币 1 和 2，计算过程也是这样，

- j = 1 时，再放入 2 个硬币 2
- j = 3 时，再放入 1 个硬币 2

现在我们知道了，如何计算的，那么上面的过程是不是可以再优化一下？

我们以硬币 1 的情况来说明，我先把计算过程写一下，

```
dp[1][5] = dp[0][5] + dp[0][4] + dp[0][3] + dp[0][2] + dp[0][1] + dp[0][0]

dp[1][4] =            dp[0][4] + dp[0][3] + dp[0][2] + dp[0][1] + dp[0][0]

dp[1][3] =                       dp[0][3] + dp[0][2] + dp[0][1] + dp[0][0]

dp[1][2] =                                  dp[0][2] + dp[0][1] + dp[0][0]

dp[1][1] =                                             dp[0][1] + dp[0][0]

dp[1][0] =                                                        dp[0][0]
```

现在是不是有点儿恍然大悟，我再来帮你简化一下

```
dp[1][5] = dp[0][5] + dp[1][4]

dp[1][4] = dp[0][4] + dp[3][3]

dp[1][3] = dp[0][3] + dp[1][2]

dp[1][2] = dp[0][2] + dp[1][1]

dp[1][1] = dp[0][1] + dp[1][0]

dp[1][0] = dp[0][0]
```

为什么是这个结果，其实也很简单，5 个 1 的情况是在 4 个 1 的基础上再加上了一个 1

也就是说，后面的计算结果是要基于前面的结果（这里标记一下，下面的一维会用到！）

我们在举例说明一下上面的结论，就用 coins = [1,2] amount = 5 来说，

`dp[2][5] = dp[0][5] + dp[2][3]` ：

- 不使用硬币 2，那就只使用前面的硬币，即 `dp[1][5] (dp[i-1][j])`
- 使用硬币 2，那就是在 3 的基础上再加上一个即可，而 j = 3 的情况在之前已经计算过，使用同一行就是不再重复已经计算过的情况

至此，就可以明白我们一开始写的递推公式了吧！

##### 一维 dp

我们在二维的基础上，直接降维成一维，

二维：

`dp[i][j] = dp[i-1][j] + dp[i][j - coins[i-1]]`

一维：

`dp[j] = dp[j] + dp[j - coins[i]]` 

在之前的 01背包问题中，我们在遍历都是要 j 从大到小去遍历，这是因为 dp 的值是根据上一行来计算的，如果从小到大遍历计算，
那么在我们还没有使用时，就会被新的值覆盖掉，结果就是有些物品会被多次选取

而在完全背包问题中，我们的物品是可以使用多次，同时我们在上面讲二维时的那处标记还记得嘛，同一行后面的值是需要前面的新值的，
因次我们需要从小到大去遍历，这样也就避免了每次从头计算

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day63/0518_coin_change_ii.go)

### 377 组合总和4

题目链接：https://leetcode.cn/problems/combination-sum-iv/

文章讲解：https://programmercarl.com/0377.%E7%BB%84%E5%90%88%E6%80%BB%E5%92%8C%E2%85%A3.html

视频讲解：https://www.bilibili.com/video/BV1V14y1n7B6/

#### 思路

本题与上一题的不同在于，本题是求排列的个数，需要考虑顺序的，而上一题是不需要考虑顺序的

我们先来说一下，先物品再背包为什么结果就是集合，

因为外层 for 是对物品的循环，即只有上一个物品考虑完所有背包的情况下才会放入下一个物品，这样一来，物品2 一定是在 物品1 之后的！

当我们需要对物品排列时，就不能用上面的顺序，而是应该反过来，先背包再物品

我们可以先想一下这样先背包再物品的场景，nums = [1,2,3] target = 4，我们考虑 target = 4 的情况，

- j = 3 时，放入 1
- j = 2 时，放入 2
- j = 1 时，放入 3

同理，j = 3 的情况，

- j = 2 时，放入 1
- j = 1 时，放入 2

发现了嘛，这就是排列的过程！

然后还有一处需要提醒的是，递推公式中 `dp[i][j] = dp[i-1][j] +  dp[len(nums)][j-nums[i-1]]`

`dp[len(nums)]` 的含义就是在不同背包容量下所以物品的方法总数

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day63/0377_combination_sum_iv.go)