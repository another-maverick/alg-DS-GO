package main

import (
	"testing"
)

func TestOneEditAwayCheck(t *testing.T) {
	testCases := []struct {
		str1 string
		str2 string
		res bool
	}{{ "pale", "bale", true,},
		{"pale", "ale", true,},
	{"pale", "paleoooo",  false,},
	{"pale", "bake", false,}}

	for _, eachCase := range testCases{
		result := oneEditAwayCheck(eachCase.str1, eachCase.str2)
		if result != eachCase.res{
			t.Errorf("Expected: %v, Actual: %v", eachCase.res, result)
		}
	}

}
