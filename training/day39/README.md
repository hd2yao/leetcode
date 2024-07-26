## day39

## 代码随想录算法训练营第三十九天| 回溯法 93 78 90

### 93 复原IP地址

题目链接：https://leetcode.cn/problems/restore-ip-addresses/

文章讲解：https://programmercarl.com/0093.%E5%A4%8D%E5%8E%9FIP%E5%9C%B0%E5%9D%80.html

视频讲解：https://www.bilibili.com/video/BV1XP4y1U73i/

#### 思路

这道题跟昨天分割回文串很像，不过这道题在分割后需要加一个 "."，以及一些约束条件也就是剪枝操作

思路比较简单，但是写代码还是有很多细节需要一点一点去 debug 打印出来修改的，加油！！！

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day39/0093_restore_ip_addresses.go)

### 78 子集

题目链接：https://leetcode.cn/problems/subsets/

文章讲解：https://programmercarl.com/0078.%E5%AD%90%E9%9B%86.html

视频讲解：https://www.bilibili.com/video/BV1U84y1q7Ci

#### 思路

从这道题开始就从组合问题过渡到子集问题了

我们现在已经知道回溯法就是构造了一棵树，那么这两种的区别就在于：

- 组合问题：只记录叶子节点
- 子集问题：记录全部的节点

因此当我们每进行一次递归，也就是每创建一个节点，都要记录

而不是像之前组合问题，只有在递归终止时才记录

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day39/0078_subsets.go)

### 90 子集2

题目链接：https://leetcode.cn/problems/subsets-ii/

文章讲解：https://programmercarl.com/0090.%E5%AD%90%E9%9B%86II.html

视频讲解：https://www.bilibili.com/video/BV1vm4y1F71J/

#### 思路

从这道题跟前面做的 [组合总和2](https://github.com/hd2yao/leetcode/tree/master/training/day38/0040_combination_sum_ii.go) 
以及 [三数之和](https://github.com/hd2yao/leetcode/tree/master/training/day9/0015_3sum.go) 在去重上是一致的

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day39/0090_subsets_ii.go)