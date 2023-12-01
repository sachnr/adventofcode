package main

import (
	"testing"

	"github.com/sachnr/adventofcode/2023/helper"
)

func TestDayOne(t *testing.T) {
	input1 := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

	input2 := `two1nine
	eightwothree
	abcone2threexyz
	xtwone3four
	4nineeightseven2
	zoneight234
	7pqrstsixteen`

	t.Run("parse only digits", func(t *testing.T) {
		numbers, err := CalculateNumbers(input1, false)
		if err != nil {
			t.Errorf("Error Parsing: %v", err.Error())
		}

		helper.AssertEq(t, []int{12, 38, 15, 77}, numbers)

		sum := Sum(numbers)
		helper.AssertEq[int](t, sum, 142)
	})

	t.Run("parse digits and letters", func(t *testing.T) {
		numbers, err := CalculateNumbers(input2, true)
		if err != nil {
			t.Errorf("Error Parsing: %v", err.Error())
		}

		helper.AssertEq(t, []int{29, 83, 13, 24, 42, 14, 76}, numbers)
		sum := Sum(numbers)
		helper.AssertEq[int](t, sum, 281)
	})

}
