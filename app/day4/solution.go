package day4

import (
	"fmt"
	"advent-of-code-2022/common"
	"strings"
	"strconv"
)

const file_name string = "day4/input.txt"

func extractRanges(line string) (int, int, int, int) {
	var rangeOne = strings.Split(strings.Split(line, ",")[0], "-")
	var lowerRangeOne, _ = strconv.Atoi(rangeOne[0])
	var upperRangeOne, _ = strconv.Atoi(rangeOne[1])
	var rangeTwo = strings.Split(strings.Split(line, ",")[1], "-")
	var lowerRangeTwo, _ = strconv.Atoi(rangeTwo[0])
	var upperRangeTwo, _ = strconv.Atoi(rangeTwo[1])
	return lowerRangeOne, upperRangeOne, lowerRangeTwo, upperRangeTwo
}

func partOne(fileLines []string) (int) {
	var containedRanges int = 0
	for _, line := range fileLines {
		var lowerRangeOne, upperRangeOne, lowerRangeTwo, upperRangeTwo int = extractRanges(line)
		if (lowerRangeOne <= lowerRangeTwo && upperRangeOne >= upperRangeTwo) || (lowerRangeOne >= lowerRangeTwo && upperRangeOne <= upperRangeTwo) {
			containedRanges += 1
		} 
    }
	return containedRanges
}

func partTwo(fileLines []string) (int) {
	var containedRanges int = 0
	for _, line := range fileLines {
		var lowerRangeOne, upperRangeOne, lowerRangeTwo, upperRangeTwo int = extractRanges(line)
		if !(lowerRangeOne > upperRangeTwo || lowerRangeTwo > upperRangeOne) {
			containedRanges += 1
		} 
    }
	return containedRanges
}

func Calculate() (string) {
	fileLines := common.ReadTextFileOfString(file_name)
	var partOneResult = partOne(fileLines)
	var partTwoResult = partTwo(fileLines)
	return fmt.Sprintf("Part one: %v \nPart two: %v", partOneResult, partTwoResult)
}
