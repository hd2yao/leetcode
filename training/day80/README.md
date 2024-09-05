## day80

## 代码随想录算法训练营第八十天| 单调栈 739 496 503

### 42 接雨水

题目链接：https://leetcode.cn/problems/trapping-rain-water/

文章讲解：https://programmercarl.com/0042.%E6%8E%A5%E9%9B%A8%E6%B0%B4.html

视频讲解：https://www.bilibili.com/video/BV1uD4y1u75P/

#### 思路

本题我一开始只写出来前半部分，这是什么意思呢

就是说，我找到了凹槽，根据昨天题目的思路，要想形成凹槽，那么至少需要后面的高度比前面的高度高

转化一下就成了，去找每个柱子右边的第一个不低于它的柱子，这样一说是不是有点儿明白了

但是我一开始去求雨水的时候，想法是不对的，我只考虑了右边界

因为是第一次做，整体的思路是理解了，但还不能很好的讲解清楚，所以这里不去过多说明

详细的解释，结合代码一起理解会更清楚

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day80/0042_trapping_rain_water.go)
