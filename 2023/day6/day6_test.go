package main

import (
	"github.com/sachnr/adventofcode/2023/helper"
	"testing"
)

func TestDaySix(t *testing.T) {
	input := `Time:      7  15   30
Distance:  9  40  200`

	races, err := Parse(input)
	if err != nil {
		t.Errorf("Parsing Faild: %v", err.Error())
	}
	t.Run("Part One", func(t *testing.T) {
		want := 288
		got := PartOne(races)
		helper.AssertEq(t, got, want)
	})
	t.Run("Part Two", func(t *testing.T) {
		want := 71503
		race, err := ParseFixKerning(input)
		if err != nil {
			t.Errorf("Parsing Faild: %v", err.Error())
		}
		got := PartTwo(*race)
		helper.AssertEq(t, got, want)
	})
}
