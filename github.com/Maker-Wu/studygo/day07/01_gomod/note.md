Go mod模块是相关Go包的集合。modules源代码交换和版本控制的单元。go命令直接支持使用 modules，包括记录和解析对其他模块的依赖性。modules替换旧的基于GOPATH的方法来指定 在给定构建中使用哪些源文件。 

设置 GO111MODULE 

- off 

  go 命令行将不会支持 module 功能，寻找依赖包的方式将会沿用旧版本那种通过 vendor 目录或者 GOPATH 模式来查找。 

- on 
  go命令行会使用modules，而一点也不会去GOPATH目录下查找。 auto 默认值，go命令行将会根据当前目录来决定是否启用module功能 当modules 功能启用时，依赖包的存放位置变更为$GOPATH/pkg，允许同一个package多个版本 并存，且多个项目可以共享缓存的 module。