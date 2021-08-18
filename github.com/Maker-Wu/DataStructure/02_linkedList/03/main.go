package main

import "fmt"

type CirCleSingleLinkedList struct {
	head *Person //first指向第一个节点
	tail *Person //tail指向最后一个节点
	len  int     //节点数量
}

type Person struct {
	no   int
	next *Person
}

// AddPerson 添加小孩节点
func (c *CirCleSingleLinkedList) init(nums int) {
	if nums < 1 {
		fmt.Println("nums的值不正确")
		return
	}

	p := &Person{no: 1}
	c.head = p
	c.tail = p
	for i := 2; i <= nums; i++ {
		p = &Person{no: i}
		c.tail.next = p
		c.tail = p
	}
	c.tail.next = c.head
	c.len += nums
}

func (c *CirCleSingleLinkedList) TailInsert(t *Person) {
	if c.head == nil {
		c.head = t
	} else {
		c.tail.next = t
	}
	t.next = c.head
	c.tail = t
	c.len++
}

func (c *CirCleSingleLinkedList) Traverse() {
	temp := c.head
	if temp == nil {
		fmt.Println("空链表")
		return
	}
	for {
		fmt.Printf("%d -> ", temp.no)
		temp = temp.next
		if temp == c.head {
			break
		}
	}
	fmt.Println()
}

func (c *CirCleSingleLinkedList) Joseph(startNo, countNum int) {
	if c.head == nil {
		fmt.Println("空链表错误")
		return
	}
	// 小孩报数前，先让head和tail移动k-1
	for i := 0; i < startNo-1; i++ {
		c.head = c.head.next
		c.tail = c.tail.next
	}
	// 小孩报数时，让head和tail指针同时移动m-1次
	for {
		for i := 0; i < countNum-1; i++ {
			c.head = c.head.next
			c.tail = c.tail.next
		}
		fmt.Printf("小孩%d出圈\n", c.head.no)
		c.head = c.head.next
		c.tail.next = c.head
		if c.head == c.tail {
			fmt.Printf("最后留在圈中的小孩编号为%d\n", c.head.no)
			break
		}
	}
}

func main() {
	persons := CirCleSingleLinkedList{}
	persons.init(125)
	persons.Traverse()
	persons.Joseph(10, 20)
}
