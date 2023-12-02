package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

const (
	MAX_BLUE  = 14
	MAX_RED   = 12
	MAX_GREEN = 13
)

type Game struct {
	id    uint
	blue  uint
	red   uint
	green uint
}

func ParseInput(input string) ([]Game, error) {
	lines := strings.Split(input, "\n")

	var games []Game

	for _, line := range lines {
		if line == "" {
			continue
		}

		records := strings.SplitAfter(line, ":")
		idAsString := records[0]
		idAsString = strings.Fields(idAsString)[1]
		idAsString = strings.TrimSuffix(idAsString, ":")
		id, err := strconv.ParseUint(idAsString, 10, 32)
		if err != nil {
			return nil, err
		}

		var maxReds uint
		var maxGreens uint
		var maxBlues uint

		sets := strings.Split(records[1], ";")
		for _, set := range sets {
			cubes := strings.Split(set, ",")

			var reds uint
			var greens uint
			var blues uint

			for _, cube := range cubes {
				fields := strings.Fields(cube)
				noAsStrings := fields[0]
				no, err := strconv.ParseUint(noAsStrings, 10, 32)
				if err != nil {
					return nil, err
				}
				name := fields[1]

				switch name {
				case "red":
					reds += uint(no)
				case "blue":
					blues += uint(no)
				case "green":
					greens += uint(no)
				}
			}

			maxBlues = maxOfTwo(maxBlues, blues)
			maxReds = maxOfTwo(maxReds, reds)
			maxGreens = maxOfTwo(maxGreens, greens)
		}

		games = append(games, Game{
			id:    uint(id),
			red:   maxReds,
			blue:  maxBlues,
			green: maxGreens,
		})
	}
	return games, nil
}

func maxOfTwo(first, second uint) uint {
	if first > second {
		return first
	}
	return second
}

func compare(game *Game) bool {
	if game.red > MAX_RED {
		return false
	} else if game.blue > MAX_BLUE {
		return false
	} else if game.green > MAX_GREEN {
		return false
	}
	return true
}

func PartOne(games []Game) uint {
	var sum uint
	for _, game := range games {
		if compare(&game) {
			sum += game.id
		}
	}
	return sum
}

func PartTwo(games []Game) uint {
	var sum uint
	for _, game := range games {
		power := game.green * game.red * game.blue
		sum += power
	}
	return sum
}

func main() {
	games, err := ParseInput(input)
	if err != nil {
		panic(fmt.Sprintf("An error occurred: %v", err))
	}

	first := PartOne(games)
	fmt.Printf("DayTwo:\n")
	fmt.Printf("\tPartOne: %v\n", first)

	second := PartTwo(games)
	fmt.Printf("\tPartTwo: %v", second)
}
