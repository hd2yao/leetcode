package _206_Reverse_Linked_List

import (
    "fmt"
    "github.com/hd2yao/leetcode/structures"
    "testing"
)

type question0206 struct {
    para0206
    ans0206
}

type para0206 struct {
    nums []int
}

type ans0206 struct {
    result []int
}

func TestReverseList(t *testing.T) {
    qs := []question0206{
        {
            para0206{[]int{1, 2, 3, 4, 5}},
            ans0206{[]int{5, 4, 3, 2, 1}},
        },
        {
            para0206{[]int{}},
            ans0206{[]int{}},
        },
    }

    fmt.Printf("------------------------Leetcode Problem 203------------------------\n")

    for _, q := range qs {
        expected, p := q.ans0206, q.para0206
        fmt.Printf("【input】:%v\n【expected】:%v\n【output】:%v\n", p, expected, structures.List2Ints(reverseListWithDummy(structures.Ints2List(p.nums))))
        fmt.Println()
    }
}
