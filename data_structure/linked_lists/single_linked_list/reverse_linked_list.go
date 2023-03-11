package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 反转链表: 存 左指 右移 右移
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func showList(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Print(curr.Val, "==>")
		curr = curr.Next
	}
	fmt.Println("\n")
}

func main() {
	l1 := new(ListNode)
	l2 := new(ListNode)
	l3 := new(ListNode)
	l4 := new(ListNode)
	l1.Val = 1
	l2.Val = 2
	l3.Val = 3
	l4.Val = 4
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = nil
	showList(l1)
	reverse := reverseList(l1)
	showList(reverse)

}
