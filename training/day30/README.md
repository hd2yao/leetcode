## day30

## 代码随想录算法训练营第三十天| 二叉树 513 112 113 105 106

### 513 找树左下角的值

题目链接：https://leetcode.cn/problems/find-bottom-left-tree-value/

文章讲解：https://programmercarl.com/0513.%E6%89%BE%E6%A0%91%E5%B7%A6%E4%B8%8B%E8%A7%92%E7%9A%84%E5%80%BC.html

视频讲解：https://www.bilibili.com/video/BV1424y1Z7pn

#### 思路
这道题虽然难度为中等，但做起来反而比前两天简单的题还容易

直接使用层序遍历，最后将最后一行的第一个值返回即可

改进：

这里可以不保存全部的值，可以只记录每一层的第一个节点的值，这样可以减少内存使用

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day30/0513_find_bottom_left_tree_value.go)