package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var result *ListNode
	if l1.Val <= l2.Val {
		result = l1
		result.Next = mergeTwoLists(l1.Next, l2)
	} else {
		result = l2
		result.Next = mergeTwoLists(l1, l2.Next)
	}

	return result
}
