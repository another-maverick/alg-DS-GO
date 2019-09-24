package main

import (
	"testing"
)

func TestIsUnique(t *testing.T) {
	testCases := []struct{
		input string
		output bool
	}{{"Hello", false}, {"Hey",true}, {"howdy", true}}

	for _, eachCase := range testCases{
		var result = IsUnique(eachCase.input)
		if result != eachCase.output{
			t.Fatalf("input is %s and output is %v  and expected is %v \n",   eachCase.input, result,  eachCase.output)
		}
	}
}
