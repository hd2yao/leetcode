## day42

## 代码随想录算法训练营第四十二天| 回溯法 491 46 47

### 491 非递减子序列

题目链接：https://leetcode.cn/problems/non-decreasing-subsequences/

文章讲解：https://programmercarl.com/0491.%E9%80%92%E5%A2%9E%E5%AD%90%E5%BA%8F%E5%88%97.html

视频讲解：https://www.bilibili.com/video/BV1EG4y1h78v/

#### 思路
这道题跟上周求 子集 是类似的

不同点在于，这道题我们不能先进行排序，这里有一个坑，我先按住不表，先说这道题的整体思路

首先，结果是要将全部的节点记录而非叶子节点，因此终止条件可以不加，因为可选择的数组一直在变小不会无限递归

不过，需要根据题意：至少有两个元素才能记录，记着要在这里增加约束条件

```go
if len(path) > 1 {
    tmp := make([]int, len(path))
    copy(tmp, path)
    result = append(result, tmp)
}
```

然后就是单层的递归逻辑，很容易总结出：

- 每一层 for 循环要进行去重
- 每一层递归要根据题意：非递减 剪枝

我们先来说第一个，对树的每一层节点去重，想前面求子集一样，判断后一个元素是否等于前一个元素
```go
if i != startIndex && nums[i] == nums[i-1] {
	continue
}
```
这里就是上面说的坑！

前面求子集我们可以这么写是因为，我们先对数组排序过，此时整个数组就是非递减的，如果 nums[i] > nums[i-1]，那么 num[i+1] 一定是大于 nums[i-1]

而在这道题中，数组是无序的，我们也不能对数组排序，也就是说，数组会出现 [4,7,6,7] 这种情况，相同的元素并不是相邻出现，那么就不能用上面的代码来去重

我们选择使用 map 来记录每一层已出现过的元素

```go
if used[nums[i]] {
    continue
}
```

第二个剪枝，每次都判断当前元素是否比上一层节点大即可

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day42/0491_non_decreasing_subsequences.go)

### 46 全排列

题目链接：https://leetcode.cn/problems/permutations/

文章讲解：https://programmercarl.com/0046.%E5%85%A8%E6%8E%92%E5%88%97.html

视频讲解：https://www.bilibili.com/video/BV19v4y1S79W/

#### 思路
这道题开始就是全排列问题，也就是需要考虑结果的顺序

和前面组合的不同，就是前面选取过的元素在另一棵子树上还可以再选取一次

也就是说，每一次可选取的元素不是看在数组中的位置，而是看是否已使用过

[1,2,3] 在第一次开始时，是从 1 开始，那么下一步可选择的是 [2,3]

当从 2 开始时，可选取的不只是 [3]，而是 [1,3] 

因此，我们需要记录每次选取的数字，因为数组最大不超过六个数，可以使用 []int 或 []bool，来代替 map

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day42/0046_permutations.go)
