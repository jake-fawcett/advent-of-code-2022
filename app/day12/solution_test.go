package day12

import (
	"testing"
	"advent-of-code-2022/common"
)

func TestResults(t *testing.T) {
	var fileLines = common.ReadTextFileOfString("testInput.txt")
	fileLines2 := make([]string, len(fileLines))
	copy(fileLines2, fileLines)
	var partOneResult, partTwoResult = partOne(fileLines), partTwo(fileLines2)
	if partOneResult != 31 {
		t.Log("error partOneResult should be 31, partOneResult was", partOneResult)
		t.Fail()
	}
	if partTwoResult != 29 {
		t.Log("error partTwoResult should be: 29, but partTwoResult was: ", partTwoResult)
		t.Fail()
	}

}