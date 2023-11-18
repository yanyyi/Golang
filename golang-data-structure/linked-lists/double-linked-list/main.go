package main

import (
	"fmt"
)

// 定义一个HeroNode
type HeroNode struct {
	no       int
	name     string
	nickname string
	pre      *HeroNode //指向前一个节点
	next     *HeroNode //指向下一个结点
}

// 增
func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	//1.先找到该链表的最后这个结点
	//2.创建一个辅助结点
	temp := head
	for {
		if temp.next == nil { //表示找到最后
			break
		}
		temp = temp.next //让temp不断指向下一个结点
	}
	//3.将newHeroNode加入到链表的最后
	temp.next = newHeroNode
	//4.将temp赋值为newHeroNode的前向结点
	newHeroNode.pre = temp
}

// 删 ————双向链表的删除比单向链表容易
// 以根据hero的no删除对应结点为例子
func DeleteHeroNode(head *HeroNode, no int) {
	//先创立一个辅助结点
	temp := head
	if no == 0 {
		fmt.Println("no不能为0")
		return
	}
	//for循环遍历,若找到对应no的英雄结点,进行删除操作;若直到最后还找不到,直接报错返回
	for {
		if temp == nil {
			fmt.Println("无法修改元素")
			return
		}
		if temp.no == no {
			//在链表里删除指定元素
			temp.pre.next = temp.next
			return
		}
		temp = temp.next
	}
}

// 改
// 以修改nickname为例子
func UpdateHeroNode(head *HeroNode, name string, nickname string) {
	//创建一个辅助结点
	temp := head
	//若temp为空结点,直接表示无法修改
	if temp.next == nil {
		fmt.Println("该链表为空链表,无法修改元素")
		return
	}
	//若链表不为空,尝试找到链表里对应的元素
	for {
		temp = temp.next
		if temp == nil {
			fmt.Println("找不到对应元素")
			return
		} else {
			if temp.name == name {
				temp.nickname = nickname
				return
			} else {
				continue
			}
		}
	}
}

// 查 ---正序和逆序
// 正序显示
func ListHeroNode(head *HeroNode) {
	//1.先创建一个辅助结点
	temp := head
	//先判断该链表是不是一个空链表
	if temp.next == nil {
		fmt.Println("空空如也....")
		return
	}
	//2.遍历这个链表
	for {
		fmt.Printf("[%d, %s, %s]==>", temp.next.no, temp.next.name, temp.next.nickname)
		//判断链表是否到最后
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}

// 逆序显示
func ListHeroNode2(head *HeroNode) {
	//1.先创建一个辅助结点
	temp := head
	//先判断该链表是不是一个空链表
	if temp.next == nil {
		fmt.Println("空空如也....")
		return
	}
	// 让temp定位到最后
	for {
		if temp.next != nil {
			temp = temp.next
			continue
		} else {
			break
		}
	}
	//前向遍历,直到遍历到结点的pre的pre(头结点的pre)为空
	for {
		if temp.pre != nil {
			fmt.Printf("[%d, %s, %s]==>", temp.no, temp.name, temp.nickname)
			temp = temp.pre
			continue
		} else {
			break
		}

	}
}

func main() {
	//1.先创建一个头结点
	head := new(HeroNode)
	//2.创建一个新的HeroNode
	hero1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
	}
	hero2 := &HeroNode{
		no:       2,
		name:     "卢俊义",
		nickname: "玉麒麟",
	}
	hero3 := &HeroNode{
		no:       3,
		name:     "林冲",
		nickname: "豹子头",
	}
	hero4 := &HeroNode{
		no:       4,
		name:     "鲁智深",
		nickname: "五大三粗",
	}
	////3.加入
	////在单链表后面直接插入
	InsertHeroNode(head, hero1)
	InsertHeroNode(head, hero2)
	InsertHeroNode(head, hero3)
	InsertHeroNode(head, hero4)
	fmt.Println("插入元素后正向显示的链表:")
	ListHeroNode(head)
	fmt.Println("\n")
	fmt.Println("插入元素后逆向显示的链表:")
	ListHeroNode2(head)
	fmt.Println("\n")
	fmt.Println("修改元素后正向显示的链表:")
	UpdateHeroNode(head, "鲁智深", "花和尚")
	ListHeroNode(head)
	fmt.Println("\n")
	fmt.Println("删除指定元素后的链表:")
	DeleteHeroNode(head, 3)
	ListHeroNode(head)
	fmt.Println("\n")
}
