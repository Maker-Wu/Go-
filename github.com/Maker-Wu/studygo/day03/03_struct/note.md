#### 使用值的列表初始化结构体

初始化结构体的时候可以简写，也就是初始化的时候不写键，直接写值

```go
p8 := &person{
	"沙河娜扎",
	"北京",
	28,
}
fmt.Printf("p8=%#v\n", p8)//p8=&main.go.person{name:"沙河娜扎", city:"北京", age:28}}
```

使用这种格式初始化时，需要注意：

1. 必须初始化结构体的所有字段。
2. 初始值的填充顺序必须与字段在结构体中的声明顺序一致。
3. 该方式不能和键值初始化方式混用。



#### 结构体内存布局

结构体占用一块连续的内存

```go
type test struct {
	a int8
	b int8
	c int8
	d int8
}
n := test{
	1, 2, 3, 4,
}
fmt.Printf("n.a %p\n", &n.a)
fmt.Printf("n.b %p\n", &n.b)
fmt.Printf("n.c %p\n", &n.c)
fmt.Printf("n.d %p\n", &n.d)
```

输出：

```go
n.a 0xc0000a0060
n.b 0xc0000a0061
n.c 0xc0000a0062
n.d 0xc0000a0063
```



#### 空结构体

空结构体不占用空间的

```go
var v struct{}
fmt.Println(unsafe.Sizeof(v))  // 0
```

#### for range踩坑

```go
type student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}
```

运行结果：

```go
小王子 => 大王八
娜扎 => 大王八
大王八 => 大王八
```

分析上述代码想要实现的功能是想用stus结构体数组来初始化map类型的m

发现全部都是==大王八==

这是因为在第一个循环中,m中插入的键值对的值都是&stu。由于range是值拷贝，也就是说三次循环传入的都是同一个地址。

改进第一个循环代码如下：

```go
for i, stu := range stus {
    m[stu.name] = &stus[i]
}
```



