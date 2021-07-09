#### 按照指定顺序遍历map

```go
func main() {
    // time.Now().UnixNano() 获取时间戳函数，单位纳秒
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap[key] = value
	}
	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}
```



#### 切片删除元素

```go
func main() {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3, 4)
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("%v\n", s)
	fmt.Printf("%v\n", m["q1mi"])
}
```

运行结果：

```go
[1 2 3]
[1 3]
[1 3 3]
```

得出结论：append删除元素，`s = append(s[:1], s[2:]...)`，其实就是把s[2:]...里的元素放在s[1:]，覆盖位置，所以这里会把[1, 2, 3]中的2覆盖成3，所以变成了[1, 3, 3]

==疑问==：那么为什么`m["q1mi"]`和s的长度不一样呢？

==答==：因为`m["q1mi"] = s`，这里就把s切片的三要素（指针其实位置、len、cap)都赋值给了m["q1mi"]。所以s的长度变化不会引起m["q1mi"]变化

但是s对底层数组的操作，却会引起m["q1mi"]的变化！因为他们对底层数组都使用的一个
