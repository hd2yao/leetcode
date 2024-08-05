## day49

## 代码随想录算法训练营第四十九天| 贪心法 1005 134 135

### 1005 K次取反后最大化的数组和

题目链接：https://leetcode.cn/problems/maximize-sum-of-array-after-k-negations/

文章讲解：https://programmercarl.com/1005.K%E6%AC%A1%E5%8F%96%E5%8F%8D%E5%90%8E%E6%9C%80%E5%A4%A7%E5%8C%96%E7%9A%84%E6%95%B0%E7%BB%84%E5%92%8C.html

视频讲解：https://www.bilibili.com/video/BV138411G7LY

#### 思路
这道题挺简单的，要不是因为划分在贪心法这一部分，真的都不知道用到了贪心法

其实还是用到了两种贪心：

- 将绝对值大的负数变为正数
- 如果都是正数，就将值小的正数变为负数

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day49/1005_maximize_sum_of_array_after_k_negations.go)

### 134 加油站

题目链接：https://leetcode.cn/problems/gas-station/

文章讲解：https://programmercarl.com/0134.%E5%8A%A0%E6%B2%B9%E7%AB%99.html

视频讲解：https://www.bilibili.com/video/BV1jA411r7WX

#### 思路
这道题跟前面买卖股票那道题有点儿像，我也想到 `gas[i] - cost[i]` 表示从当前加油站到下一个加油站是否有足够的油

- 大于 0 说明可以从当前出发
- 小于 0 说明油不足够从当前出发到达下一个加油站

然后我根据新生成的数组，一圈一圈的遍历，结果超时，哈哈哈，毕竟还是暴力

贪心法：

局部最优：当前累加 rest[i] 的和 curSum 一旦小于 0，起始位置至少要是 i+1，因为从 i 之前开始一定不行

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day49/0134_gas_station.go)

### 135 分发糖果

题目链接：https://leetcode.cn/problems/candy/

文章讲解：https://programmercarl.com/0135.%E5%88%86%E5%8F%91%E7%B3%96%E6%9E%9C.html

视频讲解：https://www.bilibili.com/video/BV1ev4y1r7wN

#### 思路
开始觉得好简单，手动模拟了一下也没觉得难，开始写代码的时候发现，后面的孩子会影响前面的孩子，一下子就无从下手了

这道题一定是要确定一边之后，再确定另一边：

- 先确定右边评分大于左边（从前向后遍历）
- 再确定左边评分大于右边（从后向前遍历）

这样代码一下子变得超级简单

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day49/0135_candy.go)
