package day3

import (
	"strings"
	"fmt"
	"advent-of-code-2022/common"
)

const file_name string = "day3/input.txt"

func partOne(fileLines []string) (int) {
	var prioritySum int = 0
	for _, line := range fileLines {
		midpoint := len(line) / 2
		var compartment1, compartment2 string = line[0:midpoint], line[midpoint:]
		for _, item := range compartment1 {
			if strings.Contains(compartment2, string(item)) {
				if item >= 97 {
					prioritySum += int(item - 96)
				} else {
					prioritySum += int(item - 38)
				}
				break
			}
		}
    }

	return prioritySum
}

func partTwo(fileLines []string) (int) {
	var prioritySum int = 0
	for i:=0; i<=len(fileLines)-1; i+=3 {
		var line, line2, line3 = fileLines[i], fileLines[i+1], fileLines[i+2]
		for _, item := range line {
			if strings.Contains(line2, string(item)) && strings.Contains(line3, string(item)) {
				if item >= 97 {
					prioritySum += int(item - 96)
				} else {
					prioritySum += int(item - 38)
				}
				break
			}
		}
	}
	return prioritySum
}

func Calculate() (string) {
	fileLines := common.ReadTextFile(file_name)
	var partOneResult = partOne(fileLines)
	var partTwoResult = partTwo(fileLines)
	return fmt.Sprintf("Part one: %v \nPart two: %v", partOneResult, partTwoResult)
}
