package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input1.txt
var input string

var ValueMap = map[string]string{
	"zero":  "z0o",
	"one":   "o1e",
	"two":   "t2o",
	"three": "t3e",
	"four":  "f4r",
	"five":  "f5e",
	"six":   "s6x",
	"seven": "s7n",
	"eight": "e8t",
	"nine":  "n9e",
}

func CalculateNumbers(input string, replaceLetters bool) ([]int, error) {
	var arr []int

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		var digits string

		if replaceLetters {
			line = LettersToInt(line)
		}

		for _, char := range line {
			if char >= '0' && char <= '9' {
				digits = digits + string(char)
			}
		}

		if digits == "" {
			continue
		}

		first := string(digits[0])
		last := string(digits[len(digits)-1])
		final := first + last
		number, err := strconv.Atoi(final)
		if err != nil {
			return nil, err
		}

		arr = append(arr, []int{number}...)
	}

	return arr, nil
}

func LettersToInt(line string) string {
	for letter, number := range ValueMap {
		if strings.Contains(line, letter) {
			line = strings.ReplaceAll(line, letter, number)
		}
	}
	return line
}

func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum = sum + number
	}
	return sum
}

func handleError(err error) {
	if err != nil {
		fmt.Printf("Error Parsing: %v", err.Error())
	}
}

func main() {
	numbers, err := CalculateNumbers(input, false)
	handleError(err)
	partOne := Sum(numbers)
	fmt.Printf("PartOne: %d", partOne)

	numbers, err = CalculateNumbers(input, true)
	handleError(err)
	partTwo := Sum(numbers)
	fmt.Printf("\nPartTwo: %d\n", partTwo)
}
