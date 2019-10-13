package main

import "testing"

func TestCompressString(t *testing.T){
	testStrings := []struct{
		inpStr string
		outStr string
	}{{
		"aabbccddee",
		"aabbccddee",
	},{
		"aaabbbcccddd",
		"a3b3c3d3",
	},{
		"aabbccddeeefffff",
		"a2b2c2d2e3f5",
	},
	}

	for _, eachCase := range testStrings{
		result := compressString(eachCase.inpStr)
		if result != eachCase.outStr{
			t.Errorf("THe actual string - %s is different from expected - %s", result, eachCase.outStr)
		}
	}
}
