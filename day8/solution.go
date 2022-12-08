package main

import (
	"fmt"
	"advent-of-code-2022/common"
)

const file_name string = "input.txt"

func partOne(fileLines []int) (int) {
	for _, line := range fileLines {
		fmt.Println(line)
	}
	return 0
}

func partTwo(fileLines []int) (int) {
	return 0
}

func main() {
	fileLines := common.ReadTextFileOfInt(file_name)
	var partOneResult = partOne(fileLines)
	var partTwoResult = partTwo(fileLines)
	fmt.Printf("Part one: %v \nPart two: %v", partOneResult, partTwoResult)
}

// func Calculate() (string) {
// 	fileLines := common.ReadTextFileOfString(file_name)
// 	var partOneResult = partOne(fileLines)
// 	var partTwoResult = partTwo(fileLines)
// 	return fmt.Sprintf("Part one: %v \nPart two: %v", partOneResult, partTwoResult)
// }
