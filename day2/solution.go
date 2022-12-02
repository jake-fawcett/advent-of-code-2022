package main

import (
	"bufio"
	"fmt"
	"os"
)

const file_name string = "input.txt"

func read_file() ([]string) {
	file, err := os.Open(file_name)
	if err != nil {
		fmt.Printf("Could not read the file due to this %s error \n", err)
	}

	fileScanner := bufio.NewScanner(file)
    fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
        fileLines = append(fileLines, fileScanner.Text())
    }
  
    file.Close()

	return fileLines
}

func main() {
	fileLines := read_file()

	var score int = 0
	for _, line := range fileLines {
		var opponent = line[0:1]
		var expectedResult = line[2:3]
		
		switch {
			case expectedResult == "X":
			if opponent == "A" {
				score += 3
			} else if opponent == "B" {
				score += 1
			} else {
				score += 2
			}

			case expectedResult == "Y":
			score += 3
			if opponent == "A" {
				score += 1
			} else if opponent == "B" {
				score += 2
			} else {
				score += 3
			}

			case expectedResult == "Z":
			score += 6
			if opponent == "A" {
				score += 2
			} else if opponent == "B" {
				score += 3
			} else {
				score += 1
			}
		}
    }

	fmt.Println(score)
}

// Part One
// for _, line := range fileLines {
// 	var player1 = line[0:1]
// 	var player2 = line[2:3]
	
// 	switch {
// 		case player2 == "X":
// 		score += 1
// 		if player1 == "C" {
// 			score += 6
// 		} else if player1 == "A" {
// 			score += 3
// 		}

// 		case player2 == "Y":
// 		score += 2
// 		if player1 == "A" {
// 			score += 6
// 		} else if player1 == "B" {
// 			score += 3
// 		}

// 		case player2 == "Z":
// 		score += 3
// 		if player1 == "B" {
// 			score += 6
// 		} else if player1 == "C" {
// 			score += 3
// 		}

// 	}
// }