package day8

import (
	"testing"
	"advent-of-code-2022/common"
)

func TestResults(t *testing.T) {
	fileLines := common.ReadTextFileOfInt("testInput.txt")
	var partOneResult, partTwoResult = partOne(fileLines), partTwo(fileLines)
	if partOneResult != 21 {
		t.Log("error partOneResult should be 21, partOneResult was", partOneResult)
		t.Fail()
	}

	if partTwoResult != 8 {
		t.Log("error partTwoResult should be 8, partTwoResult was", partTwoResult)
		t.Fail()
	}

}