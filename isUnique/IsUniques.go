package main

import "fmt"

func IsUnique(inpStr string) bool {
	present := make(map[rune]struct{})

	for _, eachChar := range inpStr{
		if _, ok := present[eachChar]; ok {
			fmt.Println("The string does not have unique charecters")
			return false
		}else{
			present[eachChar] = struct{}{}
		}
	}
	fmt.Println("Yes!!!!! The string has unique charecters")
	return true
}

func main(){
	IsUnique("Hey")
}
