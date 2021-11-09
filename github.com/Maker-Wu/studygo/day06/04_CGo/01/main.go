package main

//所有的C语言代码需要使用注释写入
/*
#include <stdio.h>
void SayHello()
{
	printf("大家好，我是伍胜强");
}
 */
import "C"

func main() {
	C.SayHello()
}
