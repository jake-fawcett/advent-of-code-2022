package day12

import (
	"os"
	"fmt"
	"strconv"
	"strings"
	"advent-of-code-2022/common"
	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
)

type Coord struct {
	xPos, yPos int
}

func coordHash(c Coord) (string) {
	return strconv.Itoa(c.xPos) + "," + strconv.Itoa(c.yPos)
}

func findLowest(fileLines []string) ([]Coord) {
	var lowestCoords []Coord
	for y, row := range fileLines {
		for x, height := range row {
			if height == rune('a') || height == rune('S') {
				lowestCoords = append(lowestCoords, Coord{x, y})
			}
		}
	}
	return lowestCoords
}

func addAllVertices(fileLines []string, g graph.Graph[string, Coord]) (graph.Graph[string, Coord], Coord, Coord, []string) {
	var start, goal Coord
	for y, row := range fileLines {
		for x, height := range row {
			if height == rune('S') {start = Coord{x, y}; fileLines[y] = strings.Replace(fileLines[y], "S", "a", 1)}
			if height == rune('E') {goal = Coord{x, y}; fileLines[y] = strings.Replace(fileLines[y], "E", "z", 1)}
			g.AddVertex(Coord{x, y})
		}
	}
	return g, start, goal, fileLines
}

func addAllEdges(fileLines []string, g graph.Graph[string, Coord], xLength, yLength int) (graph.Graph[string, Coord]) {
	for y, row := range fileLines {
		for x, height := range row {
			var adjacentNodes []Coord
			if x > 0 {adjacentNodes = append(adjacentNodes, Coord{x-1, y})}
			if x < xLength-1 { adjacentNodes = append(adjacentNodes, Coord{x+1, y})}
			if y > 0 {adjacentNodes = append(adjacentNodes, Coord{x, y-1})}
			if y < yLength-1 { adjacentNodes = append(adjacentNodes, Coord{x, y+1})}

			for _, node := range adjacentNodes {
				if rune(fileLines[node.yPos][node.xPos]) <= height+1 {
					g.AddEdge(coordHash(Coord{x, y}), coordHash(node), graph.EdgeWeight(1))
				}
			}
		}
	}
	return g
}

func findShortest(g graph.Graph[string, Coord], start, goal Coord) (int) {
	// The Graph package ShortestPath function produces inconsistent output 
	// To get the correct answer repeat 1000 times and get the shortest path
	// TODO: Replace use of graph package
	var shortest int = 999999
	for i := 0; i < 100; i++ {
		path, _ := graph.ShortestPath(g, coordHash(start), coordHash(goal))
		if len(path)-1 < shortest {
			shortest = len(path) - 1
		}
	}
	return shortest
}

func partOne(fileLines []string) (int) {
	xLength, yLength := len(fileLines[0]), len(fileLines)
	g := graph.New(coordHash, graph.Weighted(), graph.Directed())
	g, start, goal, fileLines := addAllVertices(fileLines, g)
	g = addAllEdges(fileLines, g, xLength, yLength)
	shortest := findShortest(g, start, goal)

	fmt.Println(goal)

	if os.Getenv("DRAW_GRAPH") == "true" {
		file, _ := os.Create("./mygraph.gv")
		_ = draw.DOT(g, file)
	}
	return shortest
}

func partTwo(fileLines []string) (int) {
	xLength, yLength := len(fileLines[0]), len(fileLines)
	g := graph.New(coordHash, graph.Weighted(), graph.Directed())
	g, _, goal, fileLines := addAllVertices(fileLines, g)
	g = addAllEdges(fileLines, g, xLength, yLength)

	fmt.Println(goal)

	// This is way too inefficient..
	var shortest, routeLength int = 999999, 999999
	var lowestCoords []Coord = findLowest(fileLines)
	for _, coord := range lowestCoords {
		routeLength = findShortest(g, coord, goal)
		if routeLength < shortest && routeLength > 0 {
			fmt.Println("Coord: ", coord, ", routeLength: ", routeLength, ", shortest: ", shortest)
			shortest = routeLength
		}
	}
	return shortest
}

func Calculate() (string) {
	// Have to copy the slice because of how we modify it, use arrays?
	fileLines := common.ReadTextFileOfString("day12/input.txt")
	fileLines2 := make([]string, len(fileLines))
	copy(fileLines2, fileLines)
	var partOneResult, partTwoResult = partOne(fileLines), partTwo(fileLines2)
	return fmt.Sprintf("Part one: %v \nPart two: %v", partOneResult, partTwoResult)
}
