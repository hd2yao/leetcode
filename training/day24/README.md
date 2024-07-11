## day24

## 代码随想录算法训练营第二十四天| 二叉树 226 101

### 226 翻转二叉树

题目链接：https://leetcode.cn/problems/invert-binary-tree/

文章讲解：https://programmercarl.com/0226.%E7%BF%BB%E8%BD%AC%E4%BA%8C%E5%8F%89%E6%A0%91.html

视频讲解：https://www.bilibili.com/video/BV1sP4y1f7q7

#### 思路
这道题还是比较简单的，就是从上至下，交换左右孩子节点，那么就很容易想到用递归的方法

还有一点需要补充的，递归的前序遍历和后序遍历都没有问题，而中序遍历需要注意的
```go
invert(treeNode.Left)
treeNode.Left, treeNode.Right = treeNode.Right, treeNode.Left
invert(treeNode.Left)
```
1. 先翻转左子树
2. 交换左右子树
3. 再翻转左子树

也就是说，两次翻转子树中间先进行了一次交换，因此原先没有翻转的右子树在步骤 2 后交换到新的左子树上

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day24/0226_invert_binary_tree.go)

### 101 对称二叉树

题目链接：https://leetcode.cn/problems/symmetric-tree/

文章讲解：https://programmercarl.com/0101.%E5%AF%B9%E7%A7%B0%E4%BA%8C%E5%8F%89%E6%A0%91.html

视频讲解：https://www.bilibili.com/video/BV1ue4y1Y7Mf