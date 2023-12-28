# [242. Valid Anagram](https://leetcode.com/problems/valid-anagram/)

## 题目

给定两个字符串 `s` 和 `t` ，编写一个函数来判断 `t` 是否是 `s` 的字母异位词。

**注意：** 若 `s` 和 `t` 中每个字符出现的次数都相同，则称 `s` 和 `t` 互为字母异位词。

**示例 1:**
> **输入**: s = "anagram", t = "nagaram"
>
> **输出**: true

**示例 2:**
> **输入**: s = "rat", t = "car"
>
> **输出**: false


## 解题思路

这道题可以用打表的方式做。

先把 s 中的每个字母都存在一个 26 个容量的数组里面，每个下标依次对应 26 个字母。s 中每个字母都对应表中一个字母， 每出现一次就加 1。

然后再扫字符串 t，每出现一个字母就在表里面减一。如果都出现了，最终表里面的值肯定都是 0 。最终判断表里面的值是否都是 0 即可，有非 0 的数都输出 false 。
