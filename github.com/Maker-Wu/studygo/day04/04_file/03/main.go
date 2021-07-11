package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	content, err := ioutil.ReadFile("./Maker-Wu/studygo/day04/04_file/note.md")
	if err != nil {
		fmt.Println("read file failed, err:", err)
	}
	fmt.Printf("type of content:%T\n", content)	//type of content:[]uint8
	fmt.Println(string(content))
}

