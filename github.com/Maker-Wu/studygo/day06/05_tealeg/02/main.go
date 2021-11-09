package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func main() {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("第一页")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	row := sheet.AddRow()
	cell := row.AddCell()
	cell.SetString("姓名")
	err = file.Save("train.xlsx")
	if err != nil {
		panic(err.Error())
	}
}