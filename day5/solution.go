package day5

import (
	"fmt"
	"common"
	"regexp"
	"strconv"
)

const file_name string = "input.txt"

type stack []string

func (s stack) Push(v string) stack {
    return append(s, v)
}

func (s stack) Pop() (stack, string) {
    l := len(s)
    return  s[:l-1], s[l-1]
}

func createStackList(fileLines []string) ([]stack) {
	var stackList []stack

	for crateStack:=0; crateStack<=8; crateStack++ {
		var stack stack
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

func getTopCrates(stackList []stack) (string) {
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
		var tempStack stack
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

func main() {
	fileLines := common.ReadTextFile(file_name)
	fmt.Println("Part one:", partOne(fileLines))
	fmt.Println("Part two:", partTwo(fileLines))
}
