package mylogger

import (
	"fmt"
	"os"
)

type Logger struct {
}

func NewLogger() Logger {
	return Logger{}
}

func (l Logger) Debug(x interface{}) {
	fmt.Fprintln(os.Stdout, x)
}
