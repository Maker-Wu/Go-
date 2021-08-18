//使用闭包函数遍历二叉树
package main

import "fmt"

type node struct {
	Value       int
	Left, Right *node
}

func (n *node) Print() {
	fmt.Printf("%d\n", n.Value)
}

func (n *node) Traverse() {
	if n == nil {
		return
	}
	n.Left.Traverse()
	n.Print() //只能打印
	n.Right.Traverse()
}

func (n *node) TraverseFunc(f func(*node)) {
	if n == nil {
		return
	}
	n.Left.TraverseFunc(f)
	f(n)
	n.Right.TraverseFunc(f)
}

func (n *node) Traverse2() {
	n.TraverseFunc(func(n *node) {
		n.Print()
	})
}

func main() {
	root := node{Value: 3}
	root.Left = &node{Value: 0}
	root.Right = &node{Value: 5}
	root.Left.Right = &node{Value: 2}
	root.Right.Left = &node{Value: 4}
	root.Traverse2()

	nodeCount := 0
	root.TraverseFunc(func(n *node) {
		nodeCount++
	})
	fmt.Printf("New count:%d\n", nodeCount)
}
