package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type TokenType int

const (
	Number TokenType = iota
	Symbol
	Dot
)

type Data struct {
	lenOfRows int
	lenOfCols int
	rows      [][]rune
	visited   [][]bool
	gears     map[string][]Coord
}

type Pos struct {
	x int
	y int
}

type Coord struct {
	row   int
	start int
	end   int
}

func (d *Data) GetCoords() []Coord {
	var coords []Coord
	for i, row := range d.rows {
		for j := range row {
			if d.visited[i][j] {
				continue
			}
			if checkType(d.rows[i][j]) == Number {
				curr := j
				for checkType(d.rows[i][curr]) == Number {
					d.visited[i][curr] = true
					curr += 1
					if curr >= d.lenOfRows {
						break
					}
				}
				coord := Coord{
					row:   i,
					start: j,
					end:   curr - 1,
				}
				coords = append(coords, coord)
			}
		}
	}
	return coords
}

func (d *Data) GetValidCoords(coords []Coord) []Coord {
	var validCoords []Coord

	for _, coord := range coords {
		if d.checkCoord(&coord) {
			validCoords = append(validCoords, coord)
		}
	}

	return validCoords
}

func (d *Data) coordToInt(coord Coord) (int, error) {
	var value []rune
	for i := coord.start; i <= coord.end; i++ {
		value = append(value, d.rows[coord.row][i])
	}
	no, err := strconv.ParseInt(string(value), 10, 32)
	return int(no), err
}

func (d *Data) GetValidNos(coords []Coord) ([]int, error) {
	var validNos []int

	for _, coord := range coords {
		no, err := d.coordToInt(coord)
		if err != nil {
			return nil, err
		}
		validNos = append(validNos, int(no))
	}

	return validNos, nil
}

func (d *Data) pushIntoMap(pos Pos, coord Coord) {
	key := fmt.Sprintf("%v,%v", pos.x, pos.y)
	if values, ok := d.gears[key]; ok {
		d.gears[key] = append(values, coord)
	} else {
		d.gears[key] = []Coord{coord}
	}
}

func (d *Data) checkCoord(coord *Coord) bool {
	checkLeft := coord.start > 0
	checkTop := coord.row > 0
	checkRight := coord.end < d.lenOfRows-1
	checkBottom := coord.row < d.lenOfRows-1

	var start int
	var end int
	if checkLeft {
		start = coord.start - 1
	} else {
		start = coord.start
	}
	if checkRight {
		end = coord.end + 1
	} else {
		end = coord.end
	}

	status := false

	for i := start; i <= end; i++ {
		if checkTop {
			topValue := d.rows[coord.row-1][i]
			if checkType(topValue) == Symbol {
				if topValue == '*' {
					pos := Pos{x: coord.row - 1, y: i}
					d.pushIntoMap(pos, *coord)
				}
				status = true
			}
		}
		if checkBottom {
			bottomValue := d.rows[coord.row+1][i]
			if checkType(bottomValue) == Symbol {
				if bottomValue == '*' {
					pos := Pos{x: coord.row + 1, y: i}
					d.pushIntoMap(pos, *coord)
				}
				status = true
			}
		}
		if checkLeft && i == start {
			leftValue := d.rows[coord.row][i]
			if checkType(leftValue) == Symbol {
				if leftValue == '*' {
					pos := Pos{x: coord.row, y: i}
					d.pushIntoMap(pos, *coord)
				}
				status = true
			}
		}
		if checkRight && i == end {
			rightValue := d.rows[coord.row][i]
			if checkType(rightValue) == Symbol {
				if rightValue == '*' {
					pos := Pos{x: coord.row, y: i}
					d.pushIntoMap(pos, *coord)
				}
				status = true
			}
		}
	}

	return status
}

func checkType(input rune) TokenType {
	if input >= '0' && input <= '9' {
		return Number
	} else if input == '.' {
		return Dot
	}
	return Symbol
}

func Parse(input string) Data {
	lines := strings.Split(input, "\n")
	var arr [][]rune
	for _, line := range lines {
		cells := []rune(line)
		arr = append(arr, cells)
	}

	rowSize := len(arr[0])
	colSize := len(arr)

	visisted := make([][]bool, colSize)
	for i := range visisted {
		visisted[i] = make([]bool, rowSize)
	}

	gears := make(map[string][]Coord)

	return Data{
		lenOfRows: rowSize,
		lenOfCols: colSize,
		rows:      arr,
		visited:   visisted,
		gears:     gears,
	}
}

func PartOne(data *Data) (int, error) {
	coords := data.GetCoords()
	validCoords := data.GetValidCoords(coords)
	validNos, err := data.GetValidNos(validCoords)
	if err != nil {
		return 0, err
	}
	var sum int
	for _, no := range validNos {
		sum += no
	}
	return sum, nil
}

func PartTwo(data *Data) (int, error) {
	var sum int
	coords := data.GetCoords()
	_ = data.GetValidCoords(coords)
	for _, coords := range data.gears {
		if len(coords) == 2 {
			first, err := data.coordToInt(coords[0])
			if err != nil {
				return 0, err
			}
			second, err := data.coordToInt(coords[1])
			if err != nil {
				return 0, err
			}
			sum += first * second
		}
	}
	return sum, nil
}

func main() {
	data := Parse(input)
	sum, err := PartOne(&data)
	if err != nil {
		panic(fmt.Sprintf("Parsing Failed: %v", err.Error()))
	}

	sum1, err := PartTwo(&data)
	if err != nil {
		panic(fmt.Sprintf("Parsing Failed: %v", err.Error()))
	}

	fmt.Printf("DayThree\n")
	fmt.Printf("\tPartOne: %d\n", sum)

	fmt.Printf("DayThree\n")
	fmt.Printf("\tPartTwo: %d\n", sum1)
}
