package main

import "fmt"

type ListNode []int

func main() {
	l1 := ListNode{1, 2, 3}
	l2 := ListNode{1, 2, 5}
	res := mergeTwoLists(&l1, &l2)
	fmt.Println(res)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// base case
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	// dummyNode method
	dummy := &ListNode{0, nil}
	var dummyP *ListNode = dummy
	var p1 *ListNode = l1
	var p2 *ListNode = l2

	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			dummyP.Next = p1
			p1 = p1.Next
		} else {
			dummyP.Next = p2
			p2 = p2.Next
		}
		dummyP = dummyP.Next
	}

	// ternary operator
	dummyP.Next = map[bool]*ListNode{true: p1, false: p2}[p1 != nil]

	return dummy.Next
}
