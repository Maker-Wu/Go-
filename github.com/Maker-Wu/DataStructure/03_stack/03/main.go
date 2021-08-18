package main

import (
	"fmt"
	"strconv"
	"unicode"
)

type CalStack struct {
	maxSize int
	stack   []int
}

func (c *CalStack) isEmpty() bool {
	return len(c.stack) == 0
}

func (c *CalStack) isFull() bool {
	return len(c.stack) == c.maxSize
}

func (c *CalStack) Push(num int) {
	if c.isFull() {
		fmt.Println("栈满了")
	}
	c.stack = append(c.stack, num)
}

func (c *CalStack) Pop() int {
	if c.isEmpty() {
		panic("空栈")
	}
	val := c.stack[len(c.stack)-1]
	c.stack = c.stack[:len(c.stack)-1]
	return val
}

func (c *CalStack) Top() int {
	return c.stack[len(c.stack)-1]
}

// 返回运算符的优先级
func priority(operation rune) int {
	switch operation {
	case '*', '/':
		return 1
	case '+', '-':
		return 0
	default:
		return -1
	}
}

func infixToSuffix(express string) string {
	s1 := CalStack{maxSize: 10}
	s2 := make([]rune, 0, 10)
	for _, e := range express {
		if unicode.IsDigit(e) {
			s2 = append(s2, e)
		} else if e == '(' {
			s1.Push('(')
		} else if e == ')' {
			for {
				if s1.Top() != '(' {
					s2 = append(s2, rune(s1.Pop()))
				} else {
					s1.Pop()
					break
				}
			}
		} else {
			for {
				if s1.isEmpty() || s1.Top() == '(' {
					s1.Push(int(e))
					break
				} else if priority(e) > priority(rune(s1.Top())) {
					s1.Push(int(e))
					break
				} else {
					s2 = append(s2, rune(s1.Pop()))
				}
			}
		}

	}
	for {
		if s1.isEmpty() {
			break
		}
		s2 = append(s2, rune(s1.Pop()))
	}
	fmt.Println(string(s2))
	return string(s2)
}

func Calculator(s string) int {
	numStack := CalStack{maxSize: 10}
	var res int
	for _, ele := range s {
		if unicode.IsDigit(ele) {
			num, _ := strconv.Atoi(string(ele))
			numStack.Push(num)
		} else {
			num1 := numStack.Pop()
			num2 := numStack.Pop()
			switch ele {
			case '+':
				res = num1 + num2
			case '-':
				res = num2 - num1
			case '*':
				res = num1 * num2
			case '/':
				res = num2 / num1
			default:
				panic("符号错误 ")
			}
			numStack.Push(res)
		}
	}
	return numStack.Top()
}

func main() {
	expression := "1+((2+3)*4)-5"
	expression = infixToSuffix(expression)
	fmt.Println(Calculator(expression))
}
