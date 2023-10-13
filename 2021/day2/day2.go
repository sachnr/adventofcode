package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Course struct {
	inputs []Input
}

type Input struct {
	position string
	amount   int
}

func (p *Course) Parse(input string) error {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			continue // Skip empty lines
		}

		lineData := strings.SplitN(line, " ", 2)
		position := lineData[0]
		amount, err := strconv.Atoi(lineData[1])

		if err != nil {
			return err
		}

		p.inputs = append(p.inputs, Input{position, amount})
	}

	return nil
}

func (p *Course) Part1() int {
	depth := 0
	horz_pos := 0
	for _, input := range p.inputs {
		switch input.position {
		case "up":
			depth -= input.amount
		case "down":
			depth += input.amount
		case "forward":
			horz_pos += input.amount
		}
	}
	return depth * horz_pos
}

func (p *Course) Part2() int {
	depth := 0
	horz_pos := 0
	aim := 0

	for _, input := range p.inputs {
		switch input.position {
		case "up":
			aim -= input.amount
		case "down":
			aim += input.amount
		case "forward":
			horz_pos += input.amount
			depth += aim * input.amount
		}
	}
	return depth * horz_pos
}

func main() {
	course := &Course{}
	err := course.Parse(input)

	if err != nil {
		panic("failed to parse: " + err.Error())
	}
	fmt.Printf("Part1: %d\n", course.Part1())
	fmt.Printf("Part2: %d\n", course.Part2())
}
