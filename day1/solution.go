package main

import (
	"strconv"
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

	var current, max1, max2, max3 int = 0, 0, 0, 0
	for _, line := range fileLines {
		if line == "" {
			if current > max1 {
				max3, max2, max1 = max2, max1, current
			} else if current > max2 {
				max3, max2 = max2, current
			} else if current > max3 {
				max3 = current
			}
			current = 0
		} else {
			var lineInt, _ = strconv.Atoi(line)
			current += lineInt
		}
    }

	fmt.Println(max1 + max2 + max3)
}