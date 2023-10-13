package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Report struct {
	depths []int
}

func (r *Report) Parse(v string) error {
	lines := strings.Split(v, "\n")
	r.depths = make([]int, len(lines))
	for i, str := range lines {
		if str == "" {
			continue // Skip empty lines
		}

		val, err := strconv.Atoi(str)

		if err != nil {
			return err
		}

		r.depths[i] = val
	}

	return nil
}

func (r *Report) slidingWindow() []int {
	var result []int
	for i := 0; i < len(r.depths)-2; i++ {
		sum := r.depths[i] + r.depths[i+1] + r.depths[i+2]
		result = append(result, sum)
	}
	return result
}

func (r *Report) Part2() int {
	larger := 0
	windows := r.slidingWindow()
	for i := 0; i < len(windows)-1; i++ {
		if windows[i] < windows[i+1] {
			larger += 1
		}

	}
	return larger
}

func (r *Report) Part1() int {
	larger := 0
	for i := 0; i < len(r.depths)-1; i++ {
		if r.depths[i] < r.depths[i+1] {
			larger += 1
		}

	}
	return larger
}

func main() {
	report := &Report{}
	err := report.Parse(input)
	if err != nil {
		panic("Error parsing input: " + err.Error())
	}
	fmt.Printf("part1: %d", report.Part1())
	fmt.Printf("part2: %d", report.Part2())
}
