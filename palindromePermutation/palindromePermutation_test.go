package main

import (
	"testing"
)

func TestPalindromePermutation(t *testing.T){
	testCases := []struct{
		inp string
		res bool
	}{{
		"malayalam", true,
	},
	{
		"Hello", false,
	},
	{
		"dffdghhgtyytu", true,
	}}
	for _, eachCase := range testCases{
		result :=  palindromePermutation(eachCase.inp)
		if result != eachCase.res{
			t.Errorf("Test case has failed. Expected result is %v. Actual result is %v", eachCase.res, result)
		}
	}
}
