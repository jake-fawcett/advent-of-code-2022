package main

import (
	"strconv"
	"fmt"
	"common"
)

const file_name string = "input.txt"

func partOneAndTwo(fileLines []string) (int, int) {
	var current, max1, max2, max3 int = 0, 0, 0, 0
	for _, line := range fileLines {
		if line == "" {
			if current > max1 {
				max3, max2, max1 = max2, max1, current
			} else if current > max2 {
				max3, max2 = max2, current
			} else if current > max3 {
				max3 = current
			}
			current = 0
		} else {
			var lineInt, _ = strconv.Atoi(line)
			current += lineInt
		}
    }
	return  max1 ,(max1 + max2 + max3)
}

func main() {
	fileLines := common.ReadTextFile(file_name)
	var partOneResult, partTwoResult int = partOneAndTwo(fileLines)
	fmt.Println("Part one:", partOneResult)
	fmt.Println("Part two:", partTwoResult)
}
