package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

type Entry struct {
	destination int
	source      int
	raenge      int
}

type Data struct {
	seeds                []int
	seedToSoil           []Entry
	soilToFertilize      []Entry
	fertilizerToWater    []Entry
	waterToLight         []Entry
	lightToTemprature    []Entry
	tempratureToHumidity []Entry
	humidityToLocation   []Entry
}

func stringToInt(inputs []string) ([]int, error) {
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
	data := Data{}

	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	afterSeeds := lines[2:]
	seedParts := strings.Split(lines[0], ":")
	seedValues := strings.Fields(seedParts[1])
	seeds, err := stringToInt(seedValues)
	if err != nil {
		return nil, err
	}
	data.seeds = seeds

	newEntry := true
	name := ""
	var entries []Entry
	for i, line := range afterSeeds {
		line = strings.TrimSpace(line)

		if newEntry {
			fields := strings.Fields(line)
			name = fields[0]
			newEntry = false
			continue
		}

		if line == "" || i == len(afterSeeds)-1 {
			switch name {
			case "seed-to-soil":
				data.seedToSoil = entries
			case "soil-to-fertilizer":
				data.soilToFertilize = entries
			case "fertilizer-to-water":
				data.fertilizerToWater = entries
			case "water-to-light":
				data.waterToLight = entries
			case "light-to-temperature":
				data.lightToTemprature = entries
			case "temperature-to-humidity":
				data.tempratureToHumidity = entries
			case "humidity-to-location":
				data.humidityToLocation = entries
			}
			newEntry = true
			entries = make([]Entry, 0)
			continue
		}

		fields := strings.Fields(line)
		numbers, err := stringToInt(fields)
		if err != nil {
			return nil, err
		}
		entries = append(entries, Entry{
			destination: numbers[0],
			source:      numbers[1],
			raenge:      numbers[2],
		})
	}
	return &data, nil
}

func absInt(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func calculateLocation(entries []Entry, location int) int {
	for _, entry := range entries {
		upto := entry.source + entry.raenge
		if location >= entry.source && location <= upto {
			abs := absInt(location) - absInt(entry.source)
			return entry.destination + abs
		}
	}
	return location
}

func PartOne(data *Data) int {
	var locations []int
	for _, seed := range data.seeds {
		location := seed
		for i := 0; i < 7; i++ {
			switch i {
			case 0:
				location = calculateLocation(data.seedToSoil, location)
			case 1:
				location = calculateLocation(data.soilToFertilize, location)
			case 2:
				location = calculateLocation(data.fertilizerToWater, location)
			case 3:
				location = calculateLocation(data.waterToLight, location)
			case 4:
				location = calculateLocation(data.lightToTemprature, location)
			case 5:
				location = calculateLocation(data.tempratureToHumidity, location)
			case 6:
				location = calculateLocation(data.humidityToLocation, location)
			}
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

func main() {
	fmt.Printf("Day Five:\n")

	startTime := time.Now()
	data, err := Parse(input)
	if err != nil {
		panic(fmt.Sprintf("Parsing Failed: %v", err.Error()))
	}
	one := PartOne(data)
	endTime := time.Since(startTime)
	fmt.Printf("\tPartOne: %v\n", one)
	fmt.Printf("\ttime: %v\n", endTime)
}
