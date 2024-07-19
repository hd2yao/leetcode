## day32

## 代码随想录算法训练营第三十二天| 二叉树 530 501 236

### 530 二叉搜索树的最小绝对差

题目链接：https://leetcode.cn/problems/minimum-absolute-difference-in-bst/

文章讲解：https://programmercarl.com/0530.%E4%BA%8C%E5%8F%89%E6%90%9C%E7%B4%A2%E6%A0%91%E7%9A%84%E6%9C%80%E5%B0%8F%E7%BB%9D%E5%AF%B9%E5%B7%AE.html

视频讲解：https://www.bilibili.com/video/BV1DD4y11779

#### 思路
二叉搜索树的中序遍历可以得到一个递增的有序数组，这一点要记住

因为，求任意两个节点值的最小值，只需要一次比较数组中相邻两个值即可，因为数组是递增的，nums[2] - nums[0] 一定是大于 nums[1] - nums[0]

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day31/0530_minimum_absolute_difference_in_bst.go)
