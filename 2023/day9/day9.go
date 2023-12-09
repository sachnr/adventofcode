package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Histories struct {
	values [][]int
}

func Parse(input string) (*Histories, error) {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")

	histories := Histories{}
	for _, line := range lines {
		fields := strings.Fields(line)

		var nos []int
		for _, field := range fields {
			no, err := strconv.ParseInt(field, 10, 64)
			if err != nil {
				return nil, err
			}
			nos = append(nos, int(no))
		}
		histories.values = append(histories.values, nos)
	}

	return &histories, nil
}

func PartOne(histories *Histories) int {
	var finals int
	for _, history := range histories.values {
		sum := 0
		arr := history
		for {
			var newArr []int
			breakLoop := true
			for i := 0; i < len(arr)-1; i++ {
				value := arr[i+1] - arr[i]
				newArr = append(newArr, value)
				if value != 0 {
					breakLoop = false
				}
			}
			sum += arr[len(arr)-1]
			arr = newArr
			if breakLoop {
				break
			}
		}
		finals += sum
	}
	return finals
}

func PartTwo(histories *Histories) int {
	var finals int
	for _, history := range histories.values {
		var revSum []int
		arr := history
		for {
			var newArr []int
			breakLoop := true
			for i := 0; i < len(arr)-1; i++ {
				value := arr[i+1] - arr[i]
				newArr = append(newArr, value)
				if value != 0 {
					breakLoop = false
				}
			}
			revSum = append([]int{arr[0]}, revSum...)
			arr = newArr
			if breakLoop {
				break
			}
		}
		var sum int
		for _, val := range revSum {
			sum = val - sum
		}
		finals += sum
	}
	return finals
}

func main() {
	histories, err := Parse(input)
	if err != nil {
		panic(fmt.Sprintf("parsing faild: %v", err.Error()))
	}

	fmt.Printf("Day Nine\n")

	partone := PartOne(histories)
	parttwo := PartTwo(histories)
	fmt.Printf("\tPart One: %v\n", partone)
	fmt.Printf("\tPart Two: %v", parttwo)
}
