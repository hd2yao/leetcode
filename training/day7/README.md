## day7

## 代码随想录算法训练营第七天| 哈希表 242 349 202 1

### 242 有效的字母异位词

题目链接：https://leetcode.cn/problems/valid-anagram/

文章讲解：https://programmercarl.com/%E5%93%88%E5%B8%8C%E8%A1%A8%E7%90%86%E8%AE%BA%E5%9F%BA%E7%A1%80.html

视频讲解：https://www.bilibili.com/video/BV1YG411p7BA

#### 思路
这道题很容易想到使用 map，所以不再详细说明，主要来说一下针对这道题的一种特殊方法

因为题目中有提到 “`s` 和 `t` 仅包含小写字母”，也就是说这两个字符串种包含的元素最多只有26种，选择使用数组会是更优的方法

如何将字母与数组相结合呢？ 需要把字符映射到数组也就是哈希表的索引下标上

**因为字符`a`到字符`z`的ASCII是26个连续的数值，所以字符a映射为下标0，相应的字符z映射为下标25**
```go
s[i] - 'a'
```
可将字符对应到数组下标上

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day7/0242_vaild_anagram.go)

### 349 两个数组的交集

题目链接：https://leetcode.cn/problems/intersection-of-two-arrays/

文章讲解：https://programmercarl.com/0349.%E4%B8%A4%E4%B8%AA%E6%95%B0%E7%BB%84%E7%9A%84%E4%BA%A4%E9%9B%86.html

视频讲解：https://www.bilibili.com/video/BV1ba411S7wu

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day7/0349_intersection_of_two_arrays.go)

### 202 快乐数

题目链接：https://leetcode.cn/problems/happy-number/

文章讲解：https://programmercarl.com/0202.%E5%BF%AB%E4%B9%90%E6%95%B0.html
