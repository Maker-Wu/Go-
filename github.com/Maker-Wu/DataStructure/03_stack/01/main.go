package main

import "fmt"

// Stack 切片实现栈
type Stack struct {
	maxNum int //规定顶栈最多放几个元素
	stack  []int
	top    int //栈顶
}

func New(maxNum int) *Stack {
	return &Stack{
		maxNum: maxNum,
		top:    -1,
	}
}

func (s *Stack) isFull() bool {
	return s.top == s.maxNum-1
}

func (s *Stack) isEmpty() bool {
	return s.top == -1
}

func (s *Stack) Push(val int) {
	if s.isFull() {
		fmt.Println("栈满了")
		return
	}
	s.stack = append(s.stack, val)
	s.top++
}

func (s *Stack) Pop() int {
	if s.isEmpty() {
		panic("栈空了")
	}
	val := s.stack[s.top]
	s.stack = s.stack[:s.top]
	s.top--
	return val
}

func (s *Stack) list() {
	for i := s.top; i >= 0; i-- {
		fmt.Printf("%d ->\n", s.stack[i])
	}
}

func main() {
	l := New(4)
	key := ""
	for {
		fmt.Println("show: 显示栈内容")
		fmt.Println("exit: 退出程序")
		fmt.Println("push: 入栈")
		fmt.Println("pop: 出栈")
		fmt.Scanln(&key)
		switch key {
		case "show":
			l.list()
		case "push":
			fmt.Println("请输入一个数")
			var num int
			fmt.Scanln(&num)
			l.Push(num)
		case "pop":
			res := l.Pop()
			fmt.Println("出栈的数据是", res)
		default:
			l.list()
		}
	}
}
