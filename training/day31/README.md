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

### 617 合并二叉树

题目链接：https://leetcode.cn/problems/merge-two-binary-trees/

文章讲解：https://programmercarl.com/0617.%E5%90%88%E5%B9%B6%E4%BA%8C%E5%8F%89%E6%A0%91.html

视频讲解：https://www.bilibili.com/video/BV1m14y1Y7JK

#### 思路
这道题我写的也是半推半就，最后没想到也成功了，因为总是觉得哪里少了，还是对递归不太熟悉

不过通过之后，越看代码越觉得没问题，哈哈哈

二叉树中写递归纠结的点是：总是会觉得 return 就会结束，而忘记了每一次递归都是嵌套在上一次递归中，return 只是结束这一层递归，回溯到上一层，而并非整体结束

##### 递归过程

###### 1. 确定递归参数和返回值
```go
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
    // 处理逻辑
}
```
传入两棵树在相同位置的节点 root1、root2，返回合并后的新节点

###### 2. 确定递归结束条件
这里需要分情况讨论：

- 两个节点都为空，直接返回 nil
- 有一个节点不为空，直接返回不为空的节点，这样此节点下的子树也直接跟随节点返回
```go
if root1 == nil && root2 == nil {
    return root1
}
if root1 == nil && root2 != nil {
    return root2
}
if root1 != nil && root2 == nil {
    return root1
}
```
上面的代码可简化为
```go
if root1 == nil {
    return root2
}
if root2 == nil {
    return root1
}
```

根据上面的两种情况可写出代码

###### 3. 确定单层递归逻辑
单层递归的逻辑就比较简单，两个节点都不为空，那就直接合并

使用前序遍历，先确定中间节点，然后处理左右子节点：

- 选择一个节点，将两个节点的值相加
- 在依次处理左右子节点

最后返回根节点

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day31/0617_merge_two_binary_trees.go)
