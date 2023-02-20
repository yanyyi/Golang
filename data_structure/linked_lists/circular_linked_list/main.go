package main

import "fmt"

type CatNode struct {
	no   int
	name string
	next *CatNode
}

// 增
func InsertCatNode(head *CatNode, newCatNode *CatNode) {
	//添加的是第一只猫
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head
		return
	}
	//创建辅助结点
	temp := head
	//添加的不是第一只猫,让temp指向链表最后一个添加的元素
	for {
		if temp.next == head {
			break
		}

		temp = temp.next

	}
	//将新来的元素加入到链表中
	temp.next = newCatNode
	newCatNode.next = head
}

// 删,以指定猫的名字为例
func DeleteCatNode(head *CatNode, name string) {
	fmt.Println("开始删除结点...")
	if head.next == nil {
		fmt.Println("该链表为空,无法删除元素....")
		return
	}
	temp := head
	helper := head //辅助跟随找到要删除结点的前一个节点

	//如果只有一只猫
	if temp.next == head {
		if head.name == name {
			temp.next = nil
			return
		} else {
			return
		}

	}

	//如果至少有两只猫
	//让helper指向队尾,即temp的前一个节点
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}
	for {
		if temp.next == head {
			fmt.Println("找不到要修改的元素...")
			break
		}

		if temp.name == name {
			helper.next = temp.next
			break
		}
		temp = temp.next
		helper = helper.next
	}
	return
}

// 改,以指定no修改猫的名字为例
func UpdateCatNode(head *CatNode, no int, name string) {
	fmt.Println("开始修改结点....")
	if head.next == nil {
		fmt.Println("该链表为空,无法修改元素")
		return
	}
	//将头结点赋值给临时变量temp
	temp := head
	//依次对比每个结点的no,不满足则往下一个结点走,直至返回到head
	for {
		if temp.no == no {
			temp.name = name
			return
		}
		temp = temp.next
		if temp == head {
			return
		}
	}

}

// 查
func ListCatNode(head *CatNode) {
	//设立辅助结点temp
	temp := head
	//判断是否为空链表
	if temp.next == nil {
		fmt.Println("该链表为空.....")
		return
	}
	//打印链表,直到temp.next=head为止
	for {
		fmt.Printf("[编号为%d的猫的名字叫%s]==>", temp.no, temp.name)
		if temp.next == head {
			break
		}
		temp = temp.next

	}
}

func main() {
	head := new(CatNode)
	newCatNode1 := &CatNode{
		no:   1,
		name: "Tom",
	}
	newCatNode2 := &CatNode{
		no:   2,
		name: "Mike",
	}
	//换一种初始化方法
	newCatNode3 := new(CatNode)
	newCatNode3.no = 3
	newCatNode3.name = "Jerry"

	InsertCatNode(head, newCatNode1) //增
	InsertCatNode(head, newCatNode2)
	InsertCatNode(head, newCatNode3)
	UpdateCatNode(head, 1, "Tom2") //改
	ListCatNode(head)              //查
	fmt.Println("\n")
	DeleteCatNode(head, "Mike") //删
	ListCatNode(head)
}
