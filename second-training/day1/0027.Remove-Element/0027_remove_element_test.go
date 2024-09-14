package _027_Remove_Element

import (
    "fmt"
    "testing"
)

type question0027 struct {
    paras
    ans
}

type paras struct {
    nums   []int
    target int
}

type ans struct {
    result int
}

func TestRemoveElement(t *testing.T) {
    testCases := []question0027{
        {
            paras{[]int{1, 0, 1}, 1},
            ans{1},
        },
        {
            paras{[]int{0, 1, 0, 3, 0, 12}, 0},
            ans{3},
        },

        {
            paras{[]int{0, 1, 0, 3, 0, 0, 0, 0, 1, 12}, 0},
            ans{4},
        },

        {
            paras{[]int{0, 0, 0, 0, 0}, 0},
            ans{0},
        },

        {
            paras{[]int{1}, 1},
            ans{0},
        },

        {
            paras{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2},
            ans{5},
        },
    }

    fmt.Printf("------------------------Leetcode Problem 27------------------------\n")

    for _, testCase := range testCases {
        //fmt.Printf("【input】:%v       【output】:%v\n", testCase.paras, removeElement(testCase.paras.nums, testCase.paras.target))
        // 打印原始的 input
        fmt.Printf("【input】:%v ", testCase.paras)

        // 调用 removeElement 并打印结果
        result := removeElement(testCase.paras.nums, testCase.paras.target)
        fmt.Printf("【output】:%v\n", result)
    }
    fmt.Printf("\n\n\n")
}
