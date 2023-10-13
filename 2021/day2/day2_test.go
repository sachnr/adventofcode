package main

import "testing"

func TestDayTwo(t *testing.T) {
	input := `forward 5
down 5
forward 8
up 3
down 8
forward 2`

	t.Run("Part One", func(t *testing.T) {
		course := &Course{}
		err := course.Parse(input)

		if err != nil {
			t.Errorf("Failed to parse the input")
		}

		got := course.Part1()
		want := 150

		if got != want {
			t.Errorf("\nGot: %q\nWant: %q", got, want)
		}
	})

	t.Run("Part Two", func(t *testing.T) {
		course := &Course{}
		err := course.Parse(input)

		if err != nil {
			t.Errorf("Failed to parse the input")
		}

		got := course.Part2()
		want := 900

		if got != want {
			t.Errorf("\nGot: %q\nWant: %q", got, want)
		}

	})

}
