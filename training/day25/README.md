## day25

## 代码随想录算法训练营第二十五天| 二叉树 104 111 222

### 104 二叉树的最大深度

题目链接：https://leetcode.cn/problems/maximum-depth-of-binary-tree/

文章讲解：https://programmercarl.com/0104.%E4%BA%8C%E5%8F%89%E6%A0%91%E7%9A%84%E6%9C%80%E5%A4%A7%E6%B7%B1%E5%BA%A6.html

视频讲解：https://www.bilibili.com/video/BV1Gd4y1V75u

#### 思路
这两天的做题发现，层序遍历是真的很好理解也很好写出来，不过前中后序遍历也要掌握，虽然目前只掌握了递归法

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day25/0104_maximum_depth_of_binary_tree.go)

### 111 二叉树的最小深度

题目链接：https://leetcode.cn/problems/minimum-depth-of-binary-tree/

文章讲解：https://programmercarl.com/0111.%E4%BA%8C%E5%8F%89%E6%A0%91%E7%9A%84%E6%9C%80%E5%B0%8F%E6%B7%B1%E5%BA%A6.html

视频讲解：https://www.bilibili.com/video/BV1QD4y1B7e2

#### 思路
最小深度的关键是要理解什么情况下才是最小深度：

最小深度是从根节点到最近的叶子节点的最短路径上的节点数量

上面的关键有两点：

- 从根节点开始
- 到叶子节点为止

因此，终止的条件就是，当前节点为叶子节点，即左右孩子节点都不存在

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day25/0111_minimum_depth_of_binary_tree.go)

### 222 完全二叉树的节点个数

题目链接：https://leetcode.cn/problems/count-complete-tree-nodes/

文章讲解：https://programmercarl.com/0222.%E5%AE%8C%E5%85%A8%E4%BA%8C%E5%8F%89%E6%A0%91%E7%9A%84%E8%8A%82%E7%82%B9%E4%B8%AA%E6%95%B0.html

视频讲解：https://www.bilibili.com/video/BV1eW4y1B7pD

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day25/0222_count_complete_tree_nodes.go)
