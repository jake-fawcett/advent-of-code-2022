package main

import (
	"strings"
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

	var prioritySum int = 0
	fmt.Println(fileLines[0])
	fmt.Println(fileLines[1])
	fmt.Println(fileLines[len(fileLines)-1])
	for i:=0; i<=len(fileLines)-1; i+=3 {
		var line, line2, line3 = fileLines[i], fileLines[i+1], fileLines[i+2]
		for _, item := range line {
			if strings.Contains(line2, string(item)) && strings.Contains(line3, string(item)) {
				if item >= 97 {
					prioritySum += int(item - 96)
				} else {
					prioritySum += int(item - 38)
				}
				break
			}
		}
	}
	fmt.Println(prioritySum)
}

//  Part 1
// func main() {
// 	fileLines := read_file()

// 	var prioritySum int = 0
// 	for _, line := range fileLines {
// 		midpoint := len(line) / 2
// 		var compartment1, compartment2 string = line[0:midpoint], line[midpoint:]
// 		for _, item := range compartment1 {
// 			if strings.Contains(compartment2, string(item)) {
// 				if item >= 97 {
// 					prioritySum += int(item - 96)
// 				} else {
// 					prioritySum += int(item - 38)
// 				}
// 				break
// 			}
// 		}
//     }

// 	fmt.Println(prioritySum)
// }