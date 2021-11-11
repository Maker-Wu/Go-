## 使用 GOPATH 问题

1. 代码开发必须在 GOPATH 目录下，不然，就有问题
2. 依赖手动管理
3. 依赖包没有版本可言

## Vender

1. 解决了包依赖，一个配置文件就管理
2. 依赖包全都下载到项目 vendor 下，每个项目都把有一份。拉取项目时，开始怀疑人生

## Go mod 介绍

Go mod 模块是相关 Go 包的集合。<font color='red'>modules 是源代码交换和版本控制的单元。</font>go 命令直接支持使用 modules，包括记录和解析对其他模块的依赖性。modules 替换旧的基于 GOPATH 的方法来指定在给定构建中     使用哪些源文件。

go.mod 这个文件里记录了当前项目里所有依赖包的 git 仓库地址以及对应的**版本号**，来解决了包依赖管理的问题，后续在构建编译时，就会根据对应的版本号去拉取依赖包。 

### 设置 GO111MODULE 

- off 

  go 命令行将不会支持 module 功能，寻找依赖包的方式将会沿用旧版本那种通过 vendor 目录或者 GOPATH 模式来查找。 

- on 
  go 命令行会使用 modules，而一点也不会去 GOPATH 目录下查找。 
  
- auto 

  默认值，go 命令行将会根据当前目录下是否有 go.mod 文件来决定是否启用 module 功能 
  当 modules 功能启用时，依赖包的存放位置变更为 $GOPATH/pkg，允许同一个package 多个版本并存，且多个项目可以共享缓存的 module。

注意：go env -w 在 windows 系统会将配置写到`GOENV="C:\Users\God Wu\AppData\Roaming\go\env"`

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

![image-20211109130625409](.\note.assets\image-20211109130625409.png)

go.mod 文件一旦创建后，它的内容将会被 go toolchain 全面掌控。go toolchain 会在各类命令执行时，比如 go get、go build、go mod 等修改和维护 go.mod 文件。

go.mod 提供了 module, require、replace 和 exclude 四个命令

- module 语句指定包的名字（路径）
- require 语句指定包的依赖项模块

- replace 语句可以替换依赖项模块
- exclude 语句可以忽略依赖项模块

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

![image-20211110124319096](.\note.assets\image-20211110124319096.png)

此时，我们看到了一条**依赖包以及版本号**的信息记录。

<font color='red'>这里的 v2.1.0 是因为引用的 go-cache 包在 github 上已经打标签了</font>，所以有 v2.1.0 类似的字样出现，后面有 incompatible 是因为 go-cache 包的命名没有遵循官方规范，所以加了 incompatible 加以区分。

如果引用的包没有打过标签，那就有其他的版本记录生成规则, 比如

v0.0.0-20210501091049-10806f459f65

就表示**版本号 + 日期 + 主分支**最新的 commit **哈希值**前缀。

此外，我们还发现除了 gomod 文件之外，还有 gosum 文件。此文件主要是用来记录依赖包的 hash 值，防止部署到新环境时，重新拉取的包与之前本地拉取的包不一致。

##### 实战示例

创建 main.go 文件

```go
package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
```

执行 go run main.go 运行代码会发现 go mod 会自动查找依赖自动下载 再查看 go.mod

```go
module Gone

go 1.14

require github.com/gin-gonic/gin v1.6.3
```

但有可能结果会报错

```go
main.go:3:8: no required module provides package github.com/gin-gonic/gin; to add it:
        go get github.com/gin-gonic/gin
```

##### 解决方法

执行:go mod edit -require github.com/gin-gonic/gin@latest 解决，指定Gin的版本
再次运行 go run main.go 报错

```go
go: updates to go.mod needed; to update it:
        go mod tidy
```

然后执行 go mod tidy 再次执行 go run main.go 终于跑起来了

#### go get 升级

- 运行 go get -u 将会升级到最新的次要版本或者修订版本 (x.y.z, z 是修订版本号，y 是次要版本号）
- 运行 go get -u=patch 将会升级到最新的修订版本

- 运行 go get package@version 将会升级到指定的版本号 version
- 运行 go get 如果有版本的更改，那么 go.mod 文件也会更改

#### 使用 replace 替换无法直接获取的 package

由于某些已知的原因，并不是所有的 package 都能成功下载，比如：golang.org 下的包。modules 可以通过在 go.mod 文件中使用 replace 指令替换成 github 上对应的库，比如：

```go
replace (
    golang.org/x/crypto v0.0.0-20190313024323-a1f597ede03a => github.com/golang/crypto v0.0.0-20190313024323-a1f597ede03a
)
```

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





 

### go mod 发布和使用

#### Creating a Module 

如果你设置好 go mod 了，那你就可以在任何目录下随便创建

```go
$mkdir gomodone
$cd gomodone
```

在这个目录下创建一个文件say.go

```go
package gomodone

import "fmt" 

// say Hi to someone
func SayHi(name string) string {
   return fmt.Sprintf("Hi, %s", name)
}
```

初始化一个 go.mod文件

```go
$ go mod init github.com/jacksonyoudi/gomodone
go: creating new go.mod: module github.com/jacksonyoudi/gomodone
```

查看 go.mod 内容如下：

```go
github.com/jacksonyoudi/gomodone
go 1.14
```

下面我们要将这个 module 发布到 github 上，然后在另外一个程序使用

```go
$git init
$vim .gitiiignore
$git commit -am "init"
// github 创建对应的 repo
$git remote add origin git@github.com:jacksonyoudi/gomodone.git
$git push -u origin master
```

执行完，上面我们就相当于发布完了。

如果有人需要使用，就可以使用

```go
go get github.com/jacksonyoudi/gomodone
```

这个时候没有加 tag，所以，没有版本的控制。默认是 v0.0.0 后面接上时间和 commitid。如下：

```go
gomodone@v0.0.0-20200517004046-ee882713fd1e
```

官方不建议这样做，没有进行版本控制管理。

### module versioning

使用 tag，进行版本控制

#### making a release

```go
git tag v1.0.0
git push --tags
```

操作完，我们的 module 就发布了一个 v1.0.0 的版本了。

推荐在这个状态下，再切出一个分支，用于后续 v1.0.0 的修复推送，不要直接在 master 分支修复

```go
$git checkout -b v1
$git push -u origin v1
```

### use our module

上面已经发布了一个 v1.0.0 的版本，我们可以在另一个项目中使用，创建一个 go 的项目

```go
$mkdir Gone
$cd Gone
$vim main.go
```

```go
package main

import (
    "fmt"
    "github.com/jacksonyoudi/gomodone"
)

func main() {
    fmt.Println(gomodone.SayHi("Roberto"))
}
```

代码写好了，我们生成 go mod 文件

```go
go mod init Gone
```

上面命令执行完，会生成 go mod 文件 看下 mod 文件：

```go
module Gone

go 1.14

require (
    github.com/jacksonyoudi/gomodone v1.0.0
)
```

```go
$go mod tidy
go: finding module for package github.com/jacksonyoudi/gomodone
go: found github.com/jacksonyoudi/gomodone in github.com/jacksonyoudi/gomodone v1.0.0
```

同时还生成了 go.sum, 其中包含软件包的哈希值，以确保我们具有正确的版本和文件。

```go
github.com/jacksonyoudi/gomodone v1.0.1 h1:jFd+qZlAB0R3zqrC9kwO8IgPrAdayMUS0rSHMDc/uG8=
github.com/jacksonyoudi/gomodone v1.0.1/go.mod h1:XWi+BLbuiuC2YM8Qz4yQzTSPtHt3T3hrlNN2pNlyA94=
github.com/jacksonyoudi/gomodone/v2 v2.0.0 h1:GpzGeXCx/Xv2ueiZJ8hEhFwLu7xjxLBjkOYSmg8Ya/w=
github.com/jacksonyoudi/gomodone/v2 v2.0.0/go.mod h1:L8uFPSZNHoAhpaePWUfKmGinjufYdw9c2i70xtBorSw=
```

这个内容是下面的，需要操作执行的结果

go run main.go 就可以运行了

### Making a bugfix release

假如 fix 一个 bug, 我们在 v1 版本上进行修复

修改代码如下：

```go
// say Hi to someone
func SayHi(name string) string {
-       return fmt.Sprintf("Hi, %s", name)
+       return fmt.Sprintf("Hi, %s!", name)
}
```

修复好，我们开始 push

```go
$ git commit -m "Emphasize our friendliness" say.go
$ git tag v1.0.1
$ git push --tags origin v1
```

#### Updating modules

刚才 fix bug，所以要在我们使用项目中更新

这个需要我们手动执行更新 module 操作

我们通过使用我们的好朋友来做到这一点 go get：

- 运行  go get -u 以使用最新的  minor  版本或修补程序版本（即它将从 1.0.0 更新到例如 1.0.1，或者，如果可用，则更新为 1.1.0）
- 运行  go get -u=patch 以使用最新的  修补程序  版本（即，将更新为 1.0.1 但不更新  为 1.1.0）

- 运行 go get package@version 以更新到特定版本（例如 github.com/jacksonyoudi/gomodone@v1.0.1）

目前 module 最新的也是 v1.0.1

```go
// 更新最新
$go get -u
$go get -u=patch
//指定包，指定版本
$go get github.com/jacksonyoudi/gomodone@v1.0.1
```

操作完，go.mod 文件会修改如下：

```go
module Gone

go 1.14

require (
    github.com/jacksonyoudi/gomodone v1.0.1
)
```

#### Major versions

根据语义版本语义，主要版本与次要版本 不同。主要版本可能会破坏向后兼容性。从 Go 模块的角度来看，主要版本是完全不同的软件包。乍一看这听起来很奇怪，但这是有道理的：两个不兼容的库版本是两个不同的库。比如下面修改，完全破坏了兼容性。

```go
package gomodone

import (
    "errors"
    "fmt"
)

// Hi returns a friendly greeting
// Hi returns a friendly greeting in language lang
func SayHi(name, lang string) (string, error) {
    switch lang {
    case "en":
        return fmt.Sprintf("Hi, %s!", name), nil
    case "pt":
        return fmt.Sprintf("Oi, %s!", name), nil
    case "es":
        return fmt.Sprintf("¡Hola, %s!", name), nil
    case "fr":
        return fmt.Sprintf("Bonjour, %s!", name), nil
    default:
        return "", errors.New("unknown language")
    }
}
```

如上，我们需要不同的大版本，这种情况下

修改 go.mod 如下

```go
module github.com/jacksonyoudi/gomodone/v2

go 1.14
```

然后，重新 tag，push

```go
$ git commit say.go -m "Change Hi to allow multilang"
$ git checkout -b v2 # 用于 v2 版本，后续修复 v2
$ git commit go.mod -m "Bump version to v2"
$ git tag v2.0.0
$ git push --tags origin v2 
```

### Updating to a major version

即使发布了库的新不兼容版本，现有软件 也不会中断，因为它将继续使用现有版本 1.0.1。go get -u 将不会获得版本 2.0.0。如果想使用 v2.0.0, 代码改成如下：

```go
package main

import (
    "fmt"
    "github.com/jacksonyoudi/gomodone/v2"
)

func main() {
    g, err := gomodone.SayHi("Roberto", "pt")
    if err != nil {
        panic(err)
    }
    fmt.Println(g)
}
```

执行 go mod tidy

```go
go: finding module for package github.com/jacksonyoudi/gomodone/v2
go: downloading github.com/jacksonyoudi/gomodone/v2 v2.0.0
go: found github.com/jacksonyoudi/gomodone/v2 in github.com/jacksonyoudi/gomodone/v2 v2.0.0
```

当然，两个版本都可以同时使用，使用别名如下：

```go
package main

import (
    "fmt"
    "github.com/jacksonyoudi/gomodone"
    mv2 "github.com/jacksonyoudi/gomodone/v2"
)

func main() {
    g, err := mv2.SayHi("Roberto", "pt")
    if err != nil {
        panic(err)
    }
    fmt.Println(g)

    fmt.Println(gomodone.SayHi("Roberto"))
}
```

执行一下 go mod tidy

### Vendoring

默认是忽略 vendor 的，如果想在项目目录下有 vendor 可以执行下面命令

```go
$go vendor
```

当然，如果构建程序的时候，希望使用 vendor 中的依赖

```go
$ go build -mod vendor
```

### IDEA 下开发 GO

1. 创建 go 项目

![img](.\note.assets\1636535762349-67a477c3-8ba4-4899-bdfb-24896bfa38fc.webp)

1. 创建完项目，会自动生成 go mod 文件 如果需要修改，可以手动修改，加入 git 等操作
2. 写业务逻辑代码![img](.\note.assets\1636535762370-e218d7be-97bd-48dd-b81b-cc942e940933.webp)

1. 解决依赖，更新 go.mod

   ![img](.\note.assets\1636535762396-fdcc4333-9d1d-4c41-9ed8-d842d85e47d9.webp)

   go build

