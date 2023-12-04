package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Card struct {
	no        int
	winning   []int
	available []int
}

func Parse(input string) ([]Card, error) {
	lines := strings.Split(input, "\n")
	var cards []Card

	for i, line := range lines {
		parts1 := strings.Split(line, ":")
		parts2 := strings.Split(parts1[1], "|")

		winning := strings.Fields(parts2[0])
		available := strings.Fields(parts2[1])

		var winningNos []int
		var availableNos []int

		for _, no := range winning {
			intNo, err := strconv.ParseInt(no, 10, 32)
			if err != nil {
				return nil, err
			}
			winningNos = append(winningNos, int(intNo))
		}

		for _, no := range available {
			intNo, err := strconv.ParseInt(no, 10, 32)
			if err != nil {
				return nil, err
			}
			availableNos = append(availableNos, int(intNo))
		}

		cards = append(cards, Card{
			no:        i + 1,
			winning:   winningNos,
			available: availableNos,
		})
	}

	return cards, nil
}

func findMatching(card *Card) int {
	currTotal := 0
	matches := make(map[int]bool)

	for _, no := range card.winning {
		matches[no] = true
	}

	for _, no := range card.available {
		if matches[no] {
			currTotal += 1
		}
	}
	return currTotal
}

func partOne(cards []Card) int {
	var total int
	for _, card := range cards {
		matching := findMatching(&card)
		score := 0
		for i := 0; i < matching; i++ {
			if score == 0 {
				score = 1
				continue
			}
			score *= 2
		}
		total += score
	}

	return total
}

func PartTwo(cards []Card) int {
	totalCards := make(map[int]int)

	for _, card := range cards {
		totalCards[card.no] = 1
	}

	for _, card := range cards {
		if matching := findMatching(&card); matching > 0 {
			for i := card.no + 1; i <= card.no+matching; i++ {
				totalCards[i] += totalCards[card.no]
			}
		}
	}

	var sum int
	for _, values := range totalCards {
		sum += values
	}

	return sum
}

func main() {
	text := strings.TrimSpace(input)
	cards, err := Parse(text)
	if err != nil {
		panic(fmt.Sprintf("Parsing Faild: %v", err.Error()))
	}
	sum1 := partOne(cards)

	fmt.Printf("Day Three\n")
	fmt.Printf("\tPart One: %v\n", sum1)

	sum2 := PartTwo(cards)
	fmt.Printf("\tPart Two: %v\n", sum2)
}
