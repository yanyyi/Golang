package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	llen := 0
	curr := head
	for curr != nil {
		llen += 1
		curr = curr.Next
	}
	//删除结点的索引为del
	del := llen - n
	curr = head
	//删除的是第一个结点
	if del == 0 {
		return head.Next
	}
	//删除的不是第一个结点,就令游标停在被删结点的前驱结点
	for i := 0; i < del-1; i++ {
		curr = curr.Next
	}
	curr.Next = curr.Next.Next
	return head

}
