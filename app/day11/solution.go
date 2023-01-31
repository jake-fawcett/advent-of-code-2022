package day11

import (
	"fmt"
	"strconv"
	"advent-of-code-2022/common"
)

type Monkey struct {
	items []int
	inspectedItems int
	operation string
	test string
	trueMonkey int
	falseMonkey int
}

func createMonkeys(fileLines []string) ([]Monkey) {
	var monkeyList []Monkey
	for i:=1; i<=len(fileLines); i+=7 {
		var items []int
		var startingItems = fileLines[i]
		for j:=18; j<len(startingItems); j+=4 {
			var item, _ = strconv.Atoi(startingItems[j:j+2])
			items = append(items, int(item))
		}
		var operation = fileLines[i+1][23:]
		var test = fileLines[i+2][8:]
		var trueMonkey, _ = strconv.Atoi(fileLines[i+3][29:])
		var falseMonkey, _ = strconv.Atoi(fileLines[i+4][30:])
		var monkey = Monkey{items, 0, operation, test, trueMonkey, falseMonkey}
		monkeyList = append(monkeyList, monkey)
	}
	return monkeyList
}

func takeTurn(numTurns int, worryDecrease int, monkeyList []Monkey) ([]Monkey) {
	var reduceModulo int = 1
	for _, monkey := range monkeyList {
		var testValue, _ = strconv.Atoi(monkey.test[13:])
		reduceModulo *= testValue
	}

	for i:=1; i<=numTurns; i++ {
		for j, monkey := range monkeyList {
			for _, item := range monkey.items {
				monkeyList[j].inspectedItems += 1

				var operationOperator = string(monkey.operation[0])
				var operationValue, worryLevel int
				if monkey.operation[2:] == "old" {
					operationValue = item
				} else {
					operationValue, _ = strconv.Atoi(monkey.operation[2:])
				}
				if operationOperator == "+" {
					worryLevel = ((item + operationValue) / worryDecrease) % reduceModulo
				} else if operationOperator == "*" {
					worryLevel = ((item * operationValue) / worryDecrease) % reduceModulo
				} 
				
				var testValue, _ = strconv.Atoi(monkey.test[13:])
				if worryLevel % testValue == 0 {
					monkeyList[monkey.trueMonkey].items = append(monkeyList[monkey.trueMonkey].items , worryLevel)
				} else {
					monkeyList[monkey.falseMonkey].items = append(monkeyList[monkey.falseMonkey].items , worryLevel)
				}
			}
			monkeyList[j].items = monkeyList[j].items[:0]
		}
	}
	return monkeyList
}

func calcMostActive(monkeyList []Monkey) (int, int) {
	var mostActive, secondMostActive int
	for _, monkey := range monkeyList {
		if monkey.inspectedItems > mostActive {
			secondMostActive = mostActive
			mostActive = monkey.inspectedItems
		} else if monkey.inspectedItems > secondMostActive {
			secondMostActive = monkey.inspectedItems
		}
	}
	return mostActive, secondMostActive
}

func partOne(fileLines []string) (int) {
	var monkeyList []Monkey = createMonkeys(fileLines)
	monkeyList = takeTurn(20, 3, monkeyList)
	mostActive, secondMostActive := calcMostActive(monkeyList)
	return (mostActive * secondMostActive)
}

func partTwo(fileLines []string) (int) {
	var monkeyList []Monkey = createMonkeys(fileLines)
	monkeyList = takeTurn(10000, 1, monkeyList)
	mostActive, secondMostActive := calcMostActive(monkeyList)
	return (mostActive * secondMostActive)
}

func Calculate() (string) {
	fileLines := common.ReadTextFileOfString("day11/input.txt")
	var partOneResult, partTwoResult = partOne(fileLines), partTwo(fileLines)
	return fmt.Sprintf("Part one: %v \nPart two: %v", partOneResult, partTwoResult)
}
