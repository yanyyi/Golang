package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k == 1 {
		return head
	}

	dummy := &ListNode{
		Next: head,
	}
	pre := dummy
	cur := head
	//tail和next在for循环里确定
	for cur != nil {
		tail := cur
		count := 0

		//让tail指向待翻转区域的最后一个结点
		for count < k-1 && tail != nil {
			tail = tail.Next
			count++
		}

		if tail == nil {
			break
		}

		//由于(待翻转区域)最后一个结点要反向,所以要保存下一个结点
		next := tail.Next

		//在reverseList函数中,待翻转区域最后一个结点的Next指针为空是结束reverseList函数的重要条件
		tail.Next = nil

		//待翻转区域的前驱结点指向翻转后区域的第一个结点
		pre.Next = reverseList(cur)
		//翻转后区域的最后一个结点指向翻转(后)区域的后继结点
		cur.Next = next

		//pre指向翻转(后)区域的最后一个结点,也就是下一个待翻转区域的前驱结点
		pre = cur
		//cur指向待翻转区域的第一个结点
		cur = next
	}

	return dummy.Next
}

// 反转链表: 存 左指 右移 右移
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}
