package main

import (
	"fmt"
	"strings"
)

func IntAbs(i int) int {
	if i < 0{
		return -i
	}
	return i
}

func IntMin(i int, j int) int{
	if i < j {
		return i
	}
	return j
}

func oneEditAwayCheck(str1 string, str2  string) bool {
	if IntAbs(len(str1) - len(str2)) > 1{
		fmt.Println("The difference in length is more than 1")
		return false
	}
	numOfDissimilarities := 0
	list1 := strings.Split(str1, "")
	list2 := strings.Split(str2, "")

	for i :=0; i < IntMin(len(str1), len(str2)); i++{
		if numOfDissimilarities < 2 {
			if list1[0] == list2[0]{
				list1 = list1[1:]
				list2 = list2[1:]
			}else{
				list1 = list1[1:]
				if list2[0] != list1[0]{
					list2 = list2[1:]
					numOfDissimilarities += 1
				}
			}
		}
	}
	if numOfDissimilarities <= 1 {
		fmt.Println("strings are one edit away!")
		return true
	}
	fmt.Println("NO!!! The strings are not one edit away")
	return false
}

func main() {
	oneEditAwayCheck("pale", "bale")
}
