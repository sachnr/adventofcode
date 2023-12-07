package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type HandType int

const (
	FiveOfAKind HandType = iota
	FourOfAKind
	FullHouse
	ThreeOfAKind
	TwoPair
	OnePair
	HighCard
)

type Hand struct {
	cards    []rune
	bid      int
	handType HandType
}

var cardValues = map[rune]int{
	'2': 1,
	'3': 2,
	'4': 3,
	'5': 4,
	'6': 5,
	'7': 6,
	'8': 7,
	'9': 8,
	'T': 9,
	'J': 10,
	'Q': 11,
	'K': 12,
	'A': 13,
}

func Parse(input string, enableWildcard bool) ([]Hand, error) {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")

	var hands []Hand
	for _, line := range lines {
		parts := strings.Fields(line)
		no, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			return nil, err
		}

		cards := []rune(parts[0])
		counts := make(map[rune]int)
		for _, card := range parts[0] {
			counts[card] += 1
		}

		wildcard := counts['J']
		var cardType HandType
		switch len(counts) {
		case 1:
			cardType = FiveOfAKind
		case 2:
			four := false
			for _, val := range counts {
				if val == 4 {
					four = true
				}
			}

			if four {
				if enableWildcard && wildcard == 4 {
					cardType = FiveOfAKind
				} else if enableWildcard && wildcard == 1 {
					cardType = FiveOfAKind
				} else {
					cardType = FourOfAKind
				}
			} else {
				if enableWildcard && wildcard == 2 {
					cardType = FiveOfAKind
				} else if enableWildcard && wildcard == 3 {
					cardType = FiveOfAKind
				} else {
					cardType = FullHouse
				}
			}
		case 3:
			three := false
			for _, val := range counts {
				if val == 3 {
					three = true
				}
			}

			if three {
				if enableWildcard && wildcard == 1 {
					cardType = FourOfAKind
				} else if enableWildcard && wildcard == 3 {
					cardType = FourOfAKind
				} else {
					cardType = ThreeOfAKind
				}
			} else {
				if enableWildcard && wildcard == 2 {
					cardType = FourOfAKind
				} else if enableWildcard && wildcard == 1 {
					cardType = FullHouse
				} else {
					cardType = TwoPair
				}
			}
		case 4:
			if enableWildcard && wildcard == 2 {
				cardType = ThreeOfAKind
			} else if enableWildcard && wildcard == 1 {
				cardType = ThreeOfAKind
			} else {
				cardType = OnePair
			}
		case 5:
			if enableWildcard && wildcard == 1 {
				cardType = OnePair
			} else {
				cardType = HighCard
			}
		}

		hands = append(hands, Hand{cards: cards, bid: int(no), handType: cardType})
	}
	return hands, nil
}

type CustomQuickSortType struct {
	cardValues map[rune]int
	hands      []Hand
}

func (c CustomQuickSortType) Len() int {
	return len(c.hands)
}

func (c CustomQuickSortType) Less(i, j int) bool {
	if c.hands[i].handType > c.hands[j].handType {
		return true
	} else if c.hands[i].handType < c.hands[j].handType {
		return false
	}

	for k := 0; k < 5; k++ {
		if c.cardValues[c.hands[i].cards[k]] < cardValues[c.hands[j].cards[k]] {
			return true
		} else if c.cardValues[c.hands[i].cards[k]] > c.cardValues[c.hands[j].cards[k]] {
			return false
		}
	}
	return false
}

func (c CustomQuickSortType) Swap(i, j int) {
	c.hands[i], c.hands[j] = c.hands[j], c.hands[i]
}

func PartOne(hands []Hand) int {
	sortedHands := CustomQuickSortType{
		hands:      hands,
		cardValues: cardValues,
	}
	sort.Sort(sortedHands)

	var sum int
	for i, hand := range sortedHands.hands {
		sum += hand.bid * (i + 1)
	}
	return sum
}

func PartTwo(hands []Hand) int {
	cardValuesCloned := cardValues
	cardValuesCloned['J'] = 0

	sortedHands := CustomQuickSortType{
		hands:      hands,
		cardValues: cardValuesCloned,
	}
	sort.Sort(sortedHands)

	var sum int
	for i, hand := range sortedHands.hands {
		fmt.Printf("\nhand:%d %s %v", hand.handType, string(hand.cards), hand.bid)
		sum += hand.bid * (i + 1)
	}
	return sum

}

func main() {
	cards, err := Parse(input, false)
	if err != nil {
		panic(fmt.Sprintf("Parsing Failed: %v", err.Error()))
	}

	part1 := PartOne(cards)
	fmt.Printf("Day Seven:\n")
	fmt.Printf("\tPart One: %v\n", part1)

	cardsNew, err := Parse(input, true)
	if err != nil {
		panic(fmt.Sprintf("Parsing Failed: %v", err.Error()))
	}

	part2 := PartTwo(cardsNew)
	fmt.Printf("\tPart Two: %v\n", part2)

}
