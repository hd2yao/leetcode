package day51

import "sort"

func merge(intervals [][]int) [][]int {
    result := make([][]int, 0)

    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    preHead := intervals[0][0]
    preFairy := intervals[0][1]

    for i := 1; i < len(intervals); i++ {
        if preFairy >= intervals[i][0] {
            preFairy = max(preFairy, intervals[i][1])
        } else {
            result = append(result, []int{preHead, preFairy})
            preHead = intervals[i][0]
            preFairy = intervals[i][1]
        }
    }

    result = append(result, []int{preHead, preFairy})
    return result
}
