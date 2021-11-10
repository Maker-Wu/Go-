## 使用 GOPATH 问题

1. 代码开发必须在 GOPATH 目录下，不然，就有问题
2. 依赖手动管理
3. 依赖包没有版本可言

## Vender

1. 解决了包依赖，一个配置文件就管理
2. 依赖包全都下载到项目 vendor 下，每个项目都把有一份。拉取项目时，开始怀疑人生

## Go mod 介绍

Go mod 模块是相关 Go 包的集合。modules 是源代码交换和版本控制的单元。go 命令直接支持使用 modules，包括记录和解析对其他模块的依赖性。modules 替换旧的基于 GOPATH 的方法来指定在给定构建中使用哪些源文件。 

### 设置 GO111MODULE 

- off 

  go 命令行将不会支持 module 功能，寻找依赖包的方式将会沿用旧版本那种通过 vendor 目录或者 GOPATH 模式来查找。 

- on 
  go 命令行会使用 modules，而一点也不会去 GOPATH 目录下查找。 
  
- auto 

  默认值，go 命令行将会根据当前目录下是否有 go.mod 文件来决定是否启用 module 功能 
  当 modules 功能启用时，依赖包的存放位置变更为 $GOPATH/pkg，允许同一个package 多个版本并存，且多个项目可以共享缓存的 module。

注意：go env -w 会将配置写到`GOENV="/Users/WSQ/Library/Application Support/go/env"`

### 使用Go mod 命令管理包

| 命令      | 描述                           |
| :-------- | :----------------------------- |
| go init   | 在当前目录项目下初始化 mod     |
| go tidy   | 拉取依赖的模块，移除不用的模块 |
| go vendor | 将依赖复制到 vendor 下         |
| go edit   | 编辑 go.mod                    |
| go verify | 验证依赖是否正确               |

其实工作基本上都使用 init 和 tidy 就够了。

#### gomod 文件的创建

**命令：go mod init 项目名**

比如，我的项目是 manage，那么就可以这样使用：

```shell
go mod init manage
```

此时就会在当前项目下生成 gomod 文件：

![image-20211109130625409](F:\Users\WSQ\Go\src\Go-\github.com\Maker-Wu\studygo\day07\01_gomod\note.assets\image-20211109130625409.png)

注意，如果当前的项目是要给外部使用的，最好是配合 git 仓库命名，比如

```go
go mod init github.com/Maker-Wu/manage
```

以便其他项目可以 go get 引用得到

#### 依赖包的版本生成

上面的命令只是生成一个 gomod 文件，至于依赖包的版本号信息暂时是还没有生成，可以使用下面两个目录进行获取：

```go
命令1： go get 包名 
```

如果依赖包比较多，那么 go get 就比较麻烦了。可以使用另外一个命令：

```GO
命令2： go mod tidy
```

这个命令将会扫描所有我们 import 到的包，并生成对应的记录到 gomod 文件里。

![image-20211110124319096](F:\Users\WSQ\Go\src\Go-\github.com\Maker-Wu\studygo\day07\01_gomod\note.assets\image-20211110124319096.png)

此时，我们看到了一条**依赖包以及版本号**的信息记录。

这里的 v2.1.0 是因为引用的 go-cache 包在 github 上已经打标签了，所以有 v2.1.0 类似的字样出现，后面有 incompatible 是因为 go-cache 包的命名没有遵循官方规范，所以加了 incompatible 加以区分。

如果引用的包没有打过标签，那就有其他的版本记录生成规则, 比如

v0.0.0-20210501091049-10806f459f65

就表示**版本号 + 日期 + 主分支**最新的 commit **哈希值**前缀。

此外，我们还发现除了 gomod 文件之外，还有 gosum 文件。此文件主要是用来记录依赖包的 hash 值，防止部署到新环境时，重新拉取的包与之前本地拉取的包不一致。

#### gomod 文件的使用技巧

##### 1）引用分支的包

默认情况下，go mod tidy 会拉取**主分支**的最新代码作为版本记录。

如果我们有多个功能在同时开发，按常规操作，是需要新建各自的 feature 分支来开发的，而不会在主分支上直接开发的。

所以，当有其他模块需要引用分支代码时，我们就不能按常规操作 go mod tidy 了。

此时我们需要**手动修改** gomod 引用包的版本名字，替换为对应的**分支名**，比如 gomod 文件改为:

```go
require github.com/patrickmn/go-cache develop
```

##### 2）引用本地开发的代码

golang 是根据 gomod 文件来构建程序的，如果我们引用了其他项目代码，那每次就得先提交代码到 git 仓库，然后重新构建 gomod 文件才能引用到最新的代码。

为了能直接**引用本地**正在开发的包，又不频繁提交代码，我们可以使用下面这个命令

```go
replace github.com/patrickmn/go-cache => 本地项目包的地址
```

这样就可以在构建项目时，本地联调了。等联调完毕，再一次性的提交代码到 git 仓库里。

##### 3）查看依赖包的历史版本

使用 go mod tidy 命令时总会拉取最新版本的依赖包，但当我们只想 import 某个**历史版本**时，就可以使用下面的命令来获取历史版本号了:

```go
go list -m -versions github.com/patrickmn/go-cache
```

执行结果：

```go
github.com/patrickmn/go-cache v1.0.0 v2.0.0+incompatible v2.1.0+incompatible
```

然后当我们想引用 v1.0.0 时，就可以这样改写了:

```go
require github.com/patrickmn/go-cache v1.0.0
```

需要注意的是，`go list`某个包，必须得先在此项目中引用了对应的包。例如，当前项目如果没有引用 go-cache 包，则`go list`是获取不到信息的。

#### go module 其他命令

go list -m all ：列出当前项目包名以及所有依赖到的包

go mod vendor: 将引用的包都生成到当前项目的 vendor 包下，这样可以不用每次重新构建时去拉取对应的包，直接加入到自己的 git 代码仓库管理中, 直接 git pull 即可。

另外，有点要注意的就是，如果我们在 gomod 文件里手动添加了某个依赖包，但实际在项目里并没有使用到这个依赖包时，那么在执行 go mod tidy 构建时，就会自动删除这个依赖包的相关记录。

### 总结

go mod 的使用很简单，`go mod init`、`go mod tidy`，基本就能解决很多依赖问题了。这也是 Go 官方一直提倡的简洁、优雅。