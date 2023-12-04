package main

import (
	"github.com/sachnr/adventofcode/2023/helper"
	"testing"
)

func TestDayThree(t *testing.T) {
	input := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

	t.Run("Check Direction", func(t *testing.T) {
		data := Parse(input)
		rec := Coord{row: 2, start: 2, end: 3}
		value := data.checkCoord(&rec)
		helper.Assert(t, value)

		rec1 := Coord{row: 0, start: 5, end: 7}
		value = data.checkCoord(&rec1)
		helper.Assert(t, !value)
	})

	t.Run("Check For Valid Directions", func(t *testing.T) {
		coords := []Coord{
			{row: 0, start: 0, end: 2},
			{row: 2, start: 2, end: 3},
			{row: 0, start: 5, end: 7},
		}
		data := Parse(input)
		validNos, err := data.GetValidNos(coords)
		if err != nil {
			t.Error(err.Error())
		}
		helper.AssertEq(t, validNos, []int{467, 35, 114})
	})

	t.Run("Part One", func(t *testing.T) {
		data := Parse(input)
		sum, err := PartOne(&data)
		if err != nil {
			t.Errorf("Parsing Failed: %v", err.Error())
		}
		helper.AssertEq(t, sum, 4361)
	})

	t.Run("Part Two", func(t *testing.T) {
		data := Parse(input)
		sum, err := PartTwo(&data)
		if err != nil {
			t.Errorf("Parsing Failed: %v", err.Error())
		}
		helper.AssertEq(t, sum, 467835)
	})

}
