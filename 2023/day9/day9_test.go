package main

import (
	"github.com/sachnr/adventofcode/2023/helper"
	"testing"
)

func TestDayNine(t *testing.T) {
	input := `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

	histories, err := Parse(input)
	if err != nil {
		t.Errorf("Parsing Faild : %v", err.Error())
	}

	t.Run("Part One", func(t *testing.T) {
		got := PartOne(histories)
		helper.AssertEq(t, got, 114)
	})

	t.Run("Part Two", func(t *testing.T) {
		got := PartTwo(histories)
		helper.AssertEq(t, got, 2)
	})
}
