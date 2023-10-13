package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Diagnostics struct {
	binary []string
}

func (d *Diagnostics) Parse(input string) {
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if line == "" {
			continue // Skip empty lines
		}

		d.binary = append(d.binary, line)
	}
}

func mostCommonValue(arr []string, index int) byte {
	zeroes := 0
	ones := 0

	for _, value := range arr {
		if value[index] == '0' {
			zeroes += 1
		} else {
			ones += 1
		}
	}

	if zeroes > ones {
		return '0'
	} else {
		return '1'
	}
}

func (d *Diagnostics) Part1() (int, error) {
	var gammaRate strings.Builder
	var epsilonRate strings.Builder
	len := len(d.binary[0])
	for i := 0; i < len; i++ {
		mcv := mostCommonValue(d.binary, i)

		if mcv == '0' {
			gammaRate.WriteByte('0')
			epsilonRate.WriteByte('1')
		} else {
			gammaRate.WriteByte('1')
			epsilonRate.WriteByte('0')
		}
	}

	gammaRateDecimal, err := strconv.ParseInt(gammaRate.String(), 2, 64)
	if err != nil {
		return 0, err
	}

	epsilonRateDecimal, err := strconv.ParseInt(epsilonRate.String(), 2, 64)
	if err != nil {
		return 0, err
	}

	return int(epsilonRateDecimal) * int(gammaRateDecimal), nil
}

func (d *Diagnostics) Part2() (int, error) {
	o2Gen, err := o2Gen(d.binary)
	if err != nil {
		return 0, err
	}
	o2GenBin, err := strconv.ParseInt(o2Gen, 2, 64)
	if err != nil {
		return 0, err
	}
	co2Gen, err := co2Gen(d.binary)
	if err != nil {
		return 0, err
	}
	co2GenBin, err := strconv.ParseInt(co2Gen, 2, 64)
	if err != nil {
		return 0, err
	}

	return int(o2GenBin) * int(co2GenBin), nil
}

func o2Gen(arr []string) (string, error) {
	lenArr := len(arr)
	o2Gen := arr
	i := 0
	for len(o2Gen) > 1 {
		if i >= lenArr {
			return "", fmt.Errorf("Out of bounds")
		}
		mostCommonValue := mostCommonValue(o2Gen, i)
		var o2 []string
		for _, value := range o2Gen {
			if value[i] == mostCommonValue {
				o2 = append(o2, value)
			}
		}
		o2Gen = o2
		i++
	}
	return o2Gen[0], nil
}

func co2Gen(arr []string) (string, error) {
	lenArr := len(arr)
	co2Gen := arr
	i := 0
	for len(co2Gen) > 1 {
		if i >= lenArr {
			return "", fmt.Errorf("Out of bounds")
		}
		mostCommonValue := mostCommonValue(co2Gen, i)
		var leastCommonValue byte
		if mostCommonValue == '0' {
			leastCommonValue = '1'
		} else {
			leastCommonValue = '0'
		}

		var o2 []string
		for _, value := range co2Gen {
			if value[i] == leastCommonValue {
				o2 = append(o2, value)
			}
		}
		co2Gen = o2
		i++
	}
	return co2Gen[0], nil
}

func main() {
	diag := &Diagnostics{}
	diag.Parse(input)
	part1Val, err := diag.Part1()

	if err != nil {
		panic("Failed")
	}

	fmt.Printf("Part1: %d\n", part1Val)

	part2Val, err := diag.Part2()

	if err != nil {
		panic("Failed")
	}

	fmt.Printf("Part2: %d\n", part2Val)
}
