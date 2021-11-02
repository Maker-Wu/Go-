#### 并发与并行

并发：同一时间段内执行多个任务（你在用微信和两个女朋友聊天）。

并行：同一时刻执行多个任务（你和你朋友都在用微信和女朋友聊天）。

Go语言的并发通过`goroutine`实现。`goroutine`类似于线程，属于用户态的线程，我们可以根据需要创建成千上万个`goroutine`并发工作。`goroutine`是由Go语言的运行时（runtime）调度完成，而线程是由操作系统调度完成。

#### 进程和线程

1. 进程就是程序在操作系统中的一个执行过程，是系统进行资源分配的基本单位。
2. 线程是进程的一个执行实例，是程序执行、系统调度的最小单位。
3. 一个进程可以创建和销毁多个线程，同一个进程中的多个线程可以并发执行。。
4. 一个程序至少有一个进程，一个进程至少有一个线程

#### goroutine

`goroutine`的概念类似于线程，可以理解为轻量级的线程。 `goroutine`是由Go的运行时（runtime）调度和管理的。Go程序会智能地将 goroutine 中的任务合理地分配给每个CPU。<font color='red'>Go语言之所以被称为现代化的编程语言，就是因为它在语言层面已经内置了调度和上下文切换的机制。</font>

特点：

- 有独立的栈空间
- 共享程序堆空间
- 调度由用户控制

##### 使用goroutine

Go语言中使用`goroutine`非常简单，只需要在调用函数的时候在前面加上`go`关键字，就可以为一个函数创建一个`goroutine`。

<font color='red'>一个`goroutine`必定对应一个函数，可以创建多个`goroutine`去执行相同的函数。</font>

启动多个goroutine

```go
var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Println("Hello Goroutine!", i)
}
func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go hello(i)
	}
	wg.Wait() // 等待所有登记的goroutine都结束
}
```

多次执行上面的代码，会发现每次打印的数字的顺序都不一致。这是因为10个`goroutine`是并发执行的，而`goroutine`的调度是随机的。

##### goroutine与线程

OS线程（操作系统线程）一般都有固定的栈内存（通常为2MB）,一个`goroutine`的栈在其生命周期开始时只有很小的栈（典型情况下2KB），`goroutine`的栈不是固定的，他可以按需增大和缩小，`goroutine`的栈大小限制可以达到1GB，虽然极少会用到这么大。所以在Go语言中一次创建十万左右的`goroutine`也是可以的。

##### GPM模式

`GPM`是Go语言运行时（runtime）层面的实现，是go语言自己实现的一套调度系统。区别于操作系统调度OS线程。

- `G`很好理解，就是个goroutine的，里面除了存放本goroutine信息外 还有与所在P的绑定等信息。
- `P`管理着一组goroutine队列，P里面会存储当前goroutine运行的上下文环境（函数指针，堆栈地址及地址边界），P会对自己管理的goroutine队列做一些调度（比如把占用CPU时间较长的goroutine暂停、运行后续的goroutine等等）当自己的队列消费完了就去全局队列里取，如果全局队列里也消费完了会去其他P的队列里抢任务。
- `M（machine）`是Go运行时（runtime）对操作系统内核线程的虚拟， M与内核线程一般是一一映射的关系， 一个groutine最终是要放到M上执行的；

P与M一般也是一一对应的。他们关系是： <font color='red'>P管理着一组G挂载在M上运行。</font>当一个G长久阻塞在一个M上时，runtime会新建一个M，阻塞G所在的P会把其他的G 挂载在新建的M上。当旧的G阻塞完成或者认为其已经死掉时 回收旧的M。

P的个数是通过`runtime.GOMAXPROCS`设定（最大256），Go1.5版本之后默认为物理线程数。 在并发量大的时候会增加一些P和M，但不会太多，切换太频繁的话得不偿失。

单从线程调度讲，Go语言相比起其他语言的优势在于OS线程是由OS内核来调度的，`goroutine`则是由Go运行时（runtime）自己的调度器调度的，这个调度器使用一个称为m:n调度的技术（复用/调度m个goroutine到n个OS线程）。 其一大特点是goroutine的调度是在用户态下完成的， 不涉及内核态与用户态之间的频繁切换，包括内存的分配与释放，都是在用户态维护着一块大的内存池， 不直接调用系统的malloc函数（除非内存池需要改变），成本比调度OS线程低很多。 另一方面充分利用了多核的硬件资源，近似的把若干goroutine均分在物理线程上， 再加上本身goroutine的超轻量，以上种种保证了go调度方面的性能。

##### GOMAXPROCS

Go运行时的调度器使用`GOMAXPROCS`参数来确定需要使用多少个OS线程来同时执行Go代码。默认值是机器上的CPU核心数。例如在一个8核心的机器上，调度器会把Go代码同时调度到8个OS线程上（GOMAXPROCS是m:n调度中的n）。

Go语言中可以通过`runtime.GOMAXPROCS()`函数设置当前程序并发时占用的CPU逻辑核心数。

Go1.5版本之前，默认使用的是单核心执行。Go1.5版本之后，默认使用全部的CPU逻辑核心数。

我们可以通过将任务分配到不同的CPU逻辑核心上实现并行的效果，这里举个例子：

```go
func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func main() {
	runtime.GOMAXPROCS(1)
	go a()
	go b()
	time.Sleep(time.Second)
}
```

两个任务只有一个逻辑核心，此时是做完一个任务再做另一个任务。 将逻辑核心数设为2，此时两个任务并行执行，代码如下。

```go
func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func main() {
	runtime.GOMAXPROCS(2)
	go a()
	go b()
	time.Sleep(time.Second)
}
```

Go语言中的操作系统线程和goroutine的关系：

1. 一个操作系统线程对应用户态多个goroutine。
2. go程序可以同时使用多个操作系统线程。
3. goroutine和OS线程是多对多的关系，即m:n。

#### channel

虽然可以使用共享内存进行数据交换，但是共享内存在不同的`goroutine`中容易发生竞态问题。为了保证数据交换的正确性，必须使用互斥量对内存进行加锁，这种做法势必造成性能问题。

Go语言的并发模型是`CSP（Communicating Sequential Processes）`，提倡<font color='red'>**通过通信共享内存**</font>而不是<font color='red'>**通过共享内存而实现通信**</font>。

如果说`goroutine`是Go程序并发的执行体，`channel`就是它们之间的连接。`channel`是可以让一个`goroutine`发送特定值到另一个`goroutine`的通信机制。

Go 语言中的通道（channel）是引用类型，通道类型的空值是`nil`。通道像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型。

```go
var ch1 chan int   // 声明一个传递整型的通道
var ch2 chan bool  // 声明一个传递布尔型的通道
var ch3 chan []int // 声明一个传递int切片的通道
```

声明的通道后需要使用`make`函数初始化之后才能使用。

```go
ch4 := make(chan int)
ch5 := make(chan bool)
ch6 := make(chan []int)
```

通道有发送（send）、接收(receive）和关闭（close）三种操作。

发送和接收都使用`<-`符号。

现在我们先使用以下语句定义一个通道：

```go
ch := make(chan int)
```

将一个值发送到通道中。

```go
ch <- 10 // 把10发送到ch中
```

从一个通道中接收值。

```go
x := <- ch // 从ch中接收值并赋值给变量x
<-ch       // 从ch中接收值，忽略结果
```

我们通过调用内置的`close`函数来关闭通道。

```go
close(ch)
```

关于关闭通道需要注意的事情是，只有在通知接收方goroutine所有的数据都发送完毕的时候才需要关闭通道。通道是可以被垃圾回收机制回收的，它和关闭文件是不一样的，在结束操作之后关闭文件是必须要做的，但关闭通道不是必须的。

关闭后的通道有以下特点：

1. 对一个关闭的通道再发送值就会导致panic。
2. 对一个关闭的通道进行接收会一直获取值直到通道为空。
3. 对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
4. 关闭一个已经关闭的通道会导致panic。

##### 无缓冲的通道

无缓冲的通道又称为阻塞的通道。我们来看一下下面的代码：

```go
func main() {
	ch := make(chan int)
	ch <- 10
	fmt.Println("发送成功")
}
```

上面这段代码能够通过编译，但是执行的时候会出现以下错误：

```bash
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
        .../src/github.com/Q1mi/studygo/day06/channel02/main.go:8 +0x54
```

为什么会出现`deadlock`错误呢？

因为我们使用`ch := make(chan int)`创建的是无缓冲的通道，<font color='red'>无缓冲的通道只有在有人接收值的时候才能发送值。</font>就像你住的小区没有快递柜和代收点，快递员给你打电话必须要把这个物品送到你的手中，简单来说就是无缓冲的通道必须有接收才能发送。

上面的代码会阻塞在`ch <- 10`这一行代码形成死锁，那如何解决这个问题呢？

一种方法是启用一个`goroutine`去接收值，例如：

```go
func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}
func main() {
	ch := make(chan int)
	go recv(ch) // 启用goroutine从通道接收值
	ch <- 10
	fmt.Println("发送成功")
}
```

无缓冲通道上的发送操作会阻塞，直到另一个`goroutine`在该通道上执行接收操作，这时值才能发送成功，两个`goroutine`将继续执行。相反，如果接收操作先执行，接收方的goroutine将阻塞，直到另一个`goroutine`在该通道上发送一个值。

使用无缓冲通道进行通信将导致发送和接收的`goroutine`同步化。因此，无缓冲通道也被称为`同步通道`。

##### 有缓冲的通道

解决上面问题的方法还有一种就是使用有缓冲区的通道。我们可以在使用make函数初始化通道的时候为其指定通道的容量，例如：

```go
func main() {
	ch := make(chan int, 1) // 创建一个容量为1的有缓冲区通道
	ch <- 10
	fmt.Println("发送成功")
}
```

只要通道的容量大于零，那么该通道就是有缓冲的通道，通道的容量表示通道中能存放元素的数量。就像你小区的快递柜只有那么个多格子，格子满了就装不下了，就阻塞了，等到别人取走一个快递员就能往里面放一个。

我们可以使用内置的`len`函数获取通道内元素的数量，使用`cap`函数获取通道的容量，虽然我们很少会这么做。

##### for range从通道循环取值

当向通道中发送完数据时，我们可以通过`close`函数来关闭通道。

当通道被关闭时，再往该通道发送值会引发`panic`，从该通道取值的操作会先取完通道中的值，再然后取到的值一直都是对应类型的零值。那如何判断一个通道是否被关闭了呢？

```go
// channel 练习
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 开启goroutine将0~100的数发送到ch1中
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			i, ok := <-ch1 // 通道关闭后再取值ok=false
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()
	// 在主goroutine中从ch2中接收值打印
	for i := range ch2 { // 通道关闭后会退出for range循环
		fmt.Println(i)
	}
}
```

从上面的例子中我们看到有两种方式在接收值的时候判断该通道是否被关闭，不过我们通常使用的是`for range`的方式。使用`for range`遍历通道，当通道被关闭的时候就会退出`for range`。

##### 单向通道

有的时候我们会将通道作为参数在多个任务函数间传递，很多时候我们在不同的任务函数中使用通道都会对其进行限制，比如限制通道在函数中只能发送或只能接收。

Go语言中提供了**单向通道**来处理这种情况。例如，我们把上面的例子改造如下：

```go
func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}
func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go squarer(ch2, ch1)
	printer(ch2)
}
```

其中，

- chan<- int是一个只写单向通道（只能对其写入int类型值），可以对其执行发送操作但是不能执行接收操作；
- <-chan int是一个只读单向通道（只能从其读取int类型值），可以对其执行接收操作但是不能执行发送操作。

在函数传参及任何赋值操作中可以将双向通道转换为单向通道，但反过来是不可以的。

##### 通道总结

`channel`常见的异常总结，如下图：![channel异常总结](https://www.liwenzhou.com/images/Go/concurrence/channel01.png)

关闭已经关闭的`channel`也会引发`panic`。

#### worker pool（goroutine池）

在工作中我们通常会使用可以指定启动的goroutine数量–`worker pool`模式，控制`goroutine`的数量，防止`goroutine`泄漏和暴涨。

一个简易的`work pool`示例代码如下：

```go
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	// 开启3个goroutine
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	// 5个任务
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	// 输出结果
	for a := 1; a <= 5; a++ {
		<-results
	}
}
```

#### select多路复用

在某些场景下我们需要同时从多个通道接收数据。通道在接收数据时，如果没有数据可以接收将会发生阻塞。你也许会写出如下代码使用遍历的方式来实现：

```go
for{
    // 尝试从ch1接收值
    data, ok := <-ch1
    // 尝试从ch2接收值
    data, ok := <-ch2
    …
}
```

这种方式虽然可以实现从多个通道接收值的需求，但是运行性能会差很多。为了应对这种场景，Go内置了`select`关键字，可以同时响应多个通道的操作。

`select`的使用类似于switch语句，它有一系列case分支和一个默认的分支。<font color='red'>每个case会对应一个通道的通信（接收或发送）过程。`select`会一直等待，直到某个`case`的通信操作完成时，就会执行`case`分支对应的语句。</font>具体格式如下：

```go
select{
    case <-ch1:
        ...
    case data := <-ch2:
        ...
    case ch3<-data:
        ...
    default:
        默认操作
}
```

举个小例子来演示下`select`的使用：

```go
func main() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}
```

使用`select`语句能提高代码的可读性。

- 可处理一个或多个channel的发送/接收操作。
- 如果多个`case`同时满足，`select`会随机选择一个。
- 如果所有 `case` 均阻塞，且定义了 `default` 模块，则执行 `default` 模块。若未定义 `default` 模块，则 select 语句阻塞，直到有 `case` 被唤醒。
- 对于没有`case`的`select{}`会一直等待，可用于阻塞main函数。

```go
func main() {
    intch := make(chan int, 1)
    strch := make(chan string, 1)
    intch <- 1
    strch <- "Hello"
    select {
    case value := <-intch:
        fmt.Println(value)
    case value := <-strch:
        fmt.Println(value)    
    case <-time.After(time.Second * 5)
        fmt.Println("超時")
    }
}
```



#### 并发安全和锁

有时候在Go代码中可能会存在多个`goroutine`同时操作一个资源（临界区），这种情况会发生`竞态问题`（数据竞态）。类比现实生活中的例子有十字路口被各个方向的的汽车竞争；还有火车上的卫生间被车厢里的人竞争。

举个例子：

```go
var x int64
var wg sync.WaitGroup

func add() {
	for i := 0; i < 5000; i++ {
		x = x + 1
	}
	wg.Done()
}
func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
```

上面的代码中我们开启了两个`goroutine`去累加变量x的值，这两个`goroutine`在访问和修改`x`变量的时候就会存在数据竞争，导致最后的结果与期待的不符。

##### 互斥锁

互斥锁是一种常用的控制共享资源访问的方法，它能够保证同时只有一个goroutine可以访问共享资源。Go语言中使用sync包的Mutex类型来实现互斥锁。 使用互斥锁来修复上面代码的问题：

```go
var x int64
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock() // 加锁
		x = x + 1
		lock.Unlock() // 解锁
	}
	wg.Done()
}
func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
```

使用互斥锁能够保证同一时间有且只有一个`goroutine`进入临界区，其他的`goroutine`则在等待锁；当互斥锁释放后，等待的`goroutine`才可以获取锁进入临界区，多个`goroutine`同时等待一个锁时，唤醒的策略是随机的。

##### 读写互斥锁

互斥锁是完全互斥的，但是有很多实际的场景下是读多写少的，当我们并发的去读取一个资源不涉及资源修改的时候是没有必要加锁的，这种场景下使用读写锁是更好的一种选择。读写锁在Go语言中使用`sync`包中的`RWMutex`类型。

读写锁分为两种：读锁和写锁。当一个goroutine获取读锁之后，其他的`goroutine`如果是获取读锁会继续获得锁，如果是获取写锁就会等待；当一个`goroutine`获取写锁之后，其他的`goroutine`无论是获取读锁还是写锁都会等待。

```go
var (
	x      int64
	wg     sync.WaitGroup
	lock   sync.Mutex
	rwlock sync.RWMutex
)

func write() {
	// lock.Lock()   // 加互斥锁
	rwlock.Lock() // 加写锁
	x = x + 1
	time.Sleep(10 * time.Millisecond) // 假设写操作耗时10毫秒
	rwlock.Unlock()                   // 解写锁
	// lock.Unlock()                     // 解互斥锁
	wg.Done()
}

func read() {
	// lock.Lock()                  // 加互斥锁
	rwlock.RLock()               // 加读锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	rwlock.RUnlock()             // 解读锁
	// lock.Unlock()                // 解互斥锁
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}
```

需要注意的是读写锁非常适合读多写少的场景，如果读和写的操作差别不大，读写锁的优势就发挥不出来。

#### sync.WaitGroup

在代码中生硬的使用`time.Sleep`肯定是不合适的，Go语言中可以使用`sync.WaitGroup`来实现并发任务的同步。 `sync.WaitGroup`有以下几个方法：

|             方法名              |        功能         |
| :-----------------------------: | :-----------------: |
| (wg * WaitGroup) Add(delta int) |    计数器+delta     |
|     (wg *WaitGroup) Done()      |      计数器-1       |
|     (wg *WaitGroup) Wait()      | 阻塞直到计数器变为0 |

`sync.WaitGroup`内部维护着一个计数器，计数器的值可以增加和减少。例如当我们启动了N 个并发任务时，就将计数器值增加N。每个任务完成时通过调用Done()方法将计数器减1。通过调用Wait()来等待并发任务执行完，当计数器值为0时，表示所有并发任务已经完成。

我们利用`sync.WaitGroup`将上面的代码优化一下：

```go
var wg sync.WaitGroup

func hello() {
	defer wg.Done()
	fmt.Println("Hello Goroutine!")
}
func main() {
	wg.Add(1)
	go hello() // 启动另外一个goroutine去执行hello函数
	fmt.Println("main goroutine done!")
	wg.Wait()
}
```

需要注意`sync.WaitGroup`是一个结构体，传递的时候要传递指针。

#### sync.Once

说在前面的话：这是一个进阶知识点。

在编程的很多场景下我们需要确保某些操作在高并发的场景下只执行一次，例如只加载一次配置文件、只关闭一次通道等。

作用与 **init** 函数类似。但也有所不同。

- **init** 函数是在文件包首次被加载的时候执行，且只执行一次
- **sync.Once** 是在代码运行中需要的时候执行，且只执行一次

当一个函数不希望程序在一开始的时候就被执行的时候，我们可以使用 **sync.Once** 。其最常应用于单例模式之下，例如初始化系统配置、保持数据库唯一连接等。

Go语言中的`sync`包中提供了一个针对只执行一次场景的解决方案–`sync.Once`。

`sync.Once`只有一个`Do`方法，其签名如下：

```go
func (o *Once) Do(f func()) {}
```

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}

# Output:
Only once
```

*备注：如果要执行的函数`f`需要传递参数就需要搭配闭包来使用。*

##### 代码实现

**sync.Once** 使用变量 *done* 来记录函数的执行状态，使用 *sync.Mutex* 和 *sync.atomic* 来保证线程安全的读取 *done* 。Once 结构体非常简单，其中 done 是调用标识符，Once 对象初始化时，其 done 值默认为 0，Once 仅有一个 Do() 方法，当 Once 首次调用 Do() 方法后，done 值变为 1。m 作用于初始化竞态控制，在第一次调用 Once.Do() 方法时，会通过 m 加锁，以保证在第一个 Do() 方法中的参数 f() 函数还未执行完毕时，其他此时调用 Do() 方法会被阻塞（不返回也不执行）。

```go
package sync

import (
	"sync/atomic"
)

// Once is an object that will perform exactly one action.
type Once struct {
	m    Mutex
	done uint32
}

// Do calls the function f if and only if Do is being called for the
// first time for this instance of Once. In other words, given
// 	var once Once
// if once.Do(f) is called multiple times, only the first call will invoke f,
// even if f has a different value in each invocation. A new instance of
// Once is required for each function to execute.
//
// Do is intended for initialization that must be run exactly once. Since f
// is niladic, it may be necessary to use a function literal to capture the
// arguments to a function to be invoked by Do:
// 	config.once.Do(func() { config.init(filename) })
//
// Because no call to Do returns until the one call to f returns, if f causes
// Do to be called, it will deadlock.
//
// If f panics, Do considers it to have returned; future calls of Do return
// without calling f.
//
func (o *Once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 1 {
		return
	}
	// Slow-path.
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}
```

Do() 方法的入参是一个无参数输入与返回的函数，当 o.done 值为 0 时，执行 doSlow() 方法，为1则退出 Do() 方法。doSlow() 方法很简单：加锁，再次检查 o.done 值，执行 f()，原子操作将  o.done 值置为1，最后释放锁。

##### 加载配置文件示例

延迟一个开销很大的初始化操作到真正用到它的时候再执行是一个很好的实践。因为预先初始化一个变量（比如在 init 函数中完成初始化）会增加程序的启动耗时，而且有可能实际执行过程中这个变量没有用上，那么这个初始化操作就不是必须要做的。我们来看一个例子：

```go
var icons map[string]image.Image

func loadIcons() {
	icons = map[string]image.Image{
		"left":  loadIcon("left.png"),
		"up":    loadIcon("up.png"),
		"right": loadIcon("right.png"),
		"down":  loadIcon("down.png"),
	}
}

// Icon 被多个goroutine调用时不是并发安全的
func Icon(name string) image.Image {
	if icons == nil {
		loadIcons()
	}
	return icons[name]
}
```

多个 `goroutine` 并发调用 Icon 函数时不是并发安全的，现代的编译器和 CPU 可能会在保证每个`goroutine`都满足串行一致的基础上自由地重排访问内存的顺序。loadIcons 函数可能会被重排为以下结果：

```go
func loadIcons() {
	icons = make(map[string]image.Image)
	icons["left"] = loadIcon("left.png")
	icons["up"] = loadIcon("up.png")
	icons["right"] = loadIcon("right.png")
	icons["down"] = loadIcon("down.png")
}
```

在这种情况下就会出现即使判断了`icons`不是nil也不意味着变量初始化完成了。考虑到这种情况，我们能想到的办法就是添加互斥锁，保证初始化`icons`的时候不会被其他的`goroutine`操作，但是这样做又会引发性能问题。

使用`sync.Once`改造的示例代码如下：

```go
var icons map[string]image.Image

var loadIconsOnce sync.Once

func loadIcons() {
	icons = map[string]image.Image{
		"left":  loadIcon("left.png"),
		"up":    loadIcon("up.png"),
		"right": loadIcon("right.png"),
		"down":  loadIcon("down.png"),
	}
}

// Icon 是并发安全的
func Icon(name string) image.Image {
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}
```

#####  并发安全的单例

```go
package singleton

import "sync"

type singleton struct {}

var once sync.Once
var instance *singleton

func getInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

```



#### sync.Pool

Pool 就是对象缓存池，用来减少堆上内存的反复申请和释放的。因为 golang 的内存是用户触发申请，runtime 负责回收。如果用户申请内存过于频繁，会导致 runtime 的回收压力陡增，从而影响整体性能。

有了 sync.Pool 之后就不一样了，对象申请先看池子里有没有现成的，有就直接返回。释放的时候内存也不是直接归还，而是放进池子而已，适时释放。这样就能极大的减少申请内存的频率，从而减少 gc 压力。

sync.Pool 是 golang 提供的对象重用的机制，sync.Pool 是可伸缩的，并发安全的。其大小仅受限于内存的大小，可以被看作是一个存放可重用对象的值的容器。设计的目的是存放已经分配的但是暂时不用的对象，在需要用到的时候直接从 pool 中取。

##### **初始化 Pool 实例 New**

第一个步骤就是创建一个 Pool 实例，关键一点是配置 New 属性，声明 Pool 元素创建的方法。

```GO
bufferPool := &sync.Pool {
    New: func() interface {} {
        fmt.Println("Create new instance")
        return struct{}{}
    }
}
```

通过`New`去定义你这个池子里面放的究竟是什么东西，在这个池子里面你只能放一种类型的东西。

##### 申请对象 Get

```go
buffer := bufferPool.Get()
```

`Get` 方法会返回 Pool 已经存在的对象，如果没有，那么就走慢路径，也就是调用初始化的时候定义的 New 方法（也就是最开始定义的初始化行为）来初始化一个对象。如果没有对 New 进行赋值，则返回 nil。

##### 释放对象 Put

```go
bufferPool.Put(buffer)
```

使用对象之后，调用 Put 方法把对象放回池子。注意了，这个调用之后仅仅把这个对象放回池子，池子里面的对象啥时候真正释放外界是不清楚，是不受外部控制的。

```go
// 一个[]byte的对象池，每个对象为一个[]byte
var bytePool = sync.Pool{
  New: func() interface{} {
    b := make([]byte, 512)
    return &b
  },
}
 
func main() {
  a := time.Now().Unix()
  // 不使用对象池
  for i := 0; i < 1000000000; i++{
    obj := make([]byte,512)
    _ = obj
  }
  b := time.Now().Unix()
  // 使用对象池
  for i := 0; i < 1000000000; i++{
    obj := bytePool.Get().(*[]byte)
    _ = obj
    bytePool.Put(obj)
  }
  c := time.Now().Unix()
  fmt.Println("without pool ", b - a, "s")
  fmt.Println("with    pool ", c - b, "s")
}
 
//without pool  0 s
//with    pool  26 s
```

貌似使用池化，性能弱爆了？？？这似乎与net/http使用sync.pool池化Request来优化性能的选择相违背。这同时也说明了一个问题，好的东西，如果滥用反而造成了性能成倍的下降。在看过pool原理之后，结合实例，将给出正确的使用方法，并给出预期的效果

##### sync.Pool 的实现

为了使多个 goroutine 操作同一个 pool 做到高效，sync.Pool 为每一个 p 都分配了一个子池。当执行 get 或者put 操作时，会对当前 goroutine 挂载的子池操作。每个子池都有一个私有对象和共享列表对象，私有对象只有对应的 p 能够访问，因为同一个 p 同一时间只能操作执行一个 goroutine，因此对私有对象的操作不需要加锁；但共享列表是和其他 P 分享的，因此操作是需要加锁的。

**获取对象的过程：**

- 固定某个 P，尝试从私有对象中获取， 如果是私有对象则返回该对象，并把私有对象赋空。
- 如果私有对象是空的，需要加锁，从当前固定的 p 的共享池中获取-并从该共享队列中删除这个对象。

- 如果当前的子池都是空的，尝试去其他 P 的子池的共享列表偷取一个，如果用户没有注册New函数则返回nil。

**归还对象的过程：**

- 固定到某个 p，如果私有对象为空则放到私有对象。
- 如果私有对象不为空，加锁，加入到该 P 子池的共享列表中。
  

##### 对象池的详细实现

**对象池结构**

```go
type Pool struct {
	noCopy noCopy            //防止copy
 
	local     unsafe.Pointer //本地p缓存池指针
	localSize uintptr        //本地p缓存池大小
 
	//当池中没有对象时，会调用New函数调用一个对象
	New func() interface{}
}

type poolLocal struct {  
    private interface{}   // Can be used only by the respective P.  
    shared  []interface{} // Can be used by any P.  
    Mutex                 // Protects shared.  
    pad     [128]byte     // Prevents false sharing.  
}
```

local 成员的真实类型是一个 poolLocal 数组，localSize 是数组长度。这涉及到 Pool

**获取对象池中的对象**

```go
func (p *Pool) Get() interface{} {
	if race.Enabled {
		race.Disable()
	}
        //获取本地的poolLocal对象
	l := p.pin()
 
        //先获取private池中的私有变量
	x := l.private
	l.private = nil
	runtime_procUnpin()
	if x == nil {
                //查找本地的共享池，因为本地的共享池可能被其他p访问，所以要加锁
		l.Lock()
		last := len(l.shared) - 1
		if last >= 0 {
                        //如果本地共享池有对象，取走最后一个
			x = l.shared[last]
			l.shared = l.shared[:last]
		}
		l.Unlock()
                //查找其他p的共享池
		if x == nil {
			x = p.getSlow()
		}
	}
	if race.Enabled {
		race.Enable()
		if x != nil {
			race.Acquire(poolRaceAddr(x))
		}
	}
        //未找到其他可用元素，则调用New生成
	if x == nil && p.New != nil {
		x = p.New()
	}
	return x
}
```

![img](https://img-blog.csdnimg.cn/20191219144553495.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3FxXzM1NzAzODQ4,size_16,color_FFFFFF,t_70)

**归还对象池中的对象**

```go
func (p *Pool) Put(x interface{}) {
	if x == nil {
		return
	}
	if race.Enabled {
		if fastrand()%4 == 0 {
			//1/4的概率会把该元素扔掉
			return
		}
		race.ReleaseMerge(poolRaceAddr(x))
		race.Disable()
	}
	l := p.pin()
	if l.private == nil {
                //赋值给私有变量
		l.private = x
		x = nil
	}
	runtime_procUnpin()
	if x != nil {
                //访问共享池加锁
		l.Lock()
		l.shared = append(l.shared, x)
		l.Unlock()
	}
	if race.Enabled {
		race.Enable()
	}
}
```

![img](F:\Users\WSQ\Go\src\Go-\github.com\Maker-Wu\studygo\day05\05_goroutine\note.assets\aHR0cDovL3R0Yy10YWwub3NzLWNuLWJlaWppbmcuYWxpeXVuY3MuY29tLzE1NzY2NzYyOTgvJUU2JTlDJUFBJUU1JTkxJUJEJUU1JTkwJThEJUU4JUExJUE4JUU1JThEJTk1JTIwJTI4NCUyOS5wbmc.png)

#### sync.Cond

sync.Cond 实现了一个条件变量，用于等待一个或一组 goroutines 满足条件后唤醒的场景。每个 Cond 关联一个 Locker 通常是一个 *Mutex 或 *RWMutex 根据需求初始化不同的锁。

条件等待和互斥锁有不同，互斥锁是不同协程共用一个锁，条件等待是不同协程各用一个锁，但是 wait() 方法会等待（阻塞）,直到有信号发过来

使用场景：

我需要完成一项任务，但是这项任务需要满足一定条件才可以执行，否则我就等着。
那我可以怎么获取这个条件呢？一种是循环去获取，一种是条件满足的时候通知我就可以了。显然第二种效率高很多。
通知的方式的话，golang 里面通知可以用 channe l的方式

```go
var mail = make(chan string)
    go func() {
        <- mail
        fmt.Println("get chance to do something")
    }()
    time.Sleep(5*time.Second)
    mail <- "moximoxi"
```

但是 channel 的方式还是比较适用于一对一的方式，一对多并不是很适合。下面就来介绍一下另一种方式：sync.Cond
sync.Cond 就是用于实现条件变量的，是基于 sync.Mutex 的基础上，增加了一个通知队列，通知的线程会从通知队列中唤醒一个或多个被通知的线程。
主要有以下几个方法：

```go
sync.NewCond(&mutex)：生成一个cond，需要传入一个mutex，因为阻塞等待通知的操作以及通知解除阻塞的操作就是基于sync.Mutex来实现的。
sync.Wait()：用于等待通知
sync.Signal()：用于发送单个通知
sync.Broadcat()：用于广播
```

看到上面几个方法的源码

```go
var locker sync.Mutex
var cond = sync.NewCond(&locker)
// NewCond(l Locker)里面定义的是一个接口,拥有lock和unlock方法。
// 看到sync.Mutex的方法,func (m *Mutex) Lock(),可以看到是指针有这两个方法,所以应该传递的是指针
func main() {
    for i := 0; i < 10; i++ {
        go func(x int) {
            cond.L.Lock()         // 获取锁
            defer cond.L.Unlock() // 释放锁
            cond.Wait()           // 等待通知，阻塞当前 goroutine
            // 通知到来的时候, cond.Wait()就会结束阻塞, do something. 这里仅打印
            fmt.Println(x)
        }(i)
    }
    time.Sleep(time.Second * 1) // 睡眠 1 秒，等待所有 goroutine 进入 Wait 阻塞状态
    fmt.Println("Signal...")
    cond.Signal()               // 1 秒后下发一个通知给已经获取锁的 goroutine
    time.Sleep(time.Second * 1)
    fmt.Println("Signal...")
    cond.Signal()               // 1 秒后下发下一个通知给已经获取锁的 goroutine
    time.Sleep(time.Second * 1)
    cond.Broadcast()            // 1 秒后下发广播给所有等待的goroutine
    fmt.Println("Broadcast...")
    time.Sleep(time.Second * 1) // 睡眠 1 秒，等待所有 goroutine 执行完毕

}
```

上述代码实现了主线程对多个goroutine的通知的功能。
抛出一个问题：
主线程执行的时候，如果并不想触发所有的协程，想让不同的协程可以有自己的触发条件，应该怎么使用？
下面就是一个具体的需求：
有三个 worker 和一个 master，worker 等待 master 去分配指令，master 一直在计数，计数到 5 的时候通知第一个 worker，计数到 10 的时候通知第二个和第三个 worker 。
首先列出几种解决方式
1、所有worker循环去查看master的计数值，计数值满足自己条件的时候，触发操作 >>>>>>>>>弊端：无谓的消耗资源
2、用 channel 来实现，几个 worker 几个 channel，eg: worker1 的协程里 <-channel(worker1) 进行阻塞，计数值到 5 的时候，给 worker1 的 channel 放入值，阻塞解除，worker1 开始工作。 >>>>>>>>弊端：channel 还是比较适用于一对一的场景，一对多的时候，需要起很多的 channel ，不是很美观
3、用条件变量sync.Cond，针对多个worker的话，用broadcast，就会通知到所有的worker

```go
func main()  {
    mutex := sync.Mutex{}
    var cond = sync.NewCond(&mutex)
    mail := 1
    go func() {
        for count := 0; count <= 15; count++ {
            time.Sleep(time.Second)
            mail = count
            cond.Broadcast()
        }
    }()
    // worker1
    go func() {
        for mail != 5 {          // 触发的条件，如果不等于5，就会进入cond.Wait()等待，此时cond.Broadcast()通知进来的时候，wait阻塞解除，进入下一个循环，此时发现mail != 5，跳出循环，开始工作。
            cond.L.Lock()
            cond.Wait()
            cond.L.Unlock()
        }
        fmt.Println("worker1 started to work")
        time.Sleep(3*time.Second)
        fmt.Println("worker1 work end")
    }()
    // worker2
    go func() {
        for mail != 10 {
            cond.L.Lock()
            cond.Wait()
            cond.L.Unlock()
        }
        fmt.Println("worker2 started to work")
        time.Sleep(3*time.Second)
        fmt.Println("worker2 work end")
    }()
    // worker3
    go func() {
        for mail != 10 {
            cond.L.Lock()
            cond.Wait()
            cond.L.Unlock()
        }
        fmt.Println("worker3 started to work")
        time.Sleep(3*time.Second)
        fmt.Println("worker3 work end")
    }()
    select {

    }
}
为什么每个worker里要使用for循环？而不是用if？
首先broadcast的时候，会通知到所有的worker，此时wait都会解除，但并不是所有的worker都满足通知条件的，所以加一个for循环，不满足通知条件的会再次wait。
```

#### sync.Map

Go语言中内置的map不是并发安全的。

sync.Map 有以下特性：

- 无须初始化，直接声明即可。
- sync.Map 不能使用 map 的方式进行取值和设置等操作，而是使用 sync.Map 的方法进行调用，Store表示存储，Load表示获取，Delete 表示删除。
- 使用 Range 配合一个回调函数进行遍历操作，通过回调函数返回内部遍历出来的值， Range参数中回调函数的返回值在需要继续迭代遍历时，返回 true，终止迭代遍历时，返回 false

请看下面的示例：

```go
var m = make(map[string]int)

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			set(key, n)
			fmt.Printf("k=:%v,v:=%v\n", key, get(key))
			wg.Done()
		}(i)
	}
	wg.Wait()
}
```

上面的代码开启少量几个 `goroutine` 的时候可能没什么问题，当并发多了之后执行上面的代码就会报`fatal error: concurrent map writes`错误。

像这种场景下就需要为 map 加锁来保证并发的安全性了，Go 语言的 `sync` 包中提供了一个开箱即用的并发安全版 map–`sync.Map`。开箱即用表示不用像内置的map一样使用make函数初始化就能直接使用。同时`sync.Map`内置了诸如`Store`、`Load`、`LoadOrStore`、`Delete`、`Range`等操作方法。

```go
var m = sync.Map{}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)
			value, _ := m.Load(key)
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
```

##### 迭代遍历

简单的使用例子：

```go

```

