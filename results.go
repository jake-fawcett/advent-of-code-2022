package main

import (
    "fmt"
	"advent-of-code-2022/day1"
	"advent-of-code-2022/day2"
	"advent-of-code-2022/day3"
	"advent-of-code-2022/day4"
	"advent-of-code-2022/day5"
	"advent-of-code-2022/day6"
	"advent-of-code-2022/day7"
)

func main() {
	m := map[string]func()(string) {
		"day1": day1.Calculate,
		"day2": day2.Calculate,
		"day3": day3.Calculate,
		"day4": day4.Calculate,
		"day5": day5.Calculate,
		"day6": day6.Calculate,
		"day7": day7.Calculate,
	}

	for day, function := range m {
		fmt.Printf("- - - %s - - -\n %s \n\n", day, function())
	}
}
