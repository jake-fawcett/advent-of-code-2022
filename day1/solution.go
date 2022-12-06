package day1

import (
	"strconv"
	"fmt"
	"advent-of-code-2022/common"
)

const file_name string = "day1/input.txt"

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

func Calculate() (string) {
	fileLines := common.ReadTextFile(file_name)
	var partOneResult, partTwoResult int = partOneAndTwo(fileLines)
	return fmt.Sprintf("Part one: %v \nPart two: %v", partOneResult, partTwoResult)
}
