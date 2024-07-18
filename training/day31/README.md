## day31

## 代码随想录算法训练营第三十一天| 二叉树 105 106 654 617 700 98

### 106 从中序与后序遍历序列构造二叉树

题目链接：https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/

文章讲解：https://programmercarl.com/0106.%E4%BB%8E%E4%B8%AD%E5%BA%8F%E4%B8%8E%E5%90%8E%E5%BA%8F%E9%81%8D%E5%8E%86%E5%BA%8F%E5%88%97%E6%9E%84%E9%80%A0%E4%BA%8C%E5%8F%89%E6%A0%91.html

视频讲解：https://www.bilibili.com/video/BV1vW4y1i7dn

#### 思路
手动写还是很容易的，但是提炼出代码的递归逻辑就不太行了，跟着讲解走了一遍，发现其实真的很简单

那在自己写一遍

具体的流程备注都在代码中了

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day31/0106_construct_binary_tree_from_inorder_and_postorder_traversal.go)

### 105 从前序与中序遍历序列构造二叉树

题目链接：https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day31/0105_construct_binary_tree_from_preorder_and_inorder_traversal.go)

### 654 最大二叉树

题目链接：https://leetcode.cn/problems/maximum-binary-tree/

文章讲解：https://programmercarl.com/0654.%E6%9C%80%E5%A4%A7%E4%BA%8C%E5%8F%89%E6%A0%91.html

视频讲解：https://www.bilibili.com/video/BV1MG411G7ox

#### 思路
跟上面两道题的思路一致，无非是多了一步先找到切割点，就是找到数组中的最大值的位置

本来想着说 sort 和 copy 会快一些，没想到还是遍历数组更快，应该是数组比较小，然后分割后会更小

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day31/0654_maximum_binary_tree.go)
