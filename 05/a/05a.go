package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type mappingRange struct {
	destination int
	source      int
	ranges      int
}

type almanac struct {
	seeds    []int
	mappings [][]mappingRange
}

func main() {

	lines := read("input/1")
	almanac := parse(lines)

	mapped := make([]int, 0)
	for _, seed := range almanac.seeds {
		m := seed
		for _, mapping := range almanac.mappings {
			m = resolve(m, mapping)
		}
		mapped = append(mapped, m)
	}

	min := slices.Min(mapped)
	fmt.Println(min)
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
	seeds := make([]int, 0)

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

func parseSeeds(line string) []int {
	parts := strings.Split(line, ": ")
	return parseNumbers(parts[1])
}

func parseMappingRange(line string) (mappingRange, error) {
	if strings.Contains(line, "map") {
		return mappingRange{}, errors.New("")
	}
	numbers := parseNumbers(line)
	return mappingRange{destination: numbers[0], source: numbers[1], ranges: numbers[2]}, nil
}

func parseNumbers(rawNumbers string) []int {
	parts := strings.Split(rawNumbers, " ")

	numbers := make([]int, 0)
	for _, part := range parts {
		number, err := strconv.Atoi(part)
		if err != nil {
			continue
		}
		numbers = append(numbers, number)
	}

	return numbers
}

func resolve(number int, mapping []mappingRange) int {
	mapped := number
	for _, mappingRange := range mapping {
		start := mappingRange.source
		end := mappingRange.source + mappingRange.ranges
		if start <= number && number <= end {
			diff := number - start
			return mappingRange.destination + diff
		}
	}
	return mapped
}
