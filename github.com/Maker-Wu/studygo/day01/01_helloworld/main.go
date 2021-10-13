/* 必须在源文件中非注释的第一行指明这个文件属于哪个包
   package main表示这是一个可独立执行的文件
   每个应用程序都应包含一个名为main的包 */
package main

/*
// c语言函数
#include <stdio.h>
void sayHello(){
	printf("Hello world");
}
*/
import "C"
import "fmt"

// 第一个程序
func main() {
	fmt.Println("Hello World")
	C.sayHello()
}
