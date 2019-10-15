package main

import (
	"fmt"
	"testing"
)

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestRotate90Degrees(t *testing.T){
	testCases := []struct{
		inp [][]int
		out [][]int
	}{
		{[][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}},
		[][]int{{21, 16, 11, 6, 1},{22, 17, 12, 7, 2}, {23, 18, 13, 8, 3}, {24, 19, 14, 9, 4}, {25, 20, 15, 10, 5}}},
	}
	for _, eachCase := range testCases{
		res := rotate90Degrees(eachCase.inp)
		fmt.Println(res)
		if ok := Equal(res[0],eachCase.out[0]) ; !ok {
			t.Errorf("actual %v is different from expectec - %v", res, eachCase.out)
		}
	}

}
