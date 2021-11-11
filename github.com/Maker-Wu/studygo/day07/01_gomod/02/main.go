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