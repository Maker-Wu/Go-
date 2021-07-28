package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Args是一个[]string
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			fmt.Printf("args[%v]=%v\n", index, arg)
		}
	}
}
