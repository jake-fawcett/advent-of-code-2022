package day9

import (
	"fmt"
	"strconv"
	"advent-of-code-2022/common"
)

// TODO: Improve to not hardcode gridsize and start, use alternative
const (
	grid_size int = 500
	starting_loc int = 250
	part_one_num_knots int = 2
	part_two_num_knots int = 10
)

type Coord struct {
	xPos, yPos int
}

func diff(a, b int) int {
	if a < b {
	   return b - a
	}
	return a - b
 }

func moveTail(prevKnot, knot Coord, knotNum, part int, visited [grid_size][grid_size]int) (Coord, [grid_size][grid_size]int) {
	var xDiff, yDiff = diff(knot.xPos, prevKnot.xPos), diff(knot.yPos, prevKnot.yPos)
	if xDiff == 2 && yDiff == 0 {
		knot.xPos += (prevKnot.xPos - knot.xPos) / 2 
	} else if yDiff == 2 && xDiff == 0 {
		knot.yPos += (prevKnot.yPos - knot.yPos) / 2 
	} else if xDiff == 2 && yDiff == 1 {	
		knot.xPos += (prevKnot.xPos - knot.xPos) / 2
		knot.yPos += (prevKnot.yPos - knot.yPos)
	} else if yDiff == 2 && xDiff == 1 {
		knot.xPos += (prevKnot.xPos - knot.xPos)
		knot.yPos += (prevKnot.yPos - knot.yPos) / 2
	} else if yDiff == 2 && xDiff == 2 {
		knot.xPos += (prevKnot.xPos - knot.xPos) / 2
		knot.yPos += (prevKnot.yPos - knot.yPos) / 2
	}

	if (part == 1 && knotNum == part_one_num_knots) || (part == 2 && knotNum == part_two_num_knots) {
		visited[knot.xPos][knot.yPos] = 1
	}
	return knot, visited
}

func partOne(fileLines []string) (int) {
	head, tail := Coord{starting_loc, starting_loc}, Coord{starting_loc, starting_loc}
	visited := [grid_size][grid_size]int{}
	totVisited := 0
	for _, line := range fileLines {
		var dir = string(line[0])
		var amount, _ = strconv.Atoi(line[2:])
		for i:=1; i<=amount; i++ {
			switch {
				case dir == "U":
					head.yPos += 1
				case dir == "D":
					head.yPos -= 1
				case dir == "R":
					head.xPos += 1
				case dir == "L":
					head.xPos -= 1
			}
			tail, visited = moveTail(head, tail, 2, 1, visited)
		}
	}

	for _, row := range visited {
		for _, element := range row {
			if element == 1 {
				totVisited += 1
			}
		}
	}
	return totVisited
}

func partTwo(fileLines []string) (int) {
	var ropeKnots[10]Coord
	for i:=0; i<=(part_two_num_knots-1); i++ {
		ropeKnots[i] = Coord{starting_loc, starting_loc}
	}
	visited := [grid_size][grid_size]int{}
	totVisited := 0
	for _, line := range fileLines {
		var dir = string(line[0])
		var amount, _ = strconv.Atoi(line[2:])
		for i:=1; i<=amount; i++ {
			switch {
				case dir == "U":
					ropeKnots[0].yPos += 1
				case dir == "D":
					ropeKnots[0].yPos -= 1
				case dir == "R":
					ropeKnots[0].xPos += 1
				case dir == "L":
					ropeKnots[0].xPos -= 1
			}
			for j:=1; j<=(part_two_num_knots-1); j++ {
				ropeKnots[j], visited =  moveTail(ropeKnots[j-1], ropeKnots[j], j+1, 2, visited)
			}
		}
	}

	for _, row := range visited {
		for _, element := range row {
			if element == 1 {
				totVisited += 1
			}
		}
	}
	return totVisited
}

// TODO: Improve efficiency 
func Calculate() (string) {
	fileLines := common.ReadTextFileOfString("day9/input.txt")
	var partOneResult, partTwoResult = partOne(fileLines), partTwo(fileLines)
	return fmt.Sprintf("Part one: %v \nPart two: %v", partOneResult, partTwoResult)
}
