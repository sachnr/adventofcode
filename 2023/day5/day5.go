package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Entry struct {
	from int
	to   int
	rng  int
}

var seeds []int

type Data struct {
	name    string
	entries []Entry
	next    *Data
}

func stringsToInt(inputs []string) ([]int, error) {
	var output []int
	for _, input := range inputs {
		no, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			return nil, err
		}
		output = append(output, int(no))
	}
	return output, nil
}

func Parse(input string) (*Data, error) {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")

	seedvals := strings.Fields(lines[0])[1:]
	nos, err := stringsToInt(seedvals)
	if err != nil {
		return nil, err
	}
	seeds = append(seeds, nos...)

	var curr *Data
	curr = new(Data)
	first := curr
	maps := lines[2:]
	for _, line := range maps {
		if line == "" {
			var next *Data
			next = new(Data)
			curr.next = next
			curr = next
			continue
		}
		fields := strings.Fields(line)
		if len(fields) == 2 {
			curr.name = fields[0]
			continue
		}
		no, err := stringsToInt(fields)
		if err != nil {
			return nil, err
		}
		curr.entries = append(curr.entries, Entry{to: no[0], from: no[1], rng: no[2]})
	}

	return first, nil
}

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func calculateDestination(entries []Entry, source int) int {
	for _, entry := range entries {
		start := entry.from
		end := entry.from + entry.rng
		if source >= start && source <= end {
			abs := abs(source - start)
			return entry.to + abs
		}
	}
	return source
}

func PartOne(data *Data) int {
	var locations []int
	for _, seed := range seeds {
		curr := data
		location := seed
		for curr != nil {
			location = calculateDestination(curr.entries, location)
			curr = curr.next
		}
		locations = append(locations, location)
	}

	min := locations[0]
	for _, location := range locations {
		if min > location {
			min = location
		}
	}
	return min
}

func PartTwo(data *Data) int {
	var seedsRange [][2]int
	for i := 0; i < len(seeds); i += 2 {
		seedsRange = append(seedsRange, [2]int{seeds[0], seeds[1]})
	}

	var locations []int
	for _, ranges := range seedsRange {
		seedStart := ranges[0]
		seedEnd := ranges[1] + seedStart
		for i := seedStart; i < seedEnd; i++ {
			curr := data
			location := i
			for curr != nil {
				location = calculateDestination(curr.entries, location)
				curr = curr.next
			}
			locations = append(locations, location)
		}
	}

	min := locations[0]
	for _, location := range locations {
		if min > location {
			min = location
		}
	}
	return min
}

func main() {
	fmt.Printf("Day Five:\n")

	data, err := Parse(input)
	if err != nil {
		panic(fmt.Sprintf("Parsing Failed: %v", err.Error()))
	}
	one := PartOne(data)
	fmt.Printf("\tPartOne: %v\n", one)

	two := PartTwo(data)
	fmt.Printf("\tParttwo: %v\n", two)
}
