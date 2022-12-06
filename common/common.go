package common

import (
	"bufio"
	"fmt"
	"os"
)

func ReadTextFile(file_name string) ([]string) {
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