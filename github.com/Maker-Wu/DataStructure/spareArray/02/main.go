package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	str1, err := ioutil.ReadFile("./Maker-Wu/DataStructure/spareArray/spare.txt")
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}

	// 解析成稀疏数组
	var sparseArr [][3]int
	var twoDimensionArr [][]int
	json.Unmarshal(str1, &sparseArr)

	rowNum, colNum, ele := sparseArr[0][0], sparseArr[0][1], sparseArr[0][2]
	twoDimensionArr = make([][]int, rowNum)
	for i := 0; i < rowNum; i++ {
		twoDimensionArr[i] = make([]int, colNum)
		for j := 0; j < colNum; j++ {
			twoDimensionArr[i][j] = ele
		}
	}

	for i := 1; i < len(sparseArr); i++ {
		rowIndex, colIndex, element := sparseArr[i][0], sparseArr[i][1], sparseArr[i][2]
		twoDimensionArr[rowIndex][colIndex] = element
	}
	fmt.Println(twoDimensionArr)
}
