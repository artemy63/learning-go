package main

import (
	"fmt"
)

func isElementTwicedInArray(inputArray []int32) bool {
	hashMap := make(map[int32]int32)

	for _, arrEl := range inputArray {
		if hashMap[arrEl] > 0 {
			return true
		} else {
			hashMap[arrEl] = 1
		}
	}

	return false
}

func main() {
	fmt.Println("Start")

	var inputArray = []int32{1, 2, 3, 4, 5}
	var result = isElementTwicedInArray(inputArray)

	fmt.Println("result == ", result)

	inputArray = []int32{2, 3, 4, 2}
	result = isElementTwicedInArray(inputArray)
	fmt.Println("result == ", result)
}
