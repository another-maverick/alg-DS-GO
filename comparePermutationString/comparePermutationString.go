package main

import "fmt"

func compareStrings(str1, str2 string) bool {
	if len(str1) != len(str2) {
		fmt.Println("strings are not identical in size")
		return false
	}
	result := make(map[rune]int)
	for _, eachChar := range str1 {
		result[eachChar]++
	}

	for _,eachChar := range str2{
		result[eachChar]--
	}
	for _, eachEntry := range result{
		if eachEntry != 0 {
			fmt.Println("The strings are NOT permutations of each other")
			return false
		}
	}
	fmt.Println("The strings are permutations of each other")
	return true
}


func main(){
	compareStrings("Hellooo","ooollHe")
}
