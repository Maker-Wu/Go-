#### 结构体tag介绍

`Tag`是结构体的元信息，可以在运行的时候通过反射的机制读取出来。`Tag`在结构体字段的后方定义，由一对**反引号**包裹起来，具体的格式如下：

```go
`key1:"value1" key2:"value2"`
```

结构体tag由一个或多个键值对组成。键与值使用**冒号**分隔，值用**双引号**括起来。同一个结构体字段可以设置多个键值对tag，不同的键值对之间使用**空格**分隔。

##### 使用json tag指定字段名

```go
// 使用json tag指定序列化与反序列化时的行为
type Person struct {
	Name   string `json:"name"` // 指定json序列化/反序列化时使用小写name
	Age    int64
	Weight float64
}
```

##### 忽略某个字段

如果你想在json序列化/反序列化的时候忽略掉结构体中的某个字段，可以按如下方式在tag中添加`-`。

```go
// 使用json tag指定json序列化与反序列化时的行为
type Person struct {
	Name   string `json:"name"` // 指定json序列化/反序列化时使用小写name
	Age    int64
	Weight float64 `json:"-"` // 指定json序列化/反序列化时忽略此字段
}
```

##### 忽略空值字段

当 struct 中的字段没有值时， `json.Marshal()` 序列化的时候不会忽略这些字段，而是默认输出字段的类型零值（例如`int`和`float`类型零值是 0，`string`类型零值是`""`，对象类型零值是 nil）。如果想要在序列序列化时忽略这些没有值的字段时，可以在对应字段添加`omitempty` tag。

```go
// 在tag中添加omitempty忽略空值
// 注意这里 hobby,omitempty 合起来是json tag值，中间用英文逗号分隔
type User struct {
	Name  string   `json:"name"`
	Email string   `json:"email,omitempty"`
	Hobby []string `json:"hobby,omitempty"`
}
```

此时，再执行上述的`omitemptyDemo`，输出结果如下：

```go
str:{"name":"七米"} // 序列化结果中没有email和hobby字段
```

##### 忽略嵌套结构体空值字段

```go
type User struct {
	Name  string   `json:"name"`
	Email string   `json:"email,omitempty"`
	Hobby []string `json:"hobby,omitempty"`
	Profile
}

type Profile struct {
	Website string `json:"site"`
	Slogan  string `json:"slogan"`
}

func nestedStructDemo() {
	u1 := User{
		Name:  "七米",
		Hobby: []string{"足球", "双色球"},
	}
	b, err := json.Marshal(u1)
	if err != nil {
		fmt.Printf("json.Marshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)
}
```

匿名嵌套`Profile`时序列化后的json串为单层的：

```go
str:{"name":"七米","hobby":["足球","双色球"],"site":"","slogan":""}
```

想要变成嵌套的json串，需要改为具名嵌套或定义字段tag：

```go
type User struct {
	Name    string   `json:"name"`
	Email   string   `json:"email,omitempty"`
	Hobby   []string `json:"hobby,omitempty"`
	Profile `json:"profile"`
}
// str:{"name":"七米","hobby":["足球","双色球"],"profile":{"site":"","slogan":""}}
```

想要在嵌套的结构体为空值时，忽略该字段，仅添加`omitempty`是不够的：

```go
type User struct {
	Name     string   `json:"name"`
	Email    string   `json:"email,omitempty"`
	Hobby    []string `json:"hobby,omitempty"`
	Profile `json:"profile,omitempty"`
}
// str:{"name":"七米","hobby":["足球","双色球"],"profile":{"site":"","slogan":""}}
```

还需要使用嵌套的结构体指针：

```go
type User struct {
	Name     string   `json:"name"`
	Email    string   `json:"email,omitempty"`
	Hobby    []string `json:"hobby,omitempty"`
	*Profile `json:"profile,omitempty"`
}
// str:{"name":"七米","hobby":["足球","双色球"]}
```



