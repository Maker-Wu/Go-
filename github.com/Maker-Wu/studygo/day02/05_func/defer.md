Go语言中的`defer`语句会将其后面跟随的语句进行延迟处理。在`defer`归属的函数即将返回时，将延迟处理的语句按`defer`定义的逆序进行执行，也就是说，先被`defer`的语句最后被执行，最后被`defer`的语句，最先被执行。

#### defer的执行顺序

example1

```go
func f() (result int) {
    defer func() {
        result++
    }()
    return 0
}
```

example2

```go
func f() (r int) {
     t := 5
     defer func() {
       t = t + 5
     }()
     return t
}
```

example3

```go
func f() (r int) {
    defer func(r int) {
          r = r + 5
    }(r)
    return 1
}
```

首先明确的是：defer是在return之前执行的。

然后是了解defer的实现方式，大致就是在defer出现的地方，插入指令CALL runtime.deferproc，然后在函数返回之前的地方，插入指令CALL runtime.deferreturn。再就是明确go返回值的方式跟C是不一样的，为了支持多值返回，go是用栈返回值的，而C是用寄存器。

最最重要的一点就是：**return xxx这一条语句并不是一条原子指令!**

整个return过程，没有defer之前，先在栈中写一个值，这个值会被当作返回值，然后再调用RET指令返回。return xxx语句汇编后是先给返回值赋值，再做一个空的return，( 赋值指令 ＋ RET指令)。defer的执行是被插入到return指令之前的，有了defer之后，就变成了(赋值指令 + CALL defer指令 + RET指令)。<font color='red'>而在CALL defer函数中，有可能将最终的返回值改写了...也有可能没改写。</font>总之，如果改写了，那么看上去就像defer是在return xxx之后执行的。

![defer执行时机](defer.assets/defer.png)

这是所有你所想不明白的defer故事发生的根源。

上面的基础知识都有了，然后就可以说说神奇的defer了。告诉大家一个简单的转换规则大家就再也不为defer迷糊了。改写规则是将return语句分开成两句写，return xxx会被改写成:

```go
返回值 = xxx
调用defer函数
空的return
```

先看example1，它可以改写成这样：

```go
func f() (result int) {
     result = 0  //return语句不是一条原子调用，return xxx其实是赋值＋RET指令
     func() { //defer被插入到return之前执行，也就是赋返回值和RET指令之间
         result++
     }()
     return
}
```

所以这个返回值是1。

再看example2，它可以改写成这样：

```go
func f() (r int) {
     t := 5
     r = t //赋值指令
     func() {        //defer被插入到赋值与返回之间执行，这个例子中返回值r没被修改过
         t = t + 5
     }
     return        //空的return指令
}
```

所以这个的结果是5。

最后看example3，它改写后变成：

```go
func f() (r int) {
     r = 1  //给返回值赋值
     func(r int) {        //这里改的r是传值传进去的r，不会改变要返回的那个r值
          r = r + 5
     }(r)
     return        //空的return
}
```

懂了么？ 结论：defer确实是在return之前调用的。但表现形式上却可能不像。本质原因是return xxx语句并不是一条原子指令，defer被插入到了赋值 与 RET之前，因此可能有机会改变最终的返回值。当你觉得迷糊时，可以用我给的这套规则转一下代码。

