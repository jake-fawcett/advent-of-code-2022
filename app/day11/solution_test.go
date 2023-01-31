package day11

import (
	"testing"
	"advent-of-code-2022/common"
)

func TestResults(t *testing.T) {
	var fileLines = common.ReadTextFileOfString("testInput.txt")
	var partOneResult, partTwoResult = partOne(fileLines), partTwo(fileLines)
	if partOneResult != 10605 {
		t.Log("error partOneResult should be 10605, partOneResult was", partOneResult)
		t.Fail()
	}
	if partTwoResult != 2713310158 {
		t.Log("error partTwoResult should be: 2713310158, but partTwoResult was: ", partTwoResult)
		t.Fail()
	}

}