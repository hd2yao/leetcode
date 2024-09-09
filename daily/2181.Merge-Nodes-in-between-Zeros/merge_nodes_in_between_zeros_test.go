package _181_Merge_Nodes_in_between_Zeros

import (
    "fmt"
    "github.com/hd2yao/leetcode/structures"
    "testing"
)

type question2181 struct {
    param2181
    ans2181
}

type param2181 struct {
    head []int
}

type ans2181 struct {
    result []int
}

func TestMergeNodes(t *testing.T) {
    qs := []question2181{
        {
            param2181{[]int{0, 3, 1, 0, 4, 5, 2, 0}},
            ans2181{[]int{4, 11}},
        },
        {
            param2181{[]int{0, 1, 0, 3, 0, 2, 2, 0}},
            ans2181{[]int{1, 3, 4}},
        },
    }

    fmt.Printf("------------------------Leetcode Problem 2181-------------------------\n")

    for _, q := range qs {
        expected, p := q.ans2181, q.param2181
        fmt.Printf("【input】:%v\n【expected】:%v\n【output】:%v\n", p, expected, structures.List2Ints(mergeNodes(structures.Ints2List(p.head))))
        fmt.Println()
    }
}
