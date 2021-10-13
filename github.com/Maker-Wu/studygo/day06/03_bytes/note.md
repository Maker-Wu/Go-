bytes.Buffer是一个缓冲byte类型的缓冲器存放着都是byte。Buffer 是 bytes 包中的一个 type Buffer struct{…}，是一个变长的 buffer(其实底层就是一个[]byte)，具有 Read 和Write 方法。 Buffer 的 零值 是一个 空的 buffer，但是可以使用。

### 创建Buffer缓冲器

```go
var b bytes.Buffer	//直接定义一个Buffer变量，而不用初始化
b.Write([]byte("Hello ")) //可以直接使用

b1 := new(bytes.Buffer) //直接使用new初始化，可以直接使用
//其他两种初始化方式
func NewBuffer(buf []byte) *Buffer
func NewBufferString(s string) *Buffer
```

#### NewBuffer

```go
func NewBuffer(buf []byte) *Buffer { return &Buffer{buf: buf} }
```

- Buffer既可以被读也可以被写

#### NewBufferString

```go
func NewBufferString(s string) *Buffer {
    return &Buffer{buf: []byte(s)}
}
```

- 方法NewBufferString用一个string来初始化可读Buffer，并用string的内容填充Buffer.

### 向 Buffer 中写入数据

#### Write

把字节切片 p 写入到buffer中去。

#### WriteString

使用WriteString方法，将一个字符串放到缓冲器的尾部