package main

import (
	"testing"
)

func TestDayOne(t *testing.T) {
	t.Run("Part One", func(t *testing.T) {
		input := `199
200
208
210
200
207
240
269
260
263`

		report := &Report{}
		err := report.Parse(input)

		if err != nil {
			t.Error(err)
		}

		got := report.Part1()
		want := 7

		if got != want {
			t.Errorf("\nGot: %q\nWant: %q", got, want)
		}
	})

	t.Run("Part Two", func(t *testing.T) {
		input := `199
200
208
210
200
207
240
269
260
263`

		report := &Report{}
		err := report.Parse(input)

		if err != nil {
			t.Error(err)
		}

		got := report.Part2()
		want := 5

		if got != want {
			t.Errorf("\nGot: %d\nWant: %d", got, want)
		}
	})

}
