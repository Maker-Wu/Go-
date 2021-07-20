package main

import (
	"fmt"
	"time"
)

// 将时间戳转为时间对象
func timestampDemo2(timestamp int64) {
	timeObj := time.Unix(timestamp, 0)
	fmt.Printf("type of timeObj:%T\n", timeObj)	//type of timeObj:time.Time
	year := timeObj.Year()
	fmt.Printf("type of year:%T\n", year)	//type of year:int
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second) 	//2021-07-11 11:15:55
}

// 将时间对象转换为时间戳
func timestampDemo() {
	now := time.Now()
	timestamp1 := now.Unix()		//时间戳(秒)
	timestamp2 := now.UnixNano()	//纳秒时间戳
	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)
	timestampDemo2(timestamp1)
}

func timeDemo() {
	// time.Time类型表示时间(结构体类型)
	now := time.Now()		//获取当前的时间对象
	fmt.Printf("type of now:%T\n", now)		//type of now:time.Time
	fmt.Printf("current time:%v\n", now)		//current time:2021-07-11 11:16:42.8626597 +0800 CST m=+0.003106101

	year := now.Year()
	fmt.Printf("type of year:%T\n", year)	//type of year:int
	month := now.Month()
	fmt.Printf("type of month:%T\n", month)	//type of month:time.Month
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second) 	//2021-07-11 11:15:55
}

// 设置定时器，定时器本质上是一个通道(channel)
func tickDemo() {
	ticker := time.Tick(time.Second)	//定义一个1秒间隔的定时器
	for i := range ticker {
		fmt.Printf("type of i:%T\n", i)	//type of i:time.Time
		fmt.Println(i)
	}
}

// 时间对象转换为格式化时间.时间类型有一个自带的方法Format进行格式化
func formatDemo() {
	now := time.Now()
	//格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	//24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))		//2021-07-11 12:35:50.710 Sun Jul
	//12小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 PM Mon Jan"))		//
	fmt.Println(now.Format("2006/01/02 15:04"))						//2021/07/11 12:45
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
}

//解析字符串格式的时间
func formatDemo2() {
	now := time.Now()
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")	//可以填Local
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("type of loc:%T\n", loc)		//type of loc:*time.Location

	// 按照指定的时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2021/07/10 12:54:55", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(now.Sub(timeObj))		//24h2m4.9048176s
}

func main() {
	start := time.Now()
	timeDemo()
	timestampDemo()
	// 时间间隔
	n := 5
	time.Sleep(3*time.Second)
	time.Sleep(time.Duration(n) * time.Second)

	//tickDemo()
	formatDemo()
	formatDemo2()

	fmt.Printf("程序执行时间：%.2f", time.Now().Sub(start).Seconds())
}
