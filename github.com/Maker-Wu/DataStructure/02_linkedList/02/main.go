package main

import "fmt"

// HeroNode 定义HeroNode
type HeroNode struct {
	no       int
	name     string
	nickName string
	next     *HeroNode
	pre      *HeroNode
}

type DoubleLinkedList struct {
	head *HeroNode
	len  int
}

// Traverse 正向遍历数据
func (d *DoubleLinkedList) Traverse() {
	if d.head.next == nil {
		fmt.Println("空链表")
		return
	}
	temp := d.head.next
	for temp != nil {
		fmt.Printf("[%d, %s, %s] -> ", temp.no, temp.name, temp.nickName)
		temp = temp.next
	}
	fmt.Println()
}

// TailInsert 尾插法
func (d *DoubleLinkedList) TailInsert(t *HeroNode) {
	temp := d.head
	for {
		if temp.next == nil { //找到最后一个节点
			break
		}
		temp = temp.next
	}
	temp.next = t
	t.pre = temp
	d.len++
}

// HeadInsert 头插法
func (d *DoubleLinkedList) HeadInsert(t *HeroNode) {
	// 先处理新增节点
	t.pre = d.head
	t.next = d.head.next
	d.head.next = t
	t.next.pre = t
	d.len++
}

// DeleteByIndex 根据索引删除节点
func (d *DoubleLinkedList) DeleteByIndex(n int) {
	temp := d.head.next
	if temp == nil {
		fmt.Println("空链表无法删除")
		return
	}
	if n > d.len {
		fmt.Println("索引超出范围，无法删除")
		return
	}
	num := 1
	for {
		if num == n {
			break
		}
		temp = temp.next
		num++
	}
	temp.pre.next = temp.next
	// temp如果是最后一个节点，会报异常
	if temp.next != nil {
		temp.next.pre = temp.pre
	}
	d.len--
}

func (d *DoubleLinkedList) ModifyNode(n int, name, nickName string) {
	temp := d.head.next
	if temp == nil {
		fmt.Println("空链表无法修改")
		return
	}
	if n > d.len {
		fmt.Println("索引超出范围，无法修改")
		return
	}
	num := 1
	for {
		if num == n {
			break
		}
		temp = temp.next
		num++
	}
	temp.name = name
	temp.nickName = nickName
}

func main() {
	heroLinkedList := DoubleLinkedList{
		head: new(HeroNode),
	}
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
	hero4 := HeroNode{
		no:       4,
		name:     "林冲",
		nickName: "豹子头",
	}
	heroLinkedList.TailInsert(&hero2)
	heroLinkedList.TailInsert(&hero3)
	heroLinkedList.Traverse() //[2, 卢俊义, 玉麒麟] -> [3, 吴用, 智多星] ->
	heroLinkedList.HeadInsert(&hero1)
	heroLinkedList.Traverse()       //[1, 宋江, 及时雨] -> [2, 卢俊义, 玉麒麟] -> [3, 吴用, 智多星] ->
	heroLinkedList.DeleteByIndex(4) //索引超出范围，无法删除
	heroLinkedList.DeleteByIndex(2)
	heroLinkedList.Traverse() //[1, 宋江, 及时雨] -> [3, 吴用, 智多星] ->
	heroLinkedList.TailInsert(&hero4)
	heroLinkedList.ModifyNode(3, "公孙胜", "入云龙")
	heroLinkedList.Traverse() //[1, 宋江, 及时雨] -> [3, 吴用, 智多星] -> [4, 公孙胜, 入云龙] ->
}
