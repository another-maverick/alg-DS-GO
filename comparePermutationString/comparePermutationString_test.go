package main

import (
	"testing"
)

func TestCompareStrings(t *testing.T){
	testCases := []struct{
		string1 string
		string2 string
		res bool
	}{{"Hello", "oHell", true},
	{"hsgdhas","ajsdgas", false},
	{"heyy","eyhy", true}}

	for _, eachCase := range testCases{
		actualRes := compareStrings(eachCase.string1, eachCase.string2)
		if actualRes != eachCase.res{
			t.Errorf("The test case failed. actual result is %v and expected result is %v ", actualRes,eachCase.res)
		}
	}
}


