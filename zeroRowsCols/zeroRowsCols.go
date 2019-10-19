package main

import "fmt"

func zeroRowsCols(inpMatrix [][]int) [][]int {
	l := len(inpMatrix)
	innerL := len(inpMatrix[0])
	rowZero := []int{}
	colZero := []int{}

	for i := 0; i < l ; i++{
		for j := 0; j < innerL; j++ {
			if inpMatrix[i][j] == 0 {
				rowZero = append(rowZero, i)
				colZero = append(colZero, j)
			}
		}
	}
	for _, eachRow := range rowZero{
		inpMatrix = zeroRows(inpMatrix, eachRow)
	}
	for _, eachCol := range colZero{
		inpMatrix = zeroCols(inpMatrix, eachCol)
	}

	return inpMatrix
}

func zeroRows(inpMatrix [][]int, rowNum int) [][]int {
	for j:= 0; j< len(inpMatrix[0]) ; j++ {
		inpMatrix[rowNum][j] = 0
	}
	return inpMatrix
}

func zeroCols(inpMatrix [][]int, colNum int) [][]int {
	for i := 0; i < len(inpMatrix); i++ {
		inpMatrix[i][colNum] = 0
	}
	return inpMatrix
}


func main() {
	testMatrix := [][]int{{1,2,3,4,5}, {1,0,3,5,6}, {8,7,6,5,6}, {23,11,0,11,56}, {77,8,4,2,1}}

	res := zeroRowsCols(testMatrix)

	fmt.Println(res)
}
