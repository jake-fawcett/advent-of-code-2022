package day6

import (
	"fmt"
	"advent-of-code-2022/common"
	"strings"
)

const file_name string = "day6/input.txt"

func partOne(fileLines []string) (int) {
	var datastream string = fileLines[0]
	for i:=3; i<=len(datastream)-1; i++ {
		if !(strings.Contains(string(datastream[i-3:i]), string(datastream[i])) || 
			strings.Contains(string(datastream[i-3:i-1]), string(datastream[i-1])) || 
			strings.Contains(string(datastream[i-3:i-2]), string(datastream[i-2]))) { 
			return i + 1
		}
	}
	return 0
}

func partTwo(fileLines []string) (int) {
	var datastream string = fileLines[0]
	for i:=13; i<=len(datastream)-1; i++ {
		var messageMarker = true
		for j:=0; j<=13; j++ {
			if strings.Contains(string(datastream[i-13:i-j]), string(datastream[i-j])) { 
				messageMarker = false
				break
			}
		}
		if messageMarker {
			return i + 1
		}
	}
	return 0
}

// TODO: Use Sets!
func Calculate() (string) {
	fileLines := common.ReadTextFile(file_name)
	var partOneResult = partOne(fileLines)
	var partTwoResult = partTwo(fileLines)
	return fmt.Sprintf("Part one: %v \nPart two: %v", partOneResult, partTwoResult)
}
