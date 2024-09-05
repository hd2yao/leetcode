## day80 一刷完成

## 代码随想录算法训练营第八十天| 单调栈 42 84

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

### 84 柱状图中最大的矩形

题目链接：https://leetcode.cn/problems/largest-rectangle-in-histogram/

文章讲解：https://programmercarl.com/0084.%E6%9F%B1%E7%8A%B6%E5%9B%BE%E4%B8%AD%E6%9C%80%E5%A4%A7%E7%9A%84%E7%9F%A9%E5%BD%A2.html

视频讲解：https://www.bilibili.com/video/BV1Ns4y1o7uB/

#### 思路

本题跟上题有些类似，不过要求的刚好相反

下面说一下，这段代码的作用，

```go
curHeight := 0
if i < len(heights) {
    curHeight = heights[i]
}
```

这段代码的作用是处理遍历完 heights 数组后，模拟最后一个柱子之后的情况，以确保栈中剩余的柱子都能被处理

- 当我们遍历到数组的末尾时，栈中可能还存有一些柱子的索引，这些柱子的右边界并没有被触发（即没有碰到比它们更矮的柱子），
因此这些柱子对应的最大矩形面积还没有计算

- 为了处理栈中剩余的元素，我们在遍历完 heights 数组后，手动添加一个高度为 0 的柱子，这样就能模拟出一个比所有柱子都矮的高度，
触发栈中所有元素的弹出操作，确保所有柱子形成的最大矩形都被处理

比如 `heights = [2, 4]`，如果没有上面模拟的过程，那么也就不会触发计算面积的操作

具体情况如下，

- 当 `i < len(heights)` 时，正常获取柱子高度，进行栈的操作
- 当 `i == len(heights)` 时，`curHeight` 被设置为 0，相当于在数组结尾处再加一个高度为 0 的虚拟柱子，触发栈顶元素弹出

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day80/0084_largest_rectangle_in_histogram.go)
