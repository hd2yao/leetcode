package _349_Intersection_of_Two_Arrays

import (
    "fmt"
    "testing"
)

type question0349 struct {
    param0349
    ans0349
}

type param0349 struct {
    nums1 []int
    nums2 []int
}

type ans0349 struct {
    result []int
}

func TestIsAnagram(t *testing.T) {

    qs := []question0349{

        {
            param0349{[]int{1, 2, 2, 1}, []int{2, 2}},
            ans0349{[]int{2}},
        },

        {
            param0349{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}},
            ans0349{[]int{9, 4}},
        },
    }

    fmt.Printf("------------------------Leetcode Problem 349------------------------\n")

    for _, q := range qs {
        expected, p := q.ans0349, q.param0349
        fmt.Printf("【input】:%v\n【expected】:%v\n【output】:%v\n", p, expected, intersection(p.nums1, p.nums2))
    }
    fmt.Printf("\n\n\n")
}
