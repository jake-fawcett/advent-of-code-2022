package day10

import (
	"testing"
	"advent-of-code-2022/common"
)

func TestResults(t *testing.T) {
	var fileLines = common.ReadTextFileOfString("testInput.txt")
	var partOneResult = partOne(fileLines)
	if partOneResult != 13140 {
		t.Log("error partOneResult should be 13140, partOneResult was", partOneResult)
		t.Fail()
	}

	fileLines = common.ReadTextFileOfString("testInput.txt")
	var partTwoResult = partTwo(fileLines)
	var partTwoExpectedResult = [6]string{"##..##..##..##..##..##..##..##..##..##..", 
		"###...###...###...###...###...###...###.", 
		"####....####....####....####....####....",
		"#####.....#####.....#####.....#####.....",
		"######......######......######......####",
		"#######.......#######.......#######....."}
	if partTwoResult != partTwoExpectedResult {
		t.Log("error partTwoResult should be: ", partTwoExpectedResult, "but partTwoResult was: ", partTwoResult)
		t.Fail()
	}

}