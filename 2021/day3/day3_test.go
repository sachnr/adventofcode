package main

import "testing"

func TestDayThree(t *testing.T) {
	input := `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

	t.Run("Part One", func(t *testing.T) {
		diag := &Diagnostics{}
		diag.Parse(input)

		got, err := diag.Part1()

		if err != nil {
			t.Errorf("Failed to calculate")
		}

		want := 198

		if got != want {
			t.Errorf("\nGot: %v\nWant: %v", got, want)
		}
	})

	t.Run("filter Oxygen generator", func(t *testing.T) {
		diag := &Diagnostics{}
		diag.Parse(input)

		got, err := o2Gen(diag.binary)
		if err != nil {
			t.Errorf("err")
		}

		want := "10111"

		if got != want {
			t.Errorf("\nGot: %v\nWant: %v", got, want)
		}
	})

	t.Run("filter C02 generator", func(t *testing.T) {
		diag := &Diagnostics{}
		diag.Parse(input)

		got, err := co2Gen(diag.binary)
		if err != nil {
			t.Errorf("err")
		}

		want := "01010"

		if got != want {
			t.Errorf("\nGot: %v\nWant: %v", got, want)
		}
	})

	t.Run("Part Two", func(t *testing.T) {
		diag := &Diagnostics{}
		diag.Parse(input)

		got, err := diag.Part2()

		if err != nil {
			t.Errorf("Failed to calculate")
		}

		want := 230

		if got != want {
			t.Errorf("\nGot: %v\nWant: %v", got, want)
		}
	})

}
