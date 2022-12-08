package day8

import (
	"fmt"
	"testing"
	"advent-of-code-2022/common"
)

func checkTreesVisible(tree int, treeList []int) (bool) {
	for _, prevTree := range treeList {
		if prevTree >= tree {
			return false
		}
	}
	return true
}

func checkNumTreesVisible(tree int, treeList []int) (int) {
	for num, prevTree := range treeList {
		if prevTree >= tree {
			return num + 1
		}
	}
	return len(treeList)
}

func reverse(arr []int) []int {
    reversed := make([]int, len(arr))
    j := 0
    for i := len(arr) - 1; i >= 0; i-- {
        reversed[j] = arr[i]
        j++
    }
    return reversed
}

func partOne(fileLines [][]int) (int) {
	var yLen int = len(fileLines)
	var xLen int = len(fileLines[0])
	var visibleTrees int = xLen * 2 + (yLen -2) * 2

	for y:=1; y<=len(fileLines)-2; y++ {
		var treeRow = fileLines[y]
		for x:=1; x<=len(treeRow)-2; x++ {
			var treeColumn []int
			for _, tempTreeRow := range fileLines {
				treeColumn = append(treeColumn, tempTreeRow[x])
			}
			var tree = treeRow[x]
			if checkTreesVisible(tree, treeRow[:x]) || checkTreesVisible(tree, treeRow[x+1:]) ||
				checkTreesVisible(tree, treeColumn[:y]) || checkTreesVisible(tree, treeColumn[y+1:]) {
					visibleTrees += 1
			} else {
					continue
			}
		}
	}
	return visibleTrees
}

func partTwo(fileLines [][]int) (int) {
	var scenicScore, maxScenicScore int = 0, 0
	for y:=1; y<=len(fileLines)-2; y++ {
		var treeRow = fileLines[y]
		for x:=1; x<=len(treeRow)-2; x++ {
			var treeColumn []int
			for _, tempTreeRow := range fileLines {
				treeColumn = append(treeColumn, tempTreeRow[x])
			}
			var tree = treeRow[x]
			var reversedTreeRow, reversedTreeColumn = reverse(treeRow[:x]), reverse(treeColumn[:y])
			scenicScore = checkNumTreesVisible(tree, reversedTreeRow) * checkNumTreesVisible(tree, treeRow[x+1:]) *
				checkNumTreesVisible(tree, reversedTreeColumn) * checkNumTreesVisible(tree, treeColumn[y+1:])
			if scenicScore > maxScenicScore {
				maxScenicScore = scenicScore
			}
		}
	}
	return maxScenicScore
}

func Calculate() (string) {
	fileLines := common.ReadTextFileOfInt("input.txt")
	var partOneResult, partTwoResult = partOne(fileLines), partTwo(fileLines)
	return fmt.Sprintf("Part one: %v \nPart two: %v", partOneResult, partTwoResult)
}

func TestPartOne(t *testing.T) {
	fileLines := common.ReadTextFileOfInt("testInput.txt")
	var partOneResult, partTwoResult = partOne(fileLines), partTwo(fileLines)
	if partOneResult != 21 {
		t.Log("error partOneResult should be 21")
		t.Fail()
	}

	if partTwoResult != 12 {
		t.Log("error partTwoResult should be 21")
		t.Fail()
	}

}