### strings包NewReader

函数作用

NewReader创建一个从s读取数据的Reader

方法

Len作用: 返回未读的字符串长度

Size的作用:返回字符串的长度

Read的作用: 读取字符串信息



```go
func NewScanner(r io.Reader) *Scanner {
    return &Scanner{
        r:            r,
        split:        ScanLines,
        maxTokenSize: MaxScanTokenSize,
    }
}
```

函数NewScanner返回一个Scanner，这个返回值取决于函数参数r。他的类型是io.Reader。

对于Scanner.Scan方法，相当于其他语言的迭代器iterator，并把迭代器指向的数据存放到新的缓冲区里。新的缓冲区(token)可以通过scanner.Text()或者scanner.Bytes()获取到。

```golang
func main(){
    scanner:=bufio.NewScanner(
        strings.NewReader("ABCDEFG\nHIJKELM"),
    )
    for scanner.Scan(){
        fmt.Println(scanner.Text()) // scanner.Bytes()
    }
}
```

Scanner.Scan方法默认是以换行符`\n`，作为分隔符。如果你想指定分隔符，Go语言提供了四种方法，ScanBytes(`返回单个字节作为一个 token`), ScanLines(`返回一行文本`), ScanRunes(`返回单个 UTF-8 编码的 rune 作为一个 token`)和ScanWords(`返回通过“空格”分词的单词`)。 

```golang
func main(){
    scanner:=bufio.NewScanner(
        strings.NewReader("ABCDEFG\nHIJKELM"),
    )
    scanner.Split(ScanWords/*四种方式之一，你也可以自定义, 实现SplitFunc方法*/)
    for scanner.Scan(){
        fmt.Println(scanner.Text()) // scanner.Bytes()
    }
}
```

