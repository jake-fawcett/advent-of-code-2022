package common

import (
	"strconv"
	"bufio"
	"fmt"
	"os"
)

type StringStack []string

func (s StringStack) Push(v string) StringStack {
    return append(s, v)
}

func (s StringStack) Pop() (StringStack, string) {
    l := len(s)
    return  s[:l-1], s[l-1]
}

type IntStack []int

func (s IntStack) Push(v int) IntStack {
    return append(s, v)
}

func (s IntStack) Pop() (IntStack, int) {
    l := len(s)
    return  s[:l-1], s[l-1]
}

func Diff(a, b int) int {
	if a < b {
	   return b - a
	}
	return a - b
 }

func ReadTextFileOfString(file_name string) ([]string) {
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

func ReadTextFileOfInt(file_name string) ([][]int) {
	file, err := os.Open(file_name)
	if err != nil {
		fmt.Printf("Could not read the file due to this %s error \n", err)
	}
	fileScanner := bufio.NewScanner(file)
    fileScanner.Split(bufio.ScanLines)

	var fileLines [][]int
	for fileScanner.Scan() {
		var intRow []int
		for _, i := range fileScanner.Text() {
			var intI, _ = strconv.Atoi(string(i))
			intRow = append(intRow, intI)
		}
        if err != nil {
            fmt.Printf("Could not convert fileScanner.Text() to int due to this %s error \n", err)
        }
        fileLines = append(fileLines, intRow)
    }
    file.Close()
	return fileLines
}
		