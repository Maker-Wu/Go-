package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Hero struct {
	Name string
	Age int
}

// HeroSlice 定义结构体切片类型
type HeroSlice []Hero

// Len 实现Interface接口中的Len方法
func (h HeroSlice) Len() int {
	return len(h)
}

func (h HeroSlice) Less(i, j int) bool {
	if h[i].Age <= h[j].Age {
		return true
	} else {
		return false
	}
}

func (h HeroSlice) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func main() {
	var heros HeroSlice
	fmt.Printf("%#v\n", heros)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		hero := Hero{
			Age: rand.Intn(50),
			Name: fmt.Sprintf("英雄%02d", rand.Intn(100)),
		}
		heros = append(heros, hero)
	}
	fmt.Println(heros)		//[{英雄06 43} {英雄95 46} {英雄93 33} {英雄41 36} {英雄30 9} {英雄51 19} {英雄91 17} {英雄04 32} {英雄81 19} {英雄67 16}]

	// 排序
	sort.Sort(heros)
	fmt.Println(heros)		//[{英雄72 17} {英雄87 19} {英雄36 23} {英雄32 35} {英雄82 36} {英雄67 39} {英雄04 40} {英雄61 42} {英雄74 46} {英雄27 49}]
}