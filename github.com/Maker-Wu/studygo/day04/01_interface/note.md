#### 值接收者和指针接收者实现接口的区别

有一个Mover接口和一个dog结构体

```go
type Mover interface {
	move()
}

type dog struct {}
```

##### 值接收者实现接口

```go
func (d dog) move() {
	fmt.Println("狗会动")
}

func main() {
	var x Mover
	var wangcai = dog{} // 旺财是dog类型
	x = wangcai         // x可以接收dog类型
	var fugui = &dog{}  // 富贵是*dog类型
	x = fugui           // x可以接收*dog类型
	x.move()
}
```

从上面的代码中我们可以发现，使用值接收者实现接口之后，不管是dog结构体还是结构体指针*dog类型的变量都可以赋值给该接口变量。因为Go语言中有对指针类型变量求值的语法糖，dog指针`fugui`内部会自动求值`*fugui`。



##### 指针接收者实现接口

```go
func (d *dog) move() {
	fmt.Println("狗会动")
}
func main() {
	var x Mover
	var wangcai = dog{} // 旺财是dog类型
	x = wangcai         // x不可以接收dog类型
	var fugui = &dog{}  // 富贵是*dog类型
	x = fugui           // x可以接收*dog类型
}
```

此时实现`Mover`接口的是`*dog`类型，所以不能给`x`传入`dog`类型的wangcai，此时x只能存储`*dog`类型的值。



#### 空接口

空接口是指没有定义任何方法的接口。因此任何类型都实现了空接口

空接口类型的变量可以存储任意类型的变量。

```go
func main() {
	// 定义一个空接口x
	var x interface{}
	s := "Hello 沙河"
	x = s
	fmt.Printf("type:%T value:%v\n", x, x)
	i := 100
	x = i
	fmt.Printf("type:%T value:%v\n", x, x)
	b := true
	x = b
	fmt.Printf("type:%T value:%v\n", x, x)
}
```

运行结果：

```go\
type:string value:Hello 沙河
type:int value:100
type:bool value:true
```

##### 空接口的应用

使用空接口实现可以接收任意类型的函数参数

```go
// 空接口作为函数参数
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}
```

使用空接口实现可以保存任意值的字典。

```go
// 空接口作为map值
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "沙河娜扎"
	studentInfo["age"] = 18
	studentInfo["married"] = false
	fmt.Println(studentInfo)	//map[age:18 married:false name:沙河娜扎]
```



#### 接口类型

空接口可以存储任意类型的值，那我们如何获取其存储的具体数据呢？

##### 接口值

接口类型变量除了静态类型之外还有动态类型，动态类型是由给接口类型变量赋值的具体值的类型来决定的，除了动态类型之外还有动态值。一个接口的值（简称接口值）是由<font color='red'>一个具体类型和具体类型的值</font>两部分组成的。这两部分分别称为接口的动态类型和动态值，之所以被称为动态类型，是因为接口类型的动态类型是会变化的，由被赋予的值来决定。

```go
var w io.Writer
w = os.Stdout
w = new(bytes.Buffer)
w = nil
```

![接口值图解](note.assets/interface.png)

接口类型变量赋值之前直接获取它的类型会发现返回的结果是`nil`

```go
package main
import "fmt"
func main()  {
    // 动态类型
    var b Bird
    fmt.Printf("interface类型变量的类型：%T\n", b)
}
type Bird interface {
    fly()
    sing()
}

执行结果：
interface类型变量的类型：<nil>
```

首先我们需要明确，一个接口类型变量在没有被赋值之前，它的动态类型和动态值都是 *nil* 。在使用 `fmt.Printf("%T\n")` 获取一个变量的类型时，其实是调用了reflect包的方法进行获取的，*reflect.TypeOf* 获取的是接口变量的动态类型， *reflect.valueOf()* 获取的是接口变量的动态值。所以 *fmt.Printf("%T\n",b)* 展示的是 *reflect.TypeOf* 的结果，由于接口变量 *b* 还没有被赋值，所以它的动态类型是 *nil* ，动态值也会是 *nil* 。

对比来看，为什么只是经过了声明未赋值的变量的类型不是 *nil* 呢？就像在静态类型部分中所展示的那样。原因如下： 我们先来看一下 *reflect.TypeOf* 函数的定义，*func TypeOf(i interface{}) Type{}* ，函数的参数是一个 *interface* 类型的变量，在调用 *TypeOf* 时，在接口变量 *b* 没有赋值之前，它的静态类型与参数类型一致，不需要做转换，因为 *b* 的动态类型为 *nil*，所以 *TypeOf* 返回的结果为 *nil* 。那为什么变量 *i* 和变量 *g* 的类型不为 *nil* 呢？当变量 *i* 调用 *TypeOf* 时，会进行类型的转换，将int型变量i转换为 *interface* 型，在这个过程中会将变量 *i* 的类型作为 *b* 的动态类型，变量 *i* 的值（在这里是变量 *i* 的零值0）作为 *b* 的动态值。因为 *TypeOf()* 获取的是变量 *b* 的动态类型，所以这个时候展示出的类型为 *int*。

#### 类型断言

想要判断空接口中的值这个时候就可以使用类型断言，其语法格式：

```go
x.(T)
```

其中：

- x：表示类型为`interface{}`的变量
- T：表示断言`x`可能是的类型。

该语法返回两个参数，第一个参数是`x`转化为`T`类型后的变量，第二个值是一个布尔值，若为`true`则表示断言成功，为`false`则表示断言失败。

```go
func main() {
	var x interface{}
	x = "Hello 沙河"
	v, ok := x.(string)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败")
	}
}
```

上面的示例中如果要断言多次就需要写多个`if`判断，这个时候我们可以使用`switch`语句来实现：

```go
func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}
```

