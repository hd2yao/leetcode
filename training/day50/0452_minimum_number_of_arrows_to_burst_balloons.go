package day50

import "sort"

func findMinArrowShots(points [][]int) int {
    res := 1

    // 按照气球的 x_start 来排序
    sort.Slice(points, func(i, j int) bool {
        return points[i][0] < points[j][0]
    })

    for i := 1; i < len(points); i++ {
        // 如果前一位的右边界小于后一位的左边界，则一定不重合
        if points[i-1][1] < points[i][0] {
            res++
        } else {
            // 气球重叠后，更新气球的最小右边界，留到下一个气球进行判断
            points[i][1] = min(points[i-1][1], points[i][1])
        }
    }

    return res
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
