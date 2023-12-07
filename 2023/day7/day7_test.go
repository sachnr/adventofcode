package main

import (
	"github.com/sachnr/adventofcode/2023/helper"
	"testing"
)

func TestDaySeven(t *testing.T) {
	var inputs = []struct {
		input string
		part1 int
		part2 int
	}{
		{
			input: `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`,
			part1: 6440,
			part2: 5905,
		},
		{
			input: `2345A 1
Q2KJJ 13
Q2Q2Q 19
T3T3J 17
T3Q33 11
2345J 3
J345A 2
32T3K 5
T55J5 29
KK677 7
KTJJT 34
QQQJA 31
JJJJJ 37
JAAAA 43
AAAAJ 59
AAAAA 61
2AAAA 23
2JJJJ 53
JJJJ2 41`,
			part1: 6592,
			part2: 6839,
		},
	}

	t.Run("Part One", func(t *testing.T) {
		for _, input := range inputs {
			cards, err := Parse(input.input, false)
			if err != nil {
				t.Errorf("Parsing Failed: %v", err.Error())
			}
			got := PartOne(cards)
			helper.AssertEq(t, got, input.part1)
		}
	})

	t.Run("Part Two", func(t *testing.T) {
		for _, input := range inputs {
			cards, err := Parse(input.input, true)
			if err != nil {
				t.Errorf("Parsing Failed: %v", err.Error())
			}
			got := PartTwo(cards)
			helper.AssertEq(t, got, input.part2)
		}
	})

}
