package main

import (
	"fmt"
)

// HeroNode 定义HeroNode
type HeroNode struct {
	no       int
	name     string
	nickName string
	next     *HeroNode
}

// TailInsert 尾插法
func TailInsert(head *HeroNode, t *HeroNode) {
	temp := head
	for {
		if temp.next == nil { //找到最后一个节点
			break
		}
		temp = temp.next
	}
	temp.next = t
}

// InsertById 按Id添加节点
func InsertById(head *HeroNode, t *HeroNode) {
	index := InsertIndex(head, t)
	InsertByIndex(head, t, index)
}

// HeadInsert 头插法
func HeadInsert(head *HeroNode, t *HeroNode) {
	t.next = head.next
	head.next = t
}

func InsertIndex(head *HeroNode, t *HeroNode) int {
	if head.next == nil {
		return 1
	}
	num := 1
	temp := head.next
	for {
		if temp == nil || t.no < temp.no {
			break
		}
		temp = temp.next
		num++
	}
	return num
}

// InsertByIndex 根据位置插入
func InsertByIndex(head *HeroNode, t *HeroNode, n int) {
	if n < 1 {
		fmt.Println("插入错误")
		return
	}
	pre := head
	temp := head.next
	num := 1
	for {
		if num == n || temp == nil {
			break
		}
		pre = temp
		temp = temp.next
		num++
	}
	if num == n {
		pre.next = t
		t.next = temp
	} else {
		fmt.Println("插入位置错误")
	}
}

// Traverse 遍历链表
func Traverse(head *HeroNode) {
	if head.next == nil {
		fmt.Println("空链表")
		return
	}
	temp := head.next
	for temp != nil {
		fmt.Printf("[%d, %s, %s] -> ", temp.no, temp.name, temp.nickName)
		temp = temp.next
	}
	fmt.Println()
}

// 获取链表长度
func size(head *HeroNode) int {
	temp := head.next
	num := 0
	for {
		if temp == nil {
			break
		}
		num++
		temp = temp.next
	}
	return num
}

// 查找节点
func lookupNode(head *HeroNode, n int) *HeroNode {
	temp := head.next
	if temp == nil {
		fmt.Println("空链表")
		return nil
	}
	if n > size(head) {
		fmt.Println("查找范围链表长度")
		return nil
	}
	num := 1
	for {
		if num == n {
			break
		}
		num++
		temp = temp.next
	}
	return temp
}

func DeleteNodeById(head *HeroNode, id int) {
	pre := head
	temp := head.next
	for {
		if temp == nil {
			break
		}
		if temp.no == id {
			pre.next = temp.next
			break
		}
		pre = temp
		temp = temp.next
	}
}

func main() {
	// 初始化头节点
	var head = new(HeroNode)
	hero1 := HeroNode{
		no:       1,
		name:     "宋江",
		nickName: "及时雨",
	}
	hero2 := HeroNode{
		no:       2,
		name:     "卢俊义",
		nickName: "玉麒麟",
	}
	hero3 := HeroNode{
		no:       3,
		name:     "吴用",
		nickName: "智多星",
	}
	InsertById(head, &hero1)
	InsertById(head, &hero3)
	Traverse(head)
	InsertById(head, &hero2)
	Traverse(head)

	hero4 := HeroNode{
		no:       4,
		name:     "公孙胜",
		nickName: "入云龙",
	}
	InsertByIndex(head, &hero4, 4)
	Traverse(head)

	t := lookupNode(head, 3)
	fmt.Println(t.name)

	DeleteNodeById(head, 3)
	Traverse(head)
}
