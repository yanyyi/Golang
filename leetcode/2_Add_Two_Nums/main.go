package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	curr := &ListNode{}
	result := curr
	for l1 != nil || l2 != nil {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		curr.Next = &ListNode{Val: sum % 10}
		carry = sum / 10
		curr = curr.Next
	}
	if carry > 0 {
		curr.Next = &ListNode{Val: 1}
	}

	return result.Next

}
