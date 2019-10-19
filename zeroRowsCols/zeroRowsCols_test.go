package main

import (
	"testing"
)

func TestZeroRowsCols(t *testing.T){
	testCases := []struct{
		inp [][]int
		out [][]int
	}{
		{[][]int{{1,2,3,4,5}, {1,0,3,5,6}, {8,7,6,5,6}, {23,11,0,11,56}, {77,8,4,2,1}},
			[][]int{{1,0,0,4,5}, {0,0,0,0,0}, {8,0,0,5,6}, {0,0,0,0,0}, {77,0,0,2,1}}},
	}
	for _, eachTestCase := range testCases{
		res := zeroRowsCols(eachTestCase.inp)
		if ok := Equal(res, eachTestCase.out) ; !ok {
			t.Errorf("the test case failed. \n Actual - %v  \n Expected - %v", res, eachTestCase.out)
		}
	}
}

func Equal(inp, out [][]int) bool {
	if len(inp) != len(out){
		return false
	}
	for i:=0 ; i  < len(inp); i++{
		for j :=0; j< len(inp[0]) ; j++ {
			if inp[i][j] != out[i][j]{
				return false
			}
		}
	}
	return true
}
