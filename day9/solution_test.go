package day9

import (
	"testing"
	"advent-of-code-2022/common"
)

func TestResults(t *testing.T) {
	var fileLines = common.ReadTextFileOfString("testInput.txt")
	var partOneResult = partOne(fileLines)
	if partOneResult != 13 {
		t.Log("error partOneResult should be 13, partOneResult was", partOneResult)
		t.Fail()
	}

	fileLines = common.ReadTextFileOfString("testInput2.txt")
	var partTwoResult = partTwo(fileLines)
	if partTwoResult != 36 {
		t.Log("error partTwoResult should be 36, partTwoResult was", partTwoResult)
		t.Fail()
	}

}