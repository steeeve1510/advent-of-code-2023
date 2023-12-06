package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type mappingRange struct {
	destination uint64
	source      uint64
	ranges      uint64
}

type seedRange struct {
	start  uint64
	ranges uint64
}

type almanac struct {
	seeds    []seedRange
	mappings [][]mappingRange
}

func main() {

	lines := read("input/1")
	almanac := parse(lines)

	var result uint64
	result = math.MaxUint64
	for _, seedRange := range almanac.seeds {
		fmt.Printf("range: %+v\n", seedRange)
		for i := uint64(0); i < seedRange.ranges; i++ {
			m := seedRange.start + i
			for _, mapping := range almanac.mappings {
				m = resolve(m, mapping)
			}
			if m < result {
				result = m
			}
		}
	}

	fmt.Println(result)
}

func read(file string) []string {

	readFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var lines = make([]string, 0)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) == 0 {
			continue
		}
		lines = append(lines, line)
	}

	return lines
}

func parse(lines []string) almanac {
	seeds := make([]seedRange, 0)

	mappings := make([][]mappingRange, 0)
	mappings = append(mappings, make([]mappingRange, 0))
	mappingIndex := 0

	for i, line := range lines {
		if i == 0 {
			seeds = parseSeeds(line)
			continue
		}
		if i == 1 {
			continue
		}
		m, error := parseMappingRange(line)
		if error != nil {
			mappingIndex++
			mappings = append(mappings, make([]mappingRange, 0))
			continue
		}
		mappings[mappingIndex] = append(mappings[mappingIndex], m)
	}

	return almanac{
		seeds:    seeds,
		mappings: mappings,
	}
}

func parseSeeds(line string) []seedRange {
	parts := strings.Split(line, ": ")
	numbers := parseNumbers(parts[1])
	result := make([]seedRange, 0)
	for i, number := range numbers {
		if i%2 == 1 {
			continue
		}
		start := number
		ranges := numbers[i+1]
		result = append(result, seedRange{start: start, ranges: ranges})
	}
	return result
}

func parseMappingRange(line string) (mappingRange, error) {
	if strings.Contains(line, "map") {
		return mappingRange{}, errors.New("")
	}
	numbers := parseNumbers(line)
	return mappingRange{destination: numbers[0], source: numbers[1], ranges: numbers[2]}, nil
}

func parseNumbers(rawNumbers string) []uint64 {
	parts := strings.Split(rawNumbers, " ")

	numbers := make([]uint64, 0)
	for _, part := range parts {
		number, err := strconv.ParseUint(part, 10, 64)
		if err != nil {
			continue
		}
		numbers = append(numbers, number)
	}

	return numbers
}

func resolve(number uint64, mapping []mappingRange) uint64 {
	mapped := number
	for _, mappingRange := range mapping {
		start := mappingRange.source
		end := mappingRange.source + mappingRange.ranges
		if start <= number && number < end {
			diff := number - start
			return mappingRange.destination + diff
		}
	}
	return mapped
}
