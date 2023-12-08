package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type Directions struct {
	hashmap    map[string][2]string
	directions []int
}

func Parse(input string) Directions {
	out := new(Directions)
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")

	var directions []int
	for _, direction := range lines[0] {
		if direction == 'L' {
			directions = append(directions, 0)
		} else if direction == 'R' {
			directions = append(directions, 1)
		}
	}
	out.directions = directions

	hashmap := make(map[string][2]string)
	for _, maps := range lines[2:] {
		parts := strings.Split(maps, "=")
		key := strings.TrimSpace(parts[0])
		values := strings.Split(parts[1], ",")
		value1 := strings.TrimPrefix(strings.TrimSpace(values[0]), "(")
		value2 := strings.TrimSuffix(strings.TrimSpace(values[1]), ")")
		hashmap[key] = [2]string{value1, value2}
	}
	out.hashmap = hashmap

	return *out
}

func (d *Directions) PartOne() int {
	curr := "AAA"
	steps := 0
	i := 0
	for curr != "ZZZ" {
		steps += 1
		curr = d.hashmap[curr][d.directions[i]]
		if i < len(d.directions)-1 {
			i += 1
		} else {
			i = 0
		}
	}
	return steps
}

func Lcm(arr []int) int {
	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	lcm := func(a, b int) int {
		return a * b / gcd(a, b)
	}

	result := 1
	for _, value := range arr {
		result = lcm(result, value)
	}

	return result
}

func (d *Directions) PartTwo() int {
	var curr []string
	for key := range d.hashmap {
		if []rune(key)[2] == 'A' {
			curr = append(curr, key)
		}
	}

	var steps []int
	for _, value := range curr {
		currNode := value
		stepsPerValue := 0
		i := 0
		for []rune(currNode)[2] != 'Z' {
			stepsPerValue += 1
			currNode = d.hashmap[currNode][d.directions[i]]

			if i < len(d.directions)-1 {
				i += 1
			} else {
				i = 0
			}
		}
		steps = append(steps, stepsPerValue)
	}

	final := Lcm(steps)

	return final
}

func main() {
	directions := Parse(input)
	partone := directions.PartOne()

	fmt.Printf("DayEight:\n")
	fmt.Printf("\tPartOne: %v\n", partone)

	parttwo := directions.PartTwo()
	fmt.Printf("\tPartTwo: %v\n", parttwo)
}
