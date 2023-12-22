package _019_Remove_Nth_Node_from_End_of_List

import (
    "fmt"
    "github.com/hd2yao/leetcode/structures"
    "testing"
)

type question0019 struct {
    param0019
    ans0019
}

type param0019 struct {
    nums   []int
    target int
}

type ans0019 struct {
    result []int
}

func TestRemoveNthFromEnd(t *testing.T) {
    qs := []question0019{
        {
            param0019{[]int{1, 2, 3, 4, 5}, 2},
            ans0019{[]int{1, 2, 3, 5}},
        },
        {
            param0019{[]int{1, 2}, 1},
            ans0019{[]int{1}},
        },
        {
            param0019{[]int{1}, 1},
            ans0019{[]int{}},
        },
    }

    fmt.Printf("------------------------Leetcode Problem 24-------------------------\n")

    for _, q := range qs {
        expected, p := q.ans0019, q.param0019
        fmt.Printf("【input】:%v\n【expected】:%v\n【output】:%v\n", p, expected, structures.List2Ints(removeNthFromEnd(structures.Ints2List(p.nums), p.target)))
        fmt.Println()
    }
}
