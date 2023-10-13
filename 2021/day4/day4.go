package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

type Bingo struct {
	randomNos []int
	sheets    [][][]int
	score     [][][]bool
}

//go:embed input.txt
var input string

func (b *Bingo) Parse(input string) error {
	lines := strings.Split(input, "\n\n")
	randomNos := strings.Split(lines[0], ",")

	for _, value := range randomNos {
		value, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		b.randomNos = append(b.randomNos, value)
	}

	for i := 1; i < len(lines); i++ {
		sheet := lines[i]
		rows := strings.Split(sheet, "\n")

		layer := [][]int{}
		for _, row := range rows {
			values := strings.Fields(row)
			rowArr := []int{}
			for _, value := range values {
				value, err := strconv.Atoi(value)
				if err != nil {
					return err
				}
				rowArr = append(rowArr, value)
			}
			layer = append(layer, rowArr)
		}
		b.sheets = append(b.sheets, layer)
	}

	return nil
}

func (b *Bingo) initScore() {
	layers, rows, cols := len(b.sheets), len(b.sheets[0]), len(b.sheets[0][0])

	b.score = make([][][]bool, layers)
	for i := 0; i < layers; i++ {
		b.score[i] = make([][]bool, rows)
		for j := 0; j < rows; j++ {
			b.score[i][j] = make([]bool, cols)
			for k := 0; k < cols; k++ {
				b.score[i][j][k] = false
			}
		}
	}
}

func (b *Bingo) updateScores(value int) {
	for i, sheet := range b.sheets {
		for j, row := range sheet {
			for k, column := range row {
				if column == value {
					b.score[i][j][k] = true
				}
			}
		}
	}
}

func (b *Bingo) checkScores() (int, bool) {
	for i, sheet := range b.score {
		if hasBingo(sheet) {
			return i, true
		}
	}
	return 0, false
}

func hasBingo(sheet [][]bool) bool {
	// check row
	for _, row := range sheet {
		rowCompleted := true
		for _, column := range row {
			if column == false {
				rowCompleted = false
				break
			}
		}
		if rowCompleted {
			return true
		}
	}

	// check column
	for i, row := range sheet {
		colCompleted := true
		for j := range row {
			if sheet[j][i] == false {
				colCompleted = false
				break
			}
		}
		if colCompleted {
			return true
		}
	}

	return false
}

func (b *Bingo) CalculateScore(winner int, lastNo int) int {
	sum := 0
	for i, row := range b.score[winner] {
		for j, column := range row {
			if !column {
				sum += b.sheets[winner][i][j]
			}
		}
	}
	return sum * lastNo
}

func (b *Bingo) Run() (int, error) {
	var score int
	for _, number := range b.randomNos {
		b.updateScores(number)
		id, winner := b.checkScores()
		if winner {
			score = b.CalculateScore(id, number)
			break
		}
	}
	if score != 0 {
		return score, nil
	} else {
		return 0, fmt.Errorf("No Bingo")
	}
}

func main() {
	bingo := &Bingo{}
	err := bingo.Parse(input)

	if err != nil {
		panic("Failed to parse the input.")
	}

	bingo.initScore()
	score, err := bingo.Run()

	fmt.Printf("Part1: %d\n", score)
}
