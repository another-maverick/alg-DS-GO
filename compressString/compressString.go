package main

import (
	"fmt"
	"strconv"
	"strings"
)

func myMin(a, b string) string{
	if len(b) < len(a) {
		return b
	}
	return a
}

func compressString(str1 string) string {
	outputString := []string{}
	charCount := 0

	inputArr := strings.Split(str1, "")
	//fmt.Printf("input string after converting to an array is %v \n",  inputArr)

	for i := 0; i < len(inputArr); i++ {
		if i != 0 && inputArr[i] != inputArr[i-1] {
			//fmt.Println(charCount)
			outputString = append(outputString, inputArr[i-1], strconv.Itoa(charCount))
			charCount = 0
		}
		charCount += 1
	}
	outputString = append(outputString, inputArr[len(inputArr)-1], strconv.Itoa(charCount))
	return myMin(str1, strings.Join(outputString, ""))
}

func main() {
	result := compressString("aabbccddeeefffff")
	fmt.Println(result)
}
