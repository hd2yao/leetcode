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