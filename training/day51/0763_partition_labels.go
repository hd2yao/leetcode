package day51

import (
    "sort"
    "strings"
)

func partitionLabels(s string) []int {
    alpha := [26]int{}
    intervals := make([][]int, 0)

    // 统计每个字母的区间
    for i := 0; i < len(s); i++ {
        if alpha[s[i]-'a'] == 1 {
            continue
        }
        lastIndex := strings.LastIndex(s, string(s[i]))
        intervals = append(intervals, []int{i, lastIndex})
        alpha[s[i]-'a'] = 1
    }

    // 对区间进行排序，左对齐
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    result := make([]int, 0)
    // 分别记录上一个区间的起始位置和终止位置
    preHead := 0
    preFairy := intervals[0][1]

    for i := 1; i < len(intervals); i++ {
        // 如果区间有重叠，选取最大的右边界
        if preFairy > intervals[i][0] {
            preFairy = max(preFairy, intervals[i][1])
        } else {
            // 说明开始了新的非重叠区间
            result = append(result, intervals[i][0]-preHead)
            preHead = intervals[i][0]
            preFairy = intervals[i][1]
        }
    }

    // 将最后一个区间长度存入
    result = append(result, preFairy-preHead+1)
    return result
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}
