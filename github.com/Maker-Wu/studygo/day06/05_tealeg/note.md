## 关键术语介绍

为了方便开源库的快速上手，我们先来了解 excel 中的几个关键术语，如下图所示，① 为 sheet，也就是表格中的页签；② 为 row，代表 excel 中的一行；③ 为 cell，代表 excel 中的一个单元格。

![image.png](F:\Users\WSQ\Go\src\Go-\github.com\Maker-Wu\studygo\day06\05_tealeg\note.assets\1636076847403-dd9d5886-40d0-4a12-9d35-0f536edb1356.png)

正常情况下，创建一个表格的基本流程是打开 wps 点击新建，这时会默认创建一个 sheet，然后在该 sheet 中的第一行填写表头，接下来根据表头逐行填充内容，最后将文件另存为到硬盘的某个位置。这与 Golang 开源库创建 excel 的流程基本相同，下面演示一个极简表格的创建。

## 创建表格

创建表格前需要先引入 excel 库，我们以比较热门的 tealeg/xlsx 库为例。

```go
go get github.com/tealeg/xlsx
```

首先创建一个空文件，拿到文件句柄。

```go
file := xlsx.NewFile()
```

创建一个名为人员信息收集的 sheet。

```go
sheet, err := file.AddSheet("人员信息收集")
if err != nil {
    panic(err.Error())
}
```

然后为该 sheet 创建一行，这行作为我们的表头。

```go
row := sheet.AddRow()
```

在该行中创建一个单元格。

```go
cell := row.AddCell()
```

现在给单元格填充内容，因为是表头，暂且叫姓名。

```go
cell.Value = "姓名"
```

如何创建第二个单元格呢？原理相同，此处 cell 变量已定义，再创建新单元格只需赋值即可。

```go
cell = row.AddCell()
cell.Value = "性别"
```

表头已经设置好了，可以开始创建第二行来填充内容了，方式与上述无差别。

```go
row = sheet.AddRow()
cell = row.AddCell()
cell.Value = "张三"
cell = row.AddCell()
cell.Value = "男"
```

表格设置完成后，将该文件保存，文件名可自定义。

```go
err = file.Save("demo.xlsx")
if err != nil {
	panic(err.Error())
}
```

跑起来后，可以发现目录中多了一个 demo.xlsx 文件，打开预览内容如下，达到了预期效果。

![image.png](F:\Users\WSQ\Go\src\Go-\github.com\Maker-Wu\studygo\day06\05_tealeg\note.assets\1636077616493-8c86b685-edcb-41ca-a402-d2e20c619f27.png)

**文件源码**

```go
package main

import "github.com/tealeg/xlsx"

func main() {
    file := xlsx.NewFile()
    sheet, err := file.AddSheet("人员信息收集")
    if err != nil {
        panic(err.Error())
    }
    row := sheet.AddRow()
    cell := row.AddCell()
    cell.Value = "姓名"
    cell = row.AddCell()
    cell.Value = "性别"

    row = sheet.AddRow()
    cell = row.AddCell()
    cell.Value = "张三"
    cell = row.AddCell()
    cell.Value = "男"

    err = file.Save("demo.xlsx")
    if err != nil {
        panic(err.Error())
    }
}
```
## 读取表格

表格的读取比创建简单很多，依然以上文创建的文件为例。

```go
output, err := xlsx.FileToSlice("demo.xlsx")
if err != nil {
	panic(err.Error())
}
```

只需将文件路径传入上述方法，即可自动读取并返回一个三维切片，我们来读取第一个 sheet 的第二行中的第二个单元格。

```go
log.Println(output[0][1][1]) //Output: 男
```

由此一来就非常容易遍历了。

```go
row = sheet.AddRow()
cell = row.AddCell()
cell.Value = "张三"
cell = row.AddCell()
cell.Value = "男"
```

文件源码

```go
package main

import (
    "fmt"
    "github.com/tealeg/xlsx"
    "log"
)

func main() {
    output, err := xlsx.FileToSlice("demo.xlsx")
    if err != nil {
        panic(err.Error())
    }
    log.Println(output[0][1][1])
    for rowIndex, row := range output[0] {
        for cellIndex, cell := range row {
            log.Println(fmt.Sprintf("第%d行，第%d个单元格：%s", rowIndex+1, cellIndex+1, cell))
        }
    }
}
```

## 修改表格

只是读取表格内容可能在特定场景下无法满足需求，有时候需要对表格内容进行更改。

```go
file, err := xlsx.OpenFile("demo.xlsx")
if err != nil {
	panic(err.Error())
}
```

修改表格之前依然需要先读取文件，只是这次并没有直接将其转化为三维切片。拿到文件句柄后，可以直接修改某一行的内容。

```go
file.Sheets[0].Rows[1].Cells[0].Value = "李四"
```

上述代码将第二行的张三改为了李四，但这还没有结束，接下来需要将文件重新保存。

```go
err = file.Save("demo.xlsx")
if err != nil {
	panic(err.Error())
}
```

打开文件预览，可以看到已经成功将张三改为了李四。

![image.png](F:\Users\WSQ\Go\src\Go-\github.com\Maker-Wu\studygo\day06\05_tealeg\note.assets\1636078515499-3e4fec7d-3b80-46b4-9380-10ff09214640.png)

文件源码

```go
package main

import "github.com/tealeg/xlsx"

func main() {
    file, err := xlsx.OpenFile("demo.xlsx")
    if err != nil {
        panic(err.Error())
    }
    file.Sheets[0].Rows[1].Cells[0].Value = "李四"
    err = file.Save("demo.xlsx")
    if err != nil {
        panic(err.Error())
    }
}
```

## 样式设置

该开源库不仅支持内容的编辑，还支持表格的样式设置，样式统一由结构体 Style 来负责。

```go
type Style struct {
    Border          Border
    Fill            Fill
    Font            Font
    ApplyBorder     bool
    ApplyFill       bool
    ApplyFont       bool
    ApplyAlignment  bool
    Alignment       Alignment
    NamedStyleIndex *int
}
```

拿上述生成的文件为例，假如我要将姓名所在单元格居中，首先要实例化样式对象。

```go
style := xlsx.NewStyle()
```

赋值居中属性。

```go
style.Alignment = xlsx.Alignment{
  Horizontal:   "center",
  Vertical:     "center",
}
```

给第一行第一个单元格设置样式。

```go
file.Sheets[0].Rows[0].Cells[0].SetStyle(style)
```

与修改表格处理逻辑相同，最后保存文件。

```go
err = file.Save("demo.xlsx")
if err != nil {
  panic(err.Error())
}
```

打开预览，可以看到文字已经上下左右居中。

![img](F:\Users\WSQ\Go\src\Go-\github.com\Maker-Wu\studygo\day06\05_tealeg\note.assets\1636078992072-238249a9-649d-41ba-8b3b-46ceb7414f0a.png)

同理，可以修改文字颜色和背景，同样通过 style 的属性来设置。
style.Font.Color = xlsx.RGB_Dark_Red
style.Fill.BgColor = xlsx.RGB_Dark_Green

其他还有很多属性可以设置，比如合并单元格、字体、大小等等，大家可以自行测试。
文件源码
package main

import "github.com/tealeg/xlsx"

func main() {
file, err := xlsx.OpenFile("demo.xlsx")
if err != nil {
panic(err.Error())
}
style := xlsx.NewStyle()
style.Font.Color = xlsx.RGB_Dark_Red
style.Fill.BgColor = xlsx.RGB_Dark_Green
style.Alignment = xlsx.Alignment{
Horizontal:   "center",
Vertical:     "center",
}
file.Sheets[0].Rows[0].Cells[0].SetStyle(style)
err = file.Save("demo.xlsx")
if err != nil {
panic(err.Error())
}
}

## 实战

昨天晚上回到实验室，发现同学在给老师记录本科生提交作业的情况。大概就是有一个文件列表，文件名是学生的学号、姓名、第几次作业组成。另外有一份学生名单的 excel，需要将提交了作业的学生信息记录在 excel 里。大概是下面这个图的样子：

![img](F:\Users\WSQ\Go\src\Go-\github.com\Maker-Wu\studygo\day06\05_tealeg\note.assets\1636101026691-d81b9bf3-19c3-4fb9-8a39-4b6dd8d204e1.png)

![img](F:\Users\WSQ\Go\src\Go-\github.com\Maker-Wu\studygo\day06\05_tealeg\note.assets\1636101047143-66bc92b3-390a-49bb-b1fc-fa205e2c8f69.png)

### 思路

我们需要读取文件列表，然后在excel中将对应学号的学生的提交情况打勾记录。

步骤简单的分为两步：

1. **读取文件列表，从中取出学号信息。**
2. **读取excel文件，在学号对应的行的单元格中进行记录。**

有个优化的细节：

每次我们取得学号之后，需要搜索整个表，才能知道该学号在哪一行。甚至当学生不在这个班级的时候，搜索完整个表之后没用任何收益，反而白白浪费了时间和计算资源。所以，我们读取表格的时候将**学号信息与行号信息进行缓存**，key位学号，value位行号。

### 第一版：基础功能

```go
package main

import (
	"github.com/tealeg/xlsx"
	"os"
	"strconv"
)

func main() {
	tablePath := "E:\\Desktop\\新建文件夹\\登分表.xlsx"

	// 获取表格sheet
	table, err := xlsx.OpenFile(tablePath)
	if err != nil {
		panic(err)
	}
	sheet := table.Sheets[0]

	// 缓存表格内容，key为学号，value为行号
	m := map[int]int{}
	for index, row := range sheet.Rows{
		number, err := strconv.Atoi(row.Cells[0].String())
		if err != nil {
			continue
		}
		m[number] = index
	}

	homeworkPath := "E:\\Desktop\\新建文件夹\\作业"
	files, err := os.ReadDir(homeworkPath)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		number, _ := strconv.Atoi(file.Name()[16:27])
		index, ok := m[number]
		if ok == false {
			continue
		}
		// 打勾
		sheet.Rows[index].Cells[5].SetString("√")
	}

	// 保存表格文件
	err = table.Save(tablePath)
	if err != nil {
		panic(err)
	}
}
```

### 第二版：正则匹配

然而，理想很丰满，现实很骨感，并不是所有同学都是按照统一格式来提交作业的，实际上提交的作业形式可能是这样的。

![img](F:\Users\WSQ\Go\src\Go-\github.com\Maker-Wu\studygo\day06\05_tealeg\note.assets\1636101394441-96db441f-7c6f-456c-adcf-af7aa7e3af47.webp)

可以看到，这位嘉夫1同学很有自己的想法，将自己的学号提取到了前面，这样我们就没法通过固定下表来提取到他的学号了。

所以，我们得采用一个新的办法从文件名中提取学号。

经过观察我们发现，学号都是固定长度的，比如都是11位的数字。

我们可以通过**正则表达式**来提取学号。

```go
// 用正则表达式匹配11位学号
reg := regexp.MustCompile(`\d{11}`)
for _, file := range files {
    number, _ = strconv.Atoi(reg.FindString(file.Name()))
    // 省略不变的部分
}
```

这样，不管同学将学号放在哪里，我都可以正确的提取出学号了。

### 第三版：多次作业

细心的同学已经发现，在登分册里不光有第一次作业，还有第二次作业。

作为一个上面的代码只能登记第一次作业，要想登记第二次作业还得重新修改代码。

作为一个程序员，这种机械的工作当然是交给代码来完成。

假设现在学生提交的文件变多了，里面包含的第二次作业。

![图片](F:\Users\WSQ\Go\src\Go-\github.com\Maker-Wu\studygo\day06\05_tealeg\note.assets\640.webp)

我们定义一个作业批次的数组或者切片，例如：`["第一次作业", "第二次作业"]`

这样，我们根据文件名包含哪个字符串就能知道是第几次作业了。

如果是第一次作业就往excel的第6列(索引下表为5)的单元格插入。

如果是第二次作业就往excel的第7列(索引下表为6)的单元格插入。

```go
homeworks := []string{"第一次作业", "第二次作业"}

// 用正则表达式匹配11位学号
reg := regexp.MustCompile(`[0-9]{11}`)
for _, file := range files {
  fileName := file.Name()
  for homeworkIndex, homework := range homeworks {
    if strings.Contains(fileName, homework) {
      number, _ := strconv.Atoi(reg.FindString(file.Name()))
      index, ok := m[number]
      if ok == false {
        continue
      }
      // 打勾
      sheet.Rows[index].Cells[homeworkIndex + 5].SetString("√")
    }
  }
}
```

完整的步骤解释：

1. 对每个文件名判断是否包含**第一次作业**或者**第二次作业**的字段。
2. 如果包含，则`homeworkIndex`用于记录插入的列，偏移量为`5`。

运行结束如下图所示：

![图片](F:\Users\WSQ\Go\src\Go-\github.com\Maker-Wu\studygo\day06\05_tealeg\note.assets\640.webp)

完美！不管学生把学号放在哪，我们只需要运行依次程序即可将两次的作业情况全部登记上了。

### 第四版：最终版

如果，你以为到第三版就结束了，那你可就大错特错了。

现在的学生讲究的是德智体美劳全面发展，发扬个性！

在我用上面的程序处理那600个学生的文件之后，发现还剩200多个学生没有登记上！

我一看，原来有的学生将**第一次作业**命名成**第1次作业**。

![图片](F:\Users\WSQ\Go\src\Go-\github.com\Maker-Wu\studygo\day06\05_tealeg\note.assets\640.webp)

好家伙，我直接好家伙！这都不按老师的要求来，这是我没有想到的。

继续修改代码，我们将作业的切片变为二维的，这样就不怕同学在一个批次的作业变着花命名了。

```go

homeworkSlice := [][]string{{"第一次作业", "第1次作业"},
    {"第二次作业", "第2次作业"}}

// 用正则表达式匹配11位学号
reg := regexp.MustCompile(`[0-9]{11}`)
for _, file := range files {
  fileName := file.Name()
  // homeworkIndex表示第几次作业
  for homeworkIndex, homeworks := range homeworkSlice {
    // 同一批次作业中可能有不同的命名
    for _, homework := range homeworks {
      if strings.Contains(fileName, homework) {
        number, _ := strconv.Atoi(reg.FindString(file.Name()))
        index, ok := m[number]
        if ok == false {
          continue
        }
        // 打勾
        sheet.Rows[index].Cells[homeworkIndex + 5].SetString("√")
      }
    }
  }
}
```

