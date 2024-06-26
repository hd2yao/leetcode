## day9

## 代码随想录算法训练营第九天| 哈希表 454 383 15 18

### 454 四数相加II

题目链接：https://leetcode.cn/problems/4sum-ii/

文章讲解：https://programmercarl.com/0454.%E5%9B%9B%E6%95%B0%E7%9B%B8%E5%8A%A0II.html

视频讲解：https://www.bilibili.com/video/BV1Md4y1Q7Yh

#### 思路
这道题的关键有两点：

- `nums1[i] + nums2[j] == - nums3[k] - nums4[l]`
- 如何使用 map 来保存信息

首先第一点，两两组合之后，四数之和降维成两数之和，将四重循环转变为两个二重循环，大大降低了时间复杂度，避免超时

再来看第二点，在两数之和中，map 中的 key 保存的是 `target - nums[i]` ，同样的四数之和中，key 也保存 `nums1[i] + nums2[j]` 或是 `- nums1[i] - nums2[j]`；
两数之和中只会出现一组结果，因此 value 直接保存对应的下标，而在本题中，组合的结果不止一种，而且只需要返回结果的数量，而不需要返回全部的组合，
因此直接统计数量即可
```go
for i := 0; i < len(nums1); i++ {
    for j := 0; j < len(nums2); j++ {
        targetMap[nums1[i]+nums2[j]]++
    }
}
```
[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day9/0454_4sum_ii.go)


### 383 赎金信

题目链接：https://leetcode.cn/problems/ransom-note/

文章讲解：https://programmercarl.com/0383.%E8%B5%8E%E9%87%91%E4%BF%A1.html

#### 思路
这道题还是很简单的，跟之前 [0242-vaild-anagram](https://github.com/hd2yao/leetcode/tree/master/training/day7/0242_vaild_anagram.go) 差不多

[完整代码](https://github.com/hd2yao/leetcode/tree/master/training/day9/0383_ransom_note.go)