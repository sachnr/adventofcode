package main

import (
	"fmt"
	"regexp"
)

var ValueMap = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	inputString := "two1nineeightwothreeabcone2threexyz"
	tokenRegex := regexp.MustCompile(`(one|three|four|five|six|seven|eight|nine|zero)`)
	matches := tokenRegex.FindAllStringSubmatch(inputString, -1)
	for _, match := range matches {
		fmt.Printf("%v", match)
	}
}
