package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// 先创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 //黑子
	chessMap[2][3] = 2 //白子

	// 查看原始的数组
	for _, v := range chessMap {
		for _, chess := range v {
			fmt.Printf("%d\t", chess)
		}
		fmt.Println()
	}

	var sparseArr [][3]int
	// 记录二维数组的行列和默认值
	rowNum := len(chessMap)
	colNum := len(chessMap[0])
	sparseArr = append(sparseArr, [3]int{rowNum, colNum, 0})
	// 遍历原始数组转成稀疏数组
	for i := 0; i < rowNum; i++ {
		for j := 0; j < colNum; j++ {
			if chessMap[i][j] != 0 {
				sparseArr = append(sparseArr, [3]int{i, j, chessMap[i][j]})
			}
		}
	}

	// 输出稀疏数组
	for _, v := range sparseArr {
		for _, chess := range v {
			fmt.Printf("%d\t", chess)
		}
		fmt.Println()
	}

	// 保存到本地
	file, err := os.OpenFile("./Maker-Wu/DataStructure/spareArray/spare.txt",
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("open file failed:", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, v := range sparseArr {
		for _, chess := range v {
			writer.WriteString(strconv.Itoa(chess) + "\t")
		}
		writer.WriteString("\n")
	}
	writer.Flush()
}
