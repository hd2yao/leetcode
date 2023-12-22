package _160_Intersection_of_Two_Linked_List

import (
    "fmt"
    "testing"

    "github.com/hd2yao/leetcode/structures"
)

type question0160 struct {
    param0160
    ans0160
}

type param0160 struct {
    listA []int
    listB []int
}

type ans0160 struct {
    result []int
}

func Test_Problem160(t *testing.T) {

    qs := []question0160{

        {
            param0160{[]int{}, []int{}},
            ans0160{[]int{}},
        },

        {
            param0160{[]int{3}, []int{1, 2, 3}},
            ans0160{[]int{3}},
        },

        {
            param0160{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
            ans0160{[]int{1, 2, 3, 4}},
        },

        {
            param0160{[]int{4, 1, 8, 4, 5}, []int{5, 0, 1, 8, 4, 5}},
            ans0160{[]int{8, 4, 5}},
        },

        {
            param0160{[]int{1}, []int{9, 9, 9, 9, 9}},
            ans0160{[]int{}},
        },

        {
            param0160{[]int{0, 9, 1, 2, 4}, []int{3, 2, 4}},
            ans0160{[]int{2, 4}},
        },
    }

    fmt.Printf("------------------------Leetcode Problem 160------------------------\n")

    for _, q := range qs {
        expected, p := q.ans0160, q.param0160
        fmt.Printf("【input】:%v\n【expected】:%v\n【output】:%v\n", p, expected, structures.List2Ints(getIntersectionNode(structures.Ints2List(p.listA), structures.Ints2List(p.listB))))
    }
    fmt.Printf("\n\n\n")
}
