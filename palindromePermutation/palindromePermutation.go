package main

import (
	"fmt"
	"strings"
)

func palindromePermutation(inputStr string) bool{
	var modifiedStr string
	singlesCount := 0
	modifiedStr = strings.Replace(inputStr, " ","", -1)
	result := make(map[rune]int)

	for _, eachChar := range modifiedStr{
		result[eachChar]++
	}
	fmt.Println(result)
	for _, v := range result {
		if v % 2 != 0{
			singlesCount++
		}
	}
	if singlesCount > 1{
		fmt.Println("NO!!! The string is not a permutation of a palindrome")
		return false
	}
	fmt.Println("Yes!!! The string is a permutation of a palindrome")
	return true
}

func main(){
	palindromePermutation("malayalam")
}
