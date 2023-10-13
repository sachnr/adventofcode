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

func (b *Bingo) updateScores(value int) {
	for i, sheet := range b.sheets {
		for j, row := range sheet {
			for k, column := range row {
				if column == value {
					b.sheets[i][j][k] = 0
				}
			}
		}
	}
}

func (b *Bingo) checkScores(ignore []int) (int, bool) {
	for i, sheet := range b.sheets {
		ignoreFlag := false
		for _, value := range ignore {
			if value == i {
				ignoreFlag = true
				break
			}
		}
		if ignoreFlag {
			continue
		}
		if hasBingo(sheet) {
			return i, true
		}
	}
	return 0, false
}

func hasBingo(sheet [][]int) bool {
	// check row
	for _, row := range sheet {
		rowCompleted := true
		for _, column := range row {
			if column != 0 {
				rowCompleted = false
				break
			}
		}
		if rowCompleted {
			return true
		}
	}

	// check column
	for i := 0; i < len(sheet); i++ {
		colCompleted := true
		for j := 0; j < len(sheet); j++ {
			if sheet[j][i] != 0 {
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
	for i, row := range b.sheets[winner] {
		for j, column := range row {
			fmt.Printf(" %v", column)
			if column != 0 {
				sum += b.sheets[winner][i][j]
			}
		}
		fmt.Println()
	}
	return sum * lastNo
}

func (b *Bingo) Part1() (int, error) {
	var score int
	var ignore []int
	for _, number := range b.randomNos {
		b.updateScores(number)
		id, winner := b.checkScores(ignore)
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

func (b *Bingo) Part2() int {
	var winners []struct {
		id     int
		number int
	}
	var ignore []int
	for _, number := range b.randomNos {
		b.updateScores(number)
		id, winner := b.checkScores(ignore)
		if winner {
			winners = append(winners,
				struct {
					id     int
					number int
				}{id, number})
			ignore = append(ignore, id)
			if len(winners) == len(b.sheets) {
				break
			}
		}
	}
	length := len(winners)
	lastWinner := winners[length-1]
	return b.CalculateScore(lastWinner.id, lastWinner.number)
}

func main() {
	bingo := &Bingo{}
	err := bingo.Parse(input)

	if err != nil {
		panic("Failed to parse the input.")
	}

	score, err := bingo.Part1()

	fmt.Printf("Part1: %d\n", score)
}
