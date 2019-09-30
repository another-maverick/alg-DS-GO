package main

import (
	"fmt"
	"strings"
)

func URLify(inputString string, contentLength int) string {
	inputList := strings.Split(inputString,"")
	//completeLength := len(inputList)
	numSpaces := 0

	fmt.Printf("the list after conversion from string is %v \n", inputList)
	outputList := make([]string, len(inputList))
	copy(outputList, inputList)

	for i := 0; i < contentLength ; i++ {
		if inputList[i] != " "{
			continue
		}else{
			for j := contentLength - 1 + 2 * numSpaces; j > i + 2 * numSpaces; j-- {
				outputList[j+2] = outputList[j]
			}
			copy(outputList[(i + 2*numSpaces) : (i + 2*numSpaces + 3)],strings.Split("%20","") )
			fmt.Println(outputList)
			numSpaces += 1
		}
	}
	fmt.Println(outputList)
	return strings.Join(outputList,"")
}

func main(){
	res := URLify("much ado about nothing      ", 22)
	fmt.Println("--- after URLify ---")
	fmt.Println(res)
}

