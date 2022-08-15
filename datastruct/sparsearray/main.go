package main

import "fmt"

type ValNode struct {
	row int
	col int
	val int
}

func main() {
	//1.1先创建一个原始数组
	var chessMap [11][13]int
	chessMap[1][2] = 1
	chessMap[2][3] = 2

	//1.2输出看效果
	for _, v1 := range chessMap {
		for _, v2 := range v1 {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

	//1.3转成稀疏数组
	var sparseSlice []ValNode
	count := 0
	sparseSlice = append(sparseSlice, ValNode{
		row: len(chessMap),
		col: len(chessMap[0]),
		val: count,
	})
	for i, v1 := range chessMap {
		for j, v2 := range v1 {
			if v2 != 0 {
				valNode := ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseSlice = append(sparseSlice, valNode)
				count++
			}
		}
	}
	sparseSlice[0].val = count
	fmt.Println(sparseSlice)

	//2.将稀疏数组存盘...
	//3.从磁盘读取数据...

	//4稀疏数组转原始数组
	//4.1数组初始化时长度不能为变量要为constant 所以初始化一个二维切片
	zeroNode := sparseSlice[0]
	var newChessMap [][]int = make([][]int, zeroNode.row)
	for i := range newChessMap {
		newChessMap[i] = make([]int, zeroNode.col)
	}

	//4.2遍历 spaseSlice
	for i := 1; i < len(sparseSlice); i++ {
		node := sparseSlice[i]
		newChessMap[node.row][node.col] = node.val
	}
	//4.3输出看效果
	fmt.Println("恢复后的原始数组")
	for _, v1 := range newChessMap {
		for _, v2 := range v1 {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

}
