package main

import (
	"github.com/sachnr/adventofcode/2023/helper"
	"testing"
)

func TestDayEight(t *testing.T) {

	var inputs = []struct {
		input string
		want  int
	}{
		{
			input: `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`,
			want: 2,
		},
		{
			input: `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`,
			want: 6,
		},
	}

    inputp2 := `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`

	t.Run("Part One", func(t *testing.T) {
		for _, input := range inputs {
			directions := Parse(input.input)
			got := directions.PartOne()
			helper.AssertEq(t, got, input.want)
		}
	})

	t.Run("Part Two", func(t *testing.T) {
		directions := Parse(inputp2)
		got := directions.PartTwo()
		helper.AssertEq(t, got, 6)
	})

}
