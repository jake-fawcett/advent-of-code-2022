package main

import (
	"fmt"
	"strings"
	"strconv"
	"advent-of-code-2022/common"
)

const file_name string = "input.txt"

func loadForest(fileLines []string) () {
	forest := [][]string{}

	for _, line := range fileLines {
		line := []string(strings.SplitAfter(line, ""))
		var forestRow []int
		for _, tree := range line {
			forestRow = append(forestRow, strconv.Atoi(tree))
		}
		forest = append(forest, forestRow)
	}
}

func partOne(fileLines []string) (int) {
	for _, line := range fileLines {
		for i, row := range line {
			fmt.Println(i)
			fmt.Println(string(row))
		}
		fmt.Println(line)
	}
	return 0
}

func partTwo(fileLines []string) (int) {
	return 0
}

func main() {
	fileLines := common.ReadTextFile(file_name)
	var partOneResult = partOne(fileLines)
	var partTwoResult = partTwo(fileLines)
	fmt.Printf("Part one: %v \nPart two: %v", partOneResult, partTwoResult)
}

// func Calculate() (string) {
// 	fileLines := common.ReadTextFile(file_name)
// 	var partOneResult = partOne(fileLines)
// 	var partTwoResult = partTwo(fileLines)
// 	return fmt.Sprintf("Part one: %v \nPart two: %v", partOneResult, partTwoResult)
// }
