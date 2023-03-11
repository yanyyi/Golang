package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var result *ListNode
	for _, list := range lists {
		result = mergeTwoLists(result, list)
	}
	return result
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var result *ListNode
	if l1.Val < l2.Val {
		result = l1
		result.Next = mergeTwoLists(l1.Next, l2)
	} else {
		result = l2
		result.Next = mergeTwoLists(l2.Next, l1)
	}
	return result

}
