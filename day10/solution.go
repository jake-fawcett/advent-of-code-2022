package day10

import (
	"fmt"
	"math"
	"strconv"
	"advent-of-code-2022/common"
)

func partOne(fileLines []string) (int) {
	var x, cycle, signalStrength int = 1, 1, 0
	for _, line := range fileLines {
		var instruction string = line[:4]
		if instruction == "addx" {
			var value, _ = strconv.Atoi(line[5:])
			if (cycle + 20) % 40 == 39{
				signalStrength += (x * (cycle+1))
			}
			cycle += 2
			x += value
			if (cycle + 20) % 40 == 0{
				signalStrength += (x * cycle)
			}
		} else {
			cycle += 1
			if (cycle + 20) % 40 == 0{
				signalStrength += (x * cycle)
			}
		}
	}
	return signalStrength
}

func renderPixel(crt [6]string, x, cycle int) ([6]string) {
	var xPos int = (cycle - 1) % 40
	var yPos int = int(math.Floor(float64((cycle - 1) / 40)))
	if common.Diff(xPos, x) <= 1 {
		crt[yPos] += "#"
	} else {
		crt[yPos] += "."
	}
	return crt
}

func partTwo(fileLines []string) ([6]string) {
	var x, cycle int = 1, 1
	crt := [6]string{}
	for _, line := range fileLines {
		var instruction string = line[:4]
		if instruction == "addx" {
			var value, _ = strconv.Atoi(line[5:])
			for i:=1; i<=2; i++ {
				crt = renderPixel(crt, x, cycle)
				cycle += 1
			}
			x += value
		} else {
			crt = renderPixel(crt, x, cycle)
			cycle += 1
		}
	}

	for _, value := range crt {
		fmt.Println(value)
	}
	return crt
}

func Calculate() (string) {
	fileLines := common.ReadTextFileOfString("day10/input.txt")
	var partOneResult, partTwoResult = partOne(fileLines), partTwo(fileLines)
	return fmt.Sprintf("Part one: %v \nPart two: %v", partOneResult, partTwoResult)
}
