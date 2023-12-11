package main

import (
	"github.com/sachnr/adventofcode/2023/helper"
	"testing"
)

func TestDayTen(t *testing.T) {
	input := `..F7.
.FJ|.
SJ.L7
|F--J
LJ...`

	t.Run("PartOne", func(t *testing.T) {
		maze := Parse(input)
		got := PartOne(maze)
		helper.AssertEq(t, got, 8)
	})
}
