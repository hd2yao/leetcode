## day72

## 代码随想录算法训练营第七十二天| 动态规划 300 674 718

### 300 最长递增子序列

题目链接：https://leetcode.cn/problems/longest-increasing-subsequence/

文章讲解：https://programmercarl.com/0300.%E6%9C%80%E9%95%BF%E4%B8%8A%E5%8D%87%E5%AD%90%E5%BA%8F%E5%88%97.html

视频讲解：https://www.bilibili.com/video/BV1ng411J7xP

#### 思路

本题对 `dp[i]` 的定义很巧妙

`dp[i]`：以 `nums[i]` 为结尾的最长子序列的长度为 `dp[i]`

然后依次比较 `nums[j]` 和 `nums[i]` 的大小，j 表示从 0 到 i-1，

如果 `nums[j]` 小于 `nums[i]`，那么 `dp[i]` 至少为 `dp[j] + 1`

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day72/0300_longest_increasing_subsequence.go)

### 674 最长连续递增子序列

题目链接：https://leetcode.cn/problems/longest-continuous-increasing-subsequence/

文章讲解：https://programmercarl.com/0674.%E6%9C%80%E9%95%BF%E8%BF%9E%E7%BB%AD%E9%80%92%E5%A2%9E%E5%BA%8F%E5%88%97.html

视频讲解：https://www.bilibili.com/video/BV1bD4y1778v

#### 思路

有了上一题的经验，本题就容易写出来了

因为需要连续，所以只需要与前一个元素比较即可

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day72/0674_longest_continuous_increasing_subsequence.go)

### 718 最长重复子数组

题目链接：https://leetcode.cn/problems/maximum-length-of-repeated-subarray/

文章讲解：https://programmercarl.com/0718.%E6%9C%80%E9%95%BF%E9%87%8D%E5%A4%8D%E5%AD%90%E6%95%B0%E7%BB%84.html

视频讲解：https://www.bilibili.com/video/BV178411H7hV

#### 思路

`dp[i][j]` ：以下标 `i - 1` 为结尾的 nums1，和以下标 `j - 1` 为结尾的 nums2，最长重复子数组长度为 `dp[i][j]`

正如我们开始做动态规划的题目一样，使用 i-1 和 j-1 就是为了在初始化的时候方便

如果直接使用 i 和 j，那么就必须要先初始化第一行和第一列，因为 i = 0 和 j = 0 时都是有意义的

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day72/0718_maximum_length_of_repeated_subarray.go)
