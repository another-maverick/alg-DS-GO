package main

import "fmt"

type myMatrix [][]int

func rotate90Degrees(inpMatrix myMatrix) myMatrix {
	l := len(inpMatrix)
	outMatrix := [][]int{}
	outMatrix_temp := []int{}

	for i:= 0; i < l ; i++{
		outMatrix_temp = []int{}
		for j :=l-1; j>=0; j--{
			outMatrix_temp = append(outMatrix_temp, inpMatrix[j][i])
		}
		outMatrix = append(outMatrix, outMatrix_temp)
	}
	return outMatrix
}

func main(){
	inpMatrix := [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}}
	fmt.Println(inpMatrix)
	res := rotate90Degrees(inpMatrix)
	fmt.Println(res)
}
