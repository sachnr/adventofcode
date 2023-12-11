package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type Maze struct {
	rows       int
	cols       int
	loop       [][]rune
	startIndex [2]int
}

type PipeDirection rune

const (
	Vertical     PipeDirection = '|'
	Horizontal   PipeDirection = '-'
	NorthAndEast PipeDirection = 'L'
	NorthAndWest PipeDirection = 'J'
	SouthAndWest PipeDirection = '7'
	SouthAndEast PipeDirection = 'F'
	Ground       PipeDirection = '.'
	Start                      = 'S'
)

type Direction int

const (
	North Direction = iota
	West
	South
	East
)

func Parse(input string) *Maze {
	maze := new(Maze)
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		if strings.ContainsRune(line, 'S') {
			maze.startIndex = [2]int{i, strings.IndexRune(line, 'S')}
		}
		maze.loop = append(maze.loop, []rune(line))
	}
	maze.rows = len(maze.loop)
	maze.cols = len(maze.loop[0])

	return maze
}

type Queueu struct {
	val [][2]int
}

func (q *Queueu) pushBack(r [2]int) {
	q.val = append(q.val, r)
}

func (q *Queueu) popFront() [2]int {
	val := q.val[0]
	q.val = q.val[1:]
	return val
}

func (q *Queueu) len() int {
	return len(q.val)
}

func PartOne(maze *Maze) int {
	visited := make(map[[2]int]bool)
	distances := make(map[[2]int]int)

	q := Queueu{}
	q.pushBack(maze.startIndex)

	for q.len() != 0 {
		curr := q.popFront()
		for i := North; i <= East; i++ {
			switch i {
			case North:
				if curr[0] < 1 {
					break
				}
				coords := [2]int{curr[0] - 1, curr[1]}

				if visited[coords] {
					break
				}

				char := maze.loop[coords[0]][coords[1]]
				if char == rune(Vertical) || char == rune(SouthAndEast) || char == rune(SouthAndWest) {
					q.pushBack(coords)
					distances[coords] = distances[curr] + 1
				}
			case West:
				if curr[1] >= maze.cols-1 {
					break
				}
				coords := [2]int{curr[0], curr[1] + 1}

				if visited[coords] {
					break
				}

				char := maze.loop[coords[0]][coords[1]]

				if char == rune(Horizontal) || char == rune(NorthAndWest) || char == rune(SouthAndWest) {
					q.pushBack(coords)
					distances[coords] = distances[curr] + 1
				}
			case South:
				if curr[0] >= maze.rows-1 {
					break
				}
				coords := [2]int{curr[0] + 1, curr[1]}

				if visited[coords] {
					break
				}

				char := maze.loop[coords[0]][coords[1]]

				if char == rune(NorthAndWest) || char == rune(NorthAndEast) || char == rune(Vertical) {
					q.pushBack(coords)
					distances[coords] = distances[curr] + 1
				}
			case East:
				if curr[1] < 1 {
					break
				}
				coords := [2]int{curr[0], curr[1] - 1}

				if visited[coords] {
					break
				}

				char := maze.loop[coords[0]][coords[1]]

				if char == rune(NorthAndEast) || char == rune(Horizontal) || char == rune(SouthAndEast) {
					q.pushBack(coords)
					distances[coords] = distances[curr] + 1
				}
			}
			visited[curr] = true
		}

	}
	max := 0
	for _, distance := range distances {
		if distance > max {
			max = distance
		}
	}
	return max
}

func main() {
	maze := Parse(input)
	partone := PartOne(maze)

	fmt.Printf("Day Ten:\n")
	fmt.Printf("\tPart One: %v\n", partone)
}
