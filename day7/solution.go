package main

import (
	"fmt"
	"advent-of-code-2022/common"
)

const file_name string = "input.txt"

func partOne(fileLines []string) (int) {
	for _, line := range fileLines {
		fmt.Println(line)
	}
	return 0
}

func partTwo(fileLines []string) (int) {
	for _, line := range fileLines {
		fmt.Println(line)
	}
	return 0
}

func main() {
	fileLines := common.ReadTextFile(file_name)
	fmt.Println("Part one:", partOne(fileLines))
	fmt.Println("Part two:", partTwo(fileLines))
}
