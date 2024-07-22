## day32

## 代码随想录算法训练营第三十二天| 二叉树 236

### 236 二叉树的最近公共祖先

题目链接：https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/

文章讲解：https://programmercarl.com/0236.%E4%BA%8C%E5%8F%89%E6%A0%91%E7%9A%84%E6%9C%80%E8%BF%91%E5%85%AC%E5%85%B1%E7%A5%96%E5%85%88.html

视频讲解：https://www.bilibili.com/video/BV1jd4y1B7E2

#### 思路
这道题首先可以想到有两种情况：

- 情况一：p、q 分别在某个节点的两棵子树上，最近公共祖先就是这个节点
- 情况二：p(q) 在 q(p) 的子树上，最近公共祖先就是 q(p)

常规的思维，是从下向上遍历，就可以从 p、q 找到它们的最近公共祖先

因此，我们想到使用后续遍历

##### 1. 确定递归函数的参数以及返回值

比较容易想到，某一个节点 treeNode 的左子树有 p，右子树有 q，那么我们需要一个标记来表示我们找到了 p、q，但是我们需要返回
最近的公共祖先，因此如果找到了 p 或 q 我们就直接返回 p、q，如果没有就返回 nil，用这样的方式来代替上面所说的标记

```go
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {}
```

##### 2. 确定终止条件

按照上面的描述，当我们遇到 p、q 时就返回，并且是递归，在前面几天已经理解过很多次了，递归中的返回是返回给上一层的

其次，如果节点为空，直接返回

```go
if root == p || root == q || root == nil {
    return root
}
```

##### 3. 确定单层递归逻辑

我们使用后续遍历，需要先处理左右子树然后再处理中间节点，由于返回值是节点，我们需要根据左右子树的返回值来判断是否找到了 p、q

- 左右子树都不为空，说明 p、q 分别在该节点的左右两棵子树上，直接返回当前节点
- 左子树为空，右子树为空，说明当前节点只包含了 p(q)，q(p) 应该在其他节点的子树中；这种情况需要返回到上一层，直到回到情况一
- 左右子树都为空，说明这个节点下没有 p、q 直接返回 nil

```go
leftTree := lowestCommonAncestor(root.Left, p, q)
rightTree := lowestCommonAncestor(root.Right, p, q)

if leftTree != nil && rightTree != nil {
    return root
}
if leftTree == nil {
    return rightTree
}
if rightTree == nil {
    return leftTree
}
return nil
```

上面所说的第二种情况，好像是少了我们最开始说的情况二，q 在 p 的子树中，这里解释一下

例如 q 在 p 的子树中，

递归的终止条件是包含 当遇到 q 直接返回 q，那么就没有继续判断 q 的子树中是否还有 p

这样的写法是由题目保证的：题目要求必有 p、q；那么就是说如果遍历整棵树只找到了 q，说明 p 一定就是在 q 的子树中

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day35/0236_lowest_common_ancestor_of_a_binary_tree.go)

### 235 二叉搜索树的最近公共祖先

题目链接：https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-search-tree/

文章讲解：https://programmercarl.com/0235.%E4%BA%8C%E5%8F%89%E6%90%9C%E7%B4%A2%E6%A0%91%E7%9A%84%E6%9C%80%E8%BF%91%E5%85%AC%E5%85%B1%E7%A5%96%E5%85%88.html

视频讲解：https://www.bilibili.com/video/BV1Zt4y1F7ww?share_source=copy_web

#### 思路
只要是二叉搜索树就一定要用到有序这一特性

所以，这道题很明显就是要通过当前节点的值来判断 p、q 的位置

- 节点值在 [p,q] 或 [q,p]，说明当前节点一定是最近的公共祖先
- 节点值大于 p、q，说明两个节点都在左子树
- 节点值小于 p、q，说明两个节点都在右子树

因为我们每次递归时都需要判断，也就是搜索一条边的写法：
```go
if (递归函数(root.Left)) {return }
if (递归函数(root.Right)) {return }
```
而且只需要判断值的大小，也就不涉及前中后序的遍历、

```go
if root == nil {
    return root
}

// 找到节点
if root == p || root == q {
    return root
}

// 区间中
if (root.Val < p.Val && root.Val > q.Val) || (root.Val < q.Val && root.Val > p.Val) {
    return root
}

// p、q 在左子树
if root.Val > q.Val && root.Val > p.Val {
    leftTree := lowestCommonAncestorSearchTree(root.Left, p, q)
    if leftTree != nil {
        return leftTree
    }
}

// p、q 在右子树
if root.Val < q.Val && root.Val < p.Val {
    rightTree := lowestCommonAncestorSearchTree(root.Right, p, q)
    if rightTree != nil {
        return rightTree
    }
}
```
上面是所以情况的代码

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day35/0235_lowest_common_ancestor_of_a_binary_search_tree.go)
