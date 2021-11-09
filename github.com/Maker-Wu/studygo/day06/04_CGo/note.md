Go 有强烈的 C 背景，在 Go 与 C 语言互操作方面，Go 更是提供了强大的支持。尤其是在 Go 中使用 C，甚至可以直接在 Go 源文件中编写 C 代码，这是其他语言所无法望其项背的。

CGO 是让 Go 程序在 Android 和 IOS 上运行的关键。

在如下一些场景中，可能会涉及到 Go 与 C 的互操作：

- 提升局部代码性能时，用 C 替换一些 Go 代码。

#### 前置条件

要使用 CGO 特性，需要安装 C/ C++ 构建工具链。

在 macOs 和 Linux 下安装 GCC，使用命令安装：yum install gcc

在 windows 下是需要安装 MinGW 工具 http://www.mingw-w64.org/doku.php。

同时需要保证环境变量 CGO_ENABLED 被设置为 1，这表示 CGO 是被启用的状态。在本地构建时 CGO_ENABLED 默认是启用的，当交叉构建时 CGO 默认是禁止的。

![image-20211103213135211](F:\Users\WSQ\Go\src\Go-\github.com\Maker-Wu\studygo\day06\04_CGo\note.assets\image-20211103213135211.png)

然后通过 import "C" 语句启用CGO特性。

```go
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
```

#### #cgo 语句

可以通过 #cgo 语句设置编译阶段和链接阶段的相关参数。

编译阶段的参数主要用于定义相关宏和指定头文件检索路径。

链接阶段的参数主要是指定库文件检索路径和要链接的库文件。

```go
// #cgo CFLAGS: -D PORT = 8080 -I./include 
// #cgo LDFLAGS: -L/usr/local/lib -llibevent 
// #include <libevent.h>
import "C"
```

CFLAGS 部分，`-D` 部分定义了宏 PORT，值为8080；`-I` 定义了头文件包含的检索目录。
LDFLAGS 部分，`-L`指定了链接时库文件检索目录，`-l` 指定了链接时需要链接 libevent库。

示例：

写好 hello.c 文件

```c
#include <stdio.h>
#include "hello.h"
void print()
{
	printf("hello from c\n");
}
```

写好头文件 hello.h

```c
void print();
```

编译成 .so，起名是 libhello.so

```shell
gcc hello.c -fPIC -shared -o libhello.so
```

编写 go 文件，hg.go

```go
package main
 
/*
#cgo CFLAGS: -I./
#cgo LDFLAGS: -L./ -lhello
#include "hello.h"
*/
import "C"
import (
    "fmt"
)
 
func main() {
    C.print()
    fmt.Println("vim-go")
}
```

在go文件的头文件部分需要指定cgo需要的include路径，lib库的路径，lib库，头文件等等。

然后编译

```go
go build hg.go
```

编译生成 hg，是可执行文件，通过 ldd 可以验证是否正确加载了 libhello.so 的库

#### 类型转换

在 Go 语言中访问 C 语言的符号时，一般是通过虚拟的 “C” 包访问，比如 C.int 对应 C 语言的int类型。

