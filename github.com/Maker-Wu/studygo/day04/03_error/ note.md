错误指的是可能出现问题的地方出现了问题，比如打开一个文件时可能失败，这种情况在人们的意料之中 ；而异常指的是不应该出现问题的地方出现了问题，比如引用了空指针，这种情况在人们的意料之外。可见， **错误是业务逻辑的一部分，而异常不是** 。

#### 错误处理

我们编写一个简单的程序，该程序试图打开一个不存在的文件：

```go
package main

import (  
    "fmt"
    "os"
)

func main() {  
    f, err := os.Open("/test.txt")
    if err != nil {
        fmt.Println("error:",err)
        return
    }
    fmt.Println(f.Name(), "open successfully")
}
```

可以看到我们的程序调用了os包的Open方法，该方法定义如下：

```go
// Open opens the named file for reading. If successful, methods on
// the returned file can be used for reading; the associated file
// descriptor has mode O_RDONLY.
// If there is an error, it will be of type *PathError.
func Open(name string) (*File, error) {
    return OpenFile(name, O_RDONLY, 0)
}
```

参考注释可以知道如果这个方法正常返回的话会返回一个可读的文件句柄和一个值为 *nil* 的错误，如果该方法未能成功打开文件会返回一个*PathError类型的错误。

在Go语言中，处理错误时通常都是将返回的错误与 nil 比较。nil 值表示了没有错误发生，而非 nil 值表示出现了错误。于是有个我们上面那行代码：

```go
if err != nil {
    fmt.Println("error:", err)
    return
}
```

运行程序，结果显示：

```go
error: open /test.txt: No such file or directory
```



#### error类型

Go中返回的error类型究竟是什么呢？看源码发现error类型是一个非常简单的接口类型，具体如下：

```go
// The error built-in interface type is the conventional interface for
// representing an error condition, with the nil value representing no error.
type error interface {
    Error() string
}
```

error 有一个 Error() string 的方法。所有实现该接口的类型都可以当作一个错误类型。Error() 方法给出了错误的描述。
fmt.Println 在打印错误时，会在内部调用 Error() string 方法来得到该错误的描述。上一节示例中的错误描述就是这样打印出的。



#### 自定义错误类型

现在我们回到刚才代码里的*PathError类型，首先显而易见os.Open方法返回的错误是一个error类型，故我们可以知道PathError类型一定实现了error类型，也就是说实现了Error方法。现在我们看下具体实现

```go
type PathError struct {
    Op   string
    Path string
    Err  error
}

func (e *PathError) Error() string { return e.Op + " " + e.Path + ": " + e.Err.Error() }
```

可以看到PathError类型实现了Error方法，该方法返回文件操作、路径及error字符串的拼接返回值。

为什么需要自定义错误类型呢，试想一下如果一个错误我们拿到的仅仅是错误的字符串描述，那显然无法从错误中获取更多的信息或者做一些逻辑相关的校验，这样我们就可以通过自定义错误的结构体，通过实现Error()来使该结构体成为一个错误类型，使用时做一下类型推荐，我们就可以从返回的错误通过结构体中的一些成员就可以做逻辑校验或者错误分类等工作。例如：

```go
package main

import (  
    "fmt"
    "os"
)

func main() {  
    f, err := os.Open("/test.txt")
    if err, ok := err.(*os.PathError); ok {
        fmt.Println("File at path", err.Path, "failed to open")
        return
    }
    fmt.Println(f.Name(), "opened successfully")
}
```

上面代码中我们通过将error类型推断为实际的PathError类型，就可以拿到发生错误的Op、Path等数据，更有助于实际场景中错误的处理。

我们组现在拉通了一套错误类型和错误码规范，之前工程里写的时候都是通过在代码中的controller里面去根据不同情况去返回，这种处理方法有很多缺点，例如下层仅返回一个error类型，上层怎么判断该错误是哪种错误，该使用哪种错误码呢？另外就是程序中靠程序员写死某个逻辑错误码为xxxx，使程序缺乏稳定性，错误码返回也较为随心所欲，因此我也去自定义了错误，具体如下：

```go
var (
    ErrSuccess           = StandardError{0, "成功"}
    ErrUnrecognized      = StandardError{-1, "未知错误"}
    ErrAccessForbid      = StandardError{1000, "没有访问权限"}
    ErrNamePwdIncorrect  = StandardError{1001, "用户名或密码错误"}
    ErrAuthExpired       = StandardError{1002, "证书过期"}
    ErrAuthInvalid       = StandardError{1003, "无效签名"}
    ErrClientInnerError  = StandardError{4000, "客户端内部错误"}
    ErrParamError        = StandardError{4001, "参数错误"}
    ErrReqForbidden      = StandardError{4003, "请求被拒绝"}
    ErrPathNotFount      = StandardError{4004, "请求路径不存在"}
    ErrMethodIncorrect   = StandardError{4005, "请求方法错误"}
    ErrTimeout           = StandardError{4006, "服务超时"}
    ErrServerUnavailable = StandardError{5000, "服务不可用"}
    ErrDbQueryError      = StandardError{5001, "数据库查询错误"}
)

//StandardError 标准错误，包含错误码和错误信息
type StandardError struct {
    ErrorCode int    `json:"errorCode"`
    ErrorMsg  string `json:"errorMsg"`
}

// Error 实现了 Error接口
func (err StandardError) Error() string {
    return fmt.Sprintf("errorCode: %d, errorMsg %s", err.ErrorCode, err.ErrorMsg)
}
```

这样通过直接取StandardError的ErrorCode就可以知道应该返回的错误信息及错误码，调用时候也较为方便，并且做到了标准化，解决了之前项目中错误处理的问题。

