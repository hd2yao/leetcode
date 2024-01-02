package _454_4Sum_II

import (
    "fmt"
    "testing"
)

type question0454 struct {
    param0454
    ans0454
}

type param0454 struct {
    nums1 []int
    nums2 []int
    nums3 []int
    nums4 []int
}

type ans0454 struct {
    result int
}

func TestIsAnagram(t *testing.T) {

    qs := []question0454{

        {
            param0454{[]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}},
            ans0454{2},
        },

        {
            param0454{[]int{0}, []int{0}, []int{0}, []int{0}},
            ans0454{1},
        },
    }

    fmt.Printf("------------------------Leetcode Problem 454------------------------\n")

    for _, q := range qs {
        expected, p := q.ans0454, q.param0454
        fmt.Printf("【input】:%v\n【expected】:%v\n【output】:%v\n", p, expected, fourSumCount(p.nums1, p.nums2, p.nums3, p.nums4))
    }
    fmt.Printf("\n\n\n")
}
