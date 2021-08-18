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
func (c *CalStack) priority(operation byte) int {
	switch operation {
	case '*', '/':
		return 1
	case '+', '-':
		return 0
	default:
		return -1
	}
}

func calculator(num1, num2 int, oper rune) int {
	switch oper {
	case '+':
		return num1 + num2
	case '-':
		return num2 - num1
	case '*':
		return num1 * num2
	case '/':
		return num2 / num1
	default:
		panic("符号错误")
	}
}

func isDigit(r byte) bool {
	return unicode.IsDigit(rune(r))
}

func main() {
	expression := "7*2*2-5+1-5+3-4"
	numStack := CalStack{maxSize: 10}
	operStack := CalStack{maxSize: 10}
	var res int
	for index := 0; index < len(expression); index++ {
		ch := expression[index]
		// 判断ch类型
		if isDigit(ch) {
			sum := 0
			for {
				num, _ := strconv.Atoi(string(ch))
				sum = 10*sum + num
				if index == len(expression)-1 || !isDigit(expression[index+1]) {
					break
				}
				ch = expression[index+1]
				index++
			}
			numStack.Push(sum)
		} else {
			for {
				if operStack.isEmpty() {
					operStack.Push(int(ch))
					break
				} else {
					if operStack.priority(ch) <= operStack.priority(byte(operStack.Top())) {
						num1 := numStack.Pop()
						num2 := numStack.Pop()
						oper := operStack.Pop()
						res = calculator(num1, num2, rune(oper))
						numStack.Push(res)
					} else {
						operStack.Push(int(ch))
						break
					}
				}
			}
		}
	}
	for {
		if operStack.isEmpty() {
			break
		}
		num1 := numStack.Pop()
		num2 := numStack.Pop()
		oper := operStack.Pop()
		res = calculator(num1, num2, rune(oper))
		numStack.Push(res)
	}
	fmt.Printf("表达式%s = %d\n", expression, numStack.Top())
}
