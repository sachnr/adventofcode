package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Race struct {
	time     int
	distance int
}

func convertToInt(inp string) (int, error) {
	no, err := strconv.ParseInt(inp, 10, 64)
	if err != nil {
		return 0, err
	}
	return int(no), nil
}

func Parse(input string) ([]Race, error) {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	time := strings.Fields(lines[0])[1:]
	distance := strings.Fields(lines[1])[1:]

	var races []Race
	for i := 0; i < len(time); i++ {
		time, err := convertToInt(time[i])
		if err != nil {
			return nil, err
		}
		distance, err := convertToInt(distance[i])
		if err != nil {
			return nil, err
		}
		races = append(races, Race{time, distance})
	}

	return races, nil
}

func ParseFixKerning(input string) (*Race, error) {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	time := strings.Fields(lines[0])[1:]
	distance := strings.Fields(lines[1])[1:]

	var timeFixed string
	var distanceFixed string
	for i := 0; i < len(time); i++ {
		timeFixed += time[i]
		distanceFixed += distance[i]
	}

	timeNo, err := strconv.ParseInt(timeFixed, 10, 64)
	if err != nil {
		return nil, err
	}
	distanceNo, err := strconv.ParseInt(distanceFixed, 10, 64)
	if err != nil {
		return nil, err
	}
	race := Race{
		distance: int(distanceNo),
		time:     int(timeNo),
	}

	return &race, nil
}

func PartOne(races []Race) int {
	var noOfWays []int
	for _, race := range races {
		var ways int
		for i := 1; i < race.time; i++ {
			timeLeft := race.time - i
			distanceTravelled := i * timeLeft
			if distanceTravelled > race.distance {
				ways += 1
			}
		}
		noOfWays = append(noOfWays, ways)
	}

	var product int
	for _, way := range noOfWays {
		if product == 0 {
			product = 1
		}
		product *= way
	}
	return product
}

func PartTwo(race Race) int {
	var ways int
	for i := 1; i < race.time; i++ {
		timeLeft := race.time - i
		distanceTravelled := i * timeLeft
		if distanceTravelled > race.distance {
			ways += 1
		}
	}

	return ways
}

func main() {
	races, err := Parse(input)
	if err != nil {
		panic(fmt.Sprintf("Parsing Faild: %v", err.Error()))
	}
	one := PartOne(races)
	fmt.Printf("Day 5\n")
	fmt.Printf("\tPartone: %v\n", one)

	race, err := ParseFixKerning(input)
	if err != nil {
		panic(fmt.Sprintf("Parsing Faild: %v", err.Error()))
	}
	two := PartTwo(*race)
	fmt.Printf("\tParttwo: %v", two)
}
