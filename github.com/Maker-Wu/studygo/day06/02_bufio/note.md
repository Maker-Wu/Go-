### bufio.Reader和bufio.Write

bufio包有两个New函数，分别是在任意io.Reader和io.Writer的基础上再包装一层缓冲区得到bufio.Reader和bufio.Writer。

```go
func NewReader(rd io.Reader) *Reader
func NewWriter(w io.Writer) *Writer
```

#### bufio.Write

多次进行小量的写操作会影响程序性能。每一次写操作最终都会体现为系统层调用，频繁进行该操作将有可能对 CPU 造成伤害。而且很多硬件设备更适合处理块对齐的数据，例如硬盘。为了减少进行多次写操作所需的开支，golang 提供了bufio.Write。数据将不再直接写入目的地(实现了 [io.Writer](https://golang.org/pkg/io/#Writer) 接口)，而是先写入缓存，当缓存写满后再统一写入目的地：

```go
producer --> buffer --> io.Writer
```

下面具体看一下在9次写入操作中(每次写入一个字符)具有4个字符空间的缓存是如何工作的：

```go
producer        buffer       destination (io.Writer)
   a    ----->    a
   b    ----->    ab
   c    ----->    abc
   d    ----->    abcd
   e    ----->    e      ----->   abcd
   f    ----->    ef
   g    ----->    efg
   h    ----->    efgh
   i    ----->    i      ----->   abcdefgh
```

`----->` 箭头代表写入操作

<font color='red'>bufio.Write底层使用 `[]byte` 进行缓存</font>

```go
type Writer int

func (*Writer) Write(p []byte) (n int, err error) {
    fmt.Println(len(p))
    return len(p), nil
}

func main() {
    fmt.Println("Unbuffered I/O")
    w := new(Writer)
    w.Write([]byte{'a'})
    w.Write([]byte{'b'})
    w.Write([]byte{'c'})
    w.Write([]byte{'d'})
    fmt.Println("Buffered I/O")
    bw := bufio.NewWriterSize(w, 3)
    bw.Write([]byte{'a'})
    bw.Write([]byte{'b'})
    bw.Write([]byte{'c'})
    bw.Write([]byte{'d'})
    err := bw.Flush()
    if err != nil {
        panic(err)
    }
}
```

运行结果：

```go
Unbuffered I/O
1
1
1
1
Buffered I/O
3
1
```

没有被缓存的 `I/O`：意味着每一次写操作都将直接写入目的地。我们进行4次写操作，每次写操作都映射为对 `Write` 的调用，调用时传入的参数为一个长度为1的 `byte` 切片。

使用了缓存的 `I/O`：我们使用三个字节长度的缓存来存储数据，当缓存满时进行一次 `flush` 操作(将缓存中的数据进行处理)。前三次写入写满了缓存。第四次写入时检测到缓存没有剩余空间，所以将缓存中的积累的数据写出。字母 `d` 被存储了，但在此之前 `Flush` 被调用以腾出空间。当缓存被写到末尾时，缓存中未被处理的数据需要被处理。<font color='red'>`bufio.Writer` 仅在缓存充满或者显式调用 `Flush` 方法时处理(发送)数据</font>。

> `bufio.Writer` 默认使用 4096 长度字节的缓存，可以使用NewWriterSize方法来设定该值

##### 实现

实现十分简单：

```go
type Writer struct {
    err error
    buf []byte
    n   int
    wr  io.Writer
}
```

字段 `buf` 用来存储数据，当缓存满或者 `Flush` 被调用时，消费者(`wr`)可以从缓存中获取到数据。<font color='red'>如果写入过程中发生了 I/O error，此 error 将会被赋给 `err` 字段， error 发生之后，writer 将停止操作(writer is no-op)</font>：

```go
type Writer int

func (*Writer) Write(p []byte) (n int, err error) {
    fmt.Printf("Write: %q\n", p)
    return 0, errors.New("boom!")
}

func main() {
    w := new(Writer)
    bw := bufio.NewWriterSize(w, 3)
    bw.Write([]byte{'a'})
    bw.Write([]byte{'b'})
    bw.Write([]byte{'c'})
    bw.Write([]byte{'d'})
    err := bw.Flush()
    fmt.Println(err)
}
Write: "abc"
boom!
```

这里我们可以看到 `Flush` 没有第二次调用消费者的 `write` 方法。如果发生了 error， 使用了缓存的 writer 不会尝试再次执行写操作。

字段 `n` 标识缓存内部当前操作的位置。`Buffered` 方法返回 `n` 的值：

```go
type Writer int
func (*Writer) Write(p []byte) (n int, err error) {
    return len(p), nil
}
func main() {
    w := new(Writer)
    bw := bufio.NewWriterSize(w, 3)
    fmt.Println(bw.Buffered())
    bw.Write([]byte{'a'})
    fmt.Println(bw.Buffered())
    bw.Write([]byte{'b'})
    fmt.Println(bw.Buffered())
    bw.Write([]byte{'c'})
    fmt.Println(bw.Buffered())
    bw.Write([]byte{'d'})
    fmt.Println(bw.Buffered())
}
```

运行结果：

```go
0
1
2
3
1
```

`n` 从 0 开始，当有数据被添加到缓存中时，该数据的长度值将会被加和到 `n`中(操作位置向后移动)。当`bw.Write([] byte{'d'})`被调用时，flush会被触发，`n` 会被重设为0。

##### Large writes

```go
type Writer int
func (*Writer) Write(p []byte) (n int, err error) {
    fmt.Printf("%q\n", p)
    return len(p), nil
}
func main() {
    w := new(Writer)
    bw := bufio.NewWriterSize(w, 3)
    bw.Write([]byte("abcd"))
}
```

运行结果：

```go
"abcd"
```

由于使用了 `bufio`，程序打印了 `"abcd"`。如果 `Writer` 检测到 `Write` 方法被调用时传入的数据长度大于缓存的长度(示例中是三个字节)。其将直接调用 writer(目的对象)的 `Write` 方法。当数据量足够大时，其会自动跳过内部缓存代理。

##### 重置

缓存是 `bufio` 的核心部分。通过使用 `Reset` 方法，`Writer` 可以用于不同的目的对象。重复使用 `Writer` 缓存减少了内存的分配。而且减少了额外的垃圾回收工作：

```go
type Writer1 int
func (*Writer1) Write(p []byte) (n int, err error) {
    fmt.Printf("writer#1: %q\n", p)
    return len(p), nil
}
type Writer2 int
func (*Writer2) Write(p []byte) (n int, err error) {
    fmt.Printf("writer#2: %q\n", p)
    return len(p), nil
}
func main() {
    w1 := new(Writer1)
    bw := bufio.NewWriterSize(w1, 2)
    bw.Write([]byte("ab"))
    bw.Write([]byte("cd"))
    w2 := new(Writer2)
    bw.Reset(w2)
    bw.Write([]byte("ef"))
    bw.Flush()
}
writer#1: "ab"
writer#2: "ef"
```

这段代码中有一个 bug。在调用 `Reset` 方法之前，我们应该使用 `Flush` flush缓存。 由于 [`Reset`](https://github.com/golang/go/blob/7b8a7f8272fd1941a199af1adb334bd9996e8909/src/bufio/bufio.go#L559) 只是简单的丢弃未被处理的数据，所以已经被写入的数据 `cd` 丢失了：

```go
func (b *Writer) Reset(w io.Writer) {
    b.err = nil
    b.n = 0
    b.wr = w
}
```

##### 缓存剩余空间

为了检测缓存中还剩余多少空间，我们可以使用方法`Available`

```go
w := new(Writer)
bw := bufio.NewWriterSize(w, 2)
fmt.Println(bw.Available())
bw.Write([]byte{'a'})
fmt.Println(bw.Available())
bw.Write([]byte{'b'})
fmt.Println(bw.Available())
bw.Write([]byte{'c'})
fmt.Println(bw.Available())
```

运行结果

```go
2
1
0
1
```

##### 写`{Byte,Rune,String}`的方法

为了方便, 我们有三个用来写普通类型的实用方法：

```go
w := new(Writer)
bw := bufio.NewWriterSize(w, 10)
fmt.Println(bw.Buffered())
bw.WriteByte('a')
fmt.Println(bw.Buffered())
bw.WriteRune('ł') // 'ł' occupies 2 bytes
fmt.Println(bw.Buffered())
bw.WriteString("aa")
fmt.Println(bw.Buffered())
0
1
3
5
```

#### ReadFrom

io 包中定义了 io.ReaderFrom接口。 该接口通常被 writer 实现，用于从指定的 reader 中读取所有数据(直到 EOF)并对读到的数据进行底层处理：

```go
type ReaderFrom interface {
    ReadFrom(r Reader) (n int64, err error)
}
```

> 比如 [`io.Copy`](https://golang.org/pkg/io/#Copy) 使用了 `io.ReaderFrom` 接口

`bufio.Writer` 实现了此接口：因此我们可以通过调用 `ReadFrom` 方法来处理从 `io.Reader` 获取到的所有数据：

```go
type Writer int
func (*Writer) Write(p []byte) (n int, err error) {
    fmt.Printf("%q\n", p)
    return len(p), nil
}
func main() {
    s := strings.NewReader("onetwothree")
    w := new(Writer)
    bw := bufio.NewWriterSize(w, 3)
    bw.ReadFrom(s)
    err := bw.Flush()
    if err != nil {
        panic(err)
    }
}
```

运行结果：

```go
"one"
"two"
"thr"
"ee"
```

> 使用 `ReadFrom` 方法的同时，调用 `Flush` 方法也很重要

#### bufio.Reader

```go
type Reader struct {
    		buf          []byte
    		rd           io.Reader // reader provided by the client
    		r, w         int       // buf read and write positions
    		err          error
    		lastByte     int
    		lastRuneSize int
    	}
```

通过它，我们可以从底层的 `io.Reader` 中更大批量的读取数据。

```go
io.Reader --> buffer --> consumer
```

假设消费者想要从硬盘上读取10个字符(每次读取一个字符)。在底层实现上，这将会触发10次读取操作。如果硬盘按每个数据块四个字节来读取数据，那么 `bufio.Reader` 将会起到帮助作用。底层引擎将会缓存整个数据块，然后提供一个可以挨个读取字节的 API 给消费者：

```go
abcd -----> abcd -----> a
            abcd -----> b
            abcd -----> c
            abcd -----> d
efgh -----> efgh -----> e
            efgh -----> f
            efgh -----> g
            efgh -----> h
ijkl -----> ijkl -----> i
            ijkl -----> j
```

`----->` 代表读取操作
这个方法仅需要从硬盘读取三次，而不是10次。

##### Peek

Peek 返回缓存的一个切片，该切片引用缓存中前 n 字节数据，该操作不会将数据读出，只是引用。

- 如果缓存不满，而且缓存中缓存的数据少于 `n` 个字节，其将会尝试从 `io.Reader` 中读取
- <font color='red'>如果请求的数据量大于缓存的容量，将会返回 </font>`bufio.ErrBufferFull`
- 如果 `n` 大于流的大小，将会返回 EOF

```go
s1 := strings.NewReader(strings.Repeat("a", 20))
r := bufio.NewReaderSize(s1, 16)
b, err := r.Peek(3)
if err != nil {
    fmt.Println(err)
}
fmt.Printf("%q\n", b)
b, err = r.Peek(17)
if err != nil {
    fmt.Println(err)
}
s2 := strings.NewReader("aaa")
r.Reset(s2)
b, err = r.Peek(10)
if err != nil {
    fmt.Println(err)
}
"aaa"
bufio: buffer full
EOF
```

> 被 `bufio.Reader` 使用的最小的缓存容器是 16。

<font color='red'>返回的切片和被 `bufio.Reader` 使用的内部缓存底层使用相同的数组。</font>因此引擎底层在执行任何读取操作之后内部返回的切片将会变成无效的。这是由于其将有可能被其他的缓存数据覆盖：

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

##### Read

`Read` 方法是 `bufio.Reader` 的核心。它和 io.Reader 的唯一方法具有相同的签名。因此 `bufio.Reader` 实现了这个普遍存在的接口：

```go
type Reader interface {
        Read(p []byte) (n int, err error)
}
```

`bufio.Reader` 的 `Read` 方法从底层的 `io.Reader` 中一次读 取最大的数量:

1. 如果内部缓存具有至少一个字节的数据，那么无论传入的切片的大小(`len(p)`)是多少，`Read` 方法都将仅仅从内部缓存中获取数据，不会从底层的 reader 中读取任何数据:

   ```go
   func (r *R) Read(p []byte) (n int, err error) {
    fmt.Println("Read")
    copy(p, "abcd")
    return 4, nil
   }
   func main() {
    r := new(R)
    br := bufio.NewReader(r)
    buf := make([]byte, 2)
    n, err := br.Read(buf)
    if err != nil {
        panic(err)
    }
    buf = make([]byte, 4)
    n, err = br.Read(buf)
    if err != nil {
        panic(err)
    }
    fmt.Printf("read = %q, n = %d\n", buf[:n], n)
   }
   Read
   read = "cd", n = 2
   ```

   我们的 `io.Reader` 实例无线返回「abcd」(不会返回 `io.EOF`)。 第二次调用 `Read`并传入长度为4的切片，但是内部缓存在第一次从 `io.Reader` 中读取数据之后已经具有数据「cd」，所以 `bufio.Reader` 返回缓存中的数据数据，而不和底层 reader 进行通信。

2. 如果内部缓存是空的，那么将会执行一次从底层 io.Reader 的读取操作。 从前面的例子中我们可以清晰的看到如果我们开启了一个空的缓存，然后调用:

   ```go
   n, err := br.Read(buf)
   ```

   将会触发读取操作来填充缓存。

3. 如果内部缓存是空的，但是传入的切片长度大于缓存长度，那么 `bufio.Reader` 将会跳过缓存，直接读取传入切片长度的数据到切片中:

   ```go
   type R struct{}
   func (r *R) Read(p []byte) (n int, err error) {
    fmt.Println("Read")
    copy(p, strings.Repeat("a", len(p)))
    return len(p), nil
   }
   func main() {
    r := new(R)
    br := bufio.NewReaderSize(r, 16)
    buf := make([]byte, 17)
    n, err := br.Read(buf)
    if err != nil {
        panic(err)
    }
    fmt.Printf("read = %q, n = %d\n", buf[:n], n)
    fmt.Printf("buffered = %d\n", br.Buffered())
   }
   Read
   read = "aaaaaaaaaaaaaaaaa", n = 17
   buffered = 0
   ```

   从 `bufio.Reader` 读取之后，内部缓存中没有任何数据(`buffered = 0`)
