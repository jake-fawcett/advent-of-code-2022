package day5

import (
	"fmt"
	"advent-of-code-2022/common"
	"regexp"
	"strconv"
)

const file_name string = "day5/input.txt"

func createStackList(fileLines []string) ([]common.StringStack) {
	var stackList []common.StringStack

	for crateStack:=0; crateStack<=8; crateStack++ {
		var stack common.StringStack
		for crate:=7; crate>=0; crate-- {
			var supply = string(fileLines[crate][(crateStack*4)+1])
			if supply != " " {
				stack = stack.Push(supply)
			}

		}
		stackList = append(stackList, stack)
	}
	return stackList
}

func getTopCrates(stackList []common.StringStack) (string) {
	var topCrates string
	var poppedCrate string
	for _, stack := range stackList {
		_, poppedCrate = stack.Pop()
		topCrates += poppedCrate
	}
	return topCrates
}

func partOne(fileLines []string) (string) {
	var stackList = createStackList(fileLines)

	re := regexp.MustCompile("[0-9]+")
	for _, line := range fileLines[10:] {
		var intStructions = re.FindAllString(line, -1)
		var numToMove, _ = strconv.Atoi(intStructions[0])
		var fromStack, _ = strconv.Atoi(intStructions[1])
		fromStack-- 
		var toStack, _ = strconv.Atoi(intStructions[2])
		toStack--
		var poppedCrate string
		for i:=1; i<=numToMove; i++ {
			stackList[fromStack], poppedCrate = stackList[fromStack].Pop()
			stackList[toStack] = stackList[toStack].Push(poppedCrate)
		}
    }

	return getTopCrates(stackList)
}

func partTwo(fileLines []string) (string) {
	var stackList = createStackList(fileLines)

	re := regexp.MustCompile("[0-9]+")
	for _, line := range fileLines[10:] {
		var intStructions = re.FindAllString(line, -1)
		var numToMove, _ = strconv.Atoi(intStructions[0])
		var fromStack, _ = strconv.Atoi(intStructions[1])
		fromStack-- 
		var toStack, _ = strconv.Atoi(intStructions[2])
		toStack--
		var poppedCrate string
		var tempStack common.StringStack
		for i:=1; i<=numToMove; i++ {
			stackList[fromStack], poppedCrate = stackList[fromStack].Pop()
			tempStack = tempStack.Push(poppedCrate)
		}
		for i:=1; i<=numToMove; i++ {
			tempStack, poppedCrate = tempStack.Pop()
			stackList[toStack] = stackList[toStack].Push(poppedCrate)
		}
    }

	return getTopCrates(stackList)
}

func Calculate() (string) {
	fileLines := common.ReadTextFileOfString(file_name)
	var partOneResult = partOne(fileLines)
	var partTwoResult = partTwo(fileLines)
	return fmt.Sprintf("Part one: %v \nPart two: %v", partOneResult, partTwoResult)
}
