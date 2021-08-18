#### log

Go语言内置`log`包实现了简单的日志服务。

##### 使用logger

log包定义了Logger类型，该类型提供了一些格式化输出的方法。本包也提供了一个预定义的“标准”logger，可以通过调用函数`Print系列`(Print|Printf|Println）、`Fatal系列`（Fatal|Fatalf|Fatalln）、和`Panic系列`（Panic|Panicf|Panicln）来使用，比自行创建一个logger对象更容易使用。

例如，我们可以像下面的代码一样直接通过`log`包来调用上面提到的方法，默认它们会将日志信息打印到终端界面：

```go
package main

import (
	"log"
)

func main() {
	log.Println("这是一条很普通的日志。")
	v := "很普通的"
	log.Printf("这是一条%s日志。\n", v)
	log.Fatalln("这是一条会触发fatal的日志。")
	log.Panicln("这是一条会触发panic的日志。")
}
```

编译并执行上面的代码会得到如下输出：

```bash
2017/06_socket/19 14:04:17 这是一条很普通的日志。
2017/06_socket/19 14:04:17 这是一条很普通的日志。
2017/06_socket/19 14:04:17 这是一条会触发fatal的日志。
```

<font color='red'>logger会打印每条日志信息的日期、时间，默认输出到系统的标准错误。</font>Fatal系列函数会在写入日志信息后调用os.Exit(1)。Panic系列函数会在写入日志信息后panic。

##### 配置logger

默认情况下的logger只会提供日志的时间信息，但是很多情况下我们希望得到更多信息，比如记录该日志的文件名和行号等。`log`标准库中为我们提供了定制这些设置的方法。

`log`标准库中的`Flags`函数会返回标准logger的输出配置，而`SetFlags`函数用来设置标准logger的输出配置。

```go
func Flags() int
func SetFlags(flag int)
```

###### flag选项

`log`标准库提供了如下的flag选项，它们是一系列定义好的常量。

```go
const (
    // 控制输出日志信息的细节，不能控制输出的顺序和格式。
    // 输出的日志在每一项后会有一个冒号分隔：例如2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
    Ldate         = 1 << iota     // 日期：2009/01/23
    Ltime                         // 时间：01:23:23
    Lmicroseconds                 // 微秒级别的时间：01:23:23.123123（用于增强Ltime位）
    Llongfile                     // 文件全路径名+行号： /a/b/c/d.go:23
    Lshortfile                    // 文件名+行号：d.go:23（会覆盖掉Llongfile）
    LUTC                          // 使用UTC时间
    LstdFlags     = Ldate | Ltime // 标准logger的初始值
)
```

下面我们在记录日志之前先设置一下标准logger的输出选项如下：

```go
func main() {
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
}
```

编译执行后得到的输出结果如下：

```go
2017/06/19 14:05:17.494943 .../log_demo/main.go:11: 这是一条很普通的日志。
```

##### 配置日志输出位置

```go
func SetOutput(w io.Writer)
```

`SetOutput`函数用来设置标准logger的输出目的地，默认是标准错误输出。

如果你要使用标准的logger，我们通常会把上面的配置操作写到`init`函数中。

```go
func init() {
	logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
}
```

#### 创建logger

`log`标准库中还提供了一个创建新logger对象的构造函数–`New`，支持我们创建自己的logger示例。`New`函数的签名如下：

```go
func New(out io.Writer, prefix string, flag int) *Logger
```

New创建一个Logger对象。其中，参数out设置日志信息写入的目的地。参数prefix会添加到生成的每一条日志前面。参数flag定义日志的属性（时间、文件等等）。

举个例子：

```go
func main() {
	logger := log.New(os.Stdout, "<New>", log.Lshortfile|log.Ldate|log.Ltime)
	logger.Println("这是自定义的logger记录的日志。")
}
```

将上面的代码编译执行之后，得到结果如下：

```bash
<New>2017/06_socket/19 14:06_socket:51 main.go:34: 这是自定义的logger记录的日志。
```

#### runtime.Caller

```go
func Caller(skip int) (pc uintptr, file string, line int, ok bool)
```

`runtime.Caller`能够拿到当前执行的文件名和行号，这个方法几乎所有的日志组件里都有使用。<font color='red'>实参`skip`为上溯的栈帧数，0表示`Caller`的调用者（`Caller`所在的调用栈）。</font>函数的返回值为调用栈标识符、带路径的完整文件名、该调用在文件中的行号。如果无法获得信息，ok会被设为false。

##### 根据pc获取函数名

根据返回的调用帧标识符`pc`，我们可以获取对应的函数名。

需要用到一个函数和一个方法，签名如下：

```go
func FuncForPC(pc uintptr) *Func

func (*Func) Name
```

`FuncForPC`返回一个调用栈标识符`pc`对应的调用栈的`*Func`；如果该调用栈标识符没有对应的调用栈，函数会返回`nil`。每一个调用栈必然是对某个函数的调用。
`Name`返回该调用栈所调用的函数的名字。

```go
package main

import (
	"fmt"
	"path"
	"runtime"
)

func getInfo(skip int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("runtime.Caller() failed")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file) // Base函数返回路径的最后一个元素
	return
}

func main() {
	fmt.Println(getInfo(0))
	fmt.Println(getInfo(1))
}
```

输出：

```go
main.getInfo main.go 10
main.main main.go 22
```

