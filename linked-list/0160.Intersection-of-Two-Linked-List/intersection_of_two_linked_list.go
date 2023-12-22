package _160_Intersection_of_Two_Linked_List

import "github.com/hd2yao/leetcode/structures"

type ListNode = structures.ListNode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	la, lb := headA, headB
	for la != lb {
		if la != nil {
			la = la.Next
		} else {
			la = headB
		}

		if lb != nil {
			lb = lb.Next
		} else {
			lb = headA
		}
	}
	return la
}
