package day7

import (
	"fmt"
	"strconv"
	"strings"
	"advent-of-code-2022/common"
)

const file_name string = "day7/input.txt"

func navigateDirs(currentDir, command string) (string) {
	if command == "/" {
		return "/"
	} else if command == ".." {
		var index = strings.LastIndex(currentDir[:len(currentDir)-1], "/")
		return currentDir[:index+1]
	} else {
		return (currentDir + command + "/")
	}
}

func partOne(fileLines []string) (map[string]int, int) {
	var dirs = make(map[string]int)
	var currentDir string

	for _, line := range fileLines {
		if string(line[0:4]) == "$ cd" {
			var command = line[5:]
			currentDir = navigateDirs(currentDir, command)
		} else if string(line[0:3]) == "dir" {
			continue
		} else {
			var fileSize, _ = strconv.Atoi(strings.Fields(line)[0])
			var tempDir string = currentDir
			
			// cascading file size up dirs
			for {
				if _, isMapContainsKey := dirs[tempDir]; isMapContainsKey {
					dirs[tempDir] += fileSize
				} else {
					dirs[tempDir] = fileSize
				}
				if tempDir == "/" {
					break
				}
				tempDir = navigateDirs(tempDir, "..")
			}
		}
	}

	var totalFileSize int = 0
	for _, fileSize := range dirs {
		if fileSize <= 100000 {
			totalFileSize += fileSize
		}
	}
	return dirs, totalFileSize
}

func partTwo(dirs map[string]int) (int) {
	var spaceToFree = 30000000 - (70000000 - dirs["/"])
	var smallestSuitableDir = 30000000
	for _, fileSize := range dirs {
		if fileSize < smallestSuitableDir && fileSize >= spaceToFree {
			smallestSuitableDir = fileSize
		}
	}

	return smallestSuitableDir
}

func Calculate() (string) {
	fileLines := common.ReadTextFile(file_name)
	var dirs, partOneResult = partOne(fileLines)
	var partTwoResult = partTwo(dirs)
	return fmt.Sprintf("Part one: %v \nPart two: %v", partOneResult, partTwoResult)
}
