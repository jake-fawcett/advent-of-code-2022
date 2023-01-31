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
	"advent-of-code-2022/day8"
	"advent-of-code-2022/day9"
	"advent-of-code-2022/day10"
	"advent-of-code-2022/day11"
	"advent-of-code-2022/day12"
)

func askUserDay() (string) {
	fmt.Println("Warning: Day9 is currently inefficnet and slow.")
	fmt.Print("Enter number of day to run (0 for all): ")
	var input string
    _, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("An error occured while reading input: ", err)
	}
	return input
}

func main() {
	m := map[string]func()(string) {
		"day1": day1.Calculate,
		"day2": day2.Calculate,
		"day3": day3.Calculate,
		"day4": day4.Calculate,
		"day5": day5.Calculate,
		"day6": day6.Calculate,
		"day7": day7.Calculate,
		"day8": day8.Calculate,
		"day9": day9.Calculate,
		"day10": day10.Calculate,
		"day11": day11.Calculate,
		"day12": day12.Calculate,
	}

	chosenDay := askUserDay()
	if chosenDay != "0" {
		day := "day" + chosenDay
		fmt.Printf("- - - %s - - -\n %s \n\n", day, m[day]())
	} else {
		for day, function := range m {
			if day == "day12" { fmt.Println("Skipping Day12 as it is really inefficient") } else 
			{
				fmt.Printf("- - - %s - - -\n %s \n\n", day, function())
			}
		}
	}
}

	
