package main

import (
	"fmt"
	"time"
)

func timeTaken(fName string) func() {
	startTime := time.Now()
	return func() {
		fmt.Printf("Time Taken by the function -- %s is %v \n \n", fName, time.Since(startTime))
	}
}

func testFunc() {
	defer timeTaken("testFunc")()
	for i := 0 ; i<= 20; i++ {
		fmt.Println(i)
		time.Sleep(200 * time.Millisecond)
	}
}

func main(){
	testFunc()
}
