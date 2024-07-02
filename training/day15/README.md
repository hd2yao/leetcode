## day15

## 代码随想录算法训练营第十五天| 栈与队列 20 1047 150

### 20 有效的括号

题目链接：https://leetcode.cn/problems/valid-parentheses/

文章讲解：https://programmercarl.com/0020.%E6%9C%89%E6%95%88%E7%9A%84%E6%8B%AC%E5%8F%B7.html

视频讲解：https://www.bilibili.com/video/BV1AF411w78g

#### 思路
这道题还是很简单的，只需要把各种不符合的情况全部罗列清楚，然后用代码写出来就可以了

- 左括号可以直接入栈
- 右括号需要依次匹配才可以进行下一步

在上面的过程中，还会出现 '((([{' 全是左括号 以及 ')['、'()][' 打头是右括号的情况，这些情况未必可以在第一次写的时候都能够考虑到

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day15/0020_valid_parentheses.go)

