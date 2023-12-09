package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := read("input/1")
	measurements := make([][]int, 0)
	for _, line := range lines {
		measurement := parseMeasurement(line)
		measurements = append(measurements, measurement)
	}

	result := 0
	for _, measurement := range measurements {
		predicted := predict(measurement)
		result += predicted
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

func parseMeasurement(line string) []int {
	parts := strings.Split(line, " ")
	measurement := make([]int, 0)
	for _, part := range parts {
		number, _ := strconv.Atoi(part)
		measurement = append(measurement, number)
	}
	return measurement
}

func predict(measurement []int) int {
	derivates := make([][]int, 0)

	current := make([]int, len(measurement))
	copy(current, measurement)
	derivates = append(derivates, current)

	for !allZero(current) {
		derivate := derivate(current)
		derivates = append(derivates, derivate)
		current = derivate
	}

	predicted := 0
	for i := len(derivates) - 2; i >= 0; i-- {
		derivate := derivates[i]
		first := derivate[0]
		predicted = first - predicted
	}

	return predicted
}

func derivate(numbers []int) []int {
	diffs := make([]int, 0)
	for i := 0; i < len(numbers)-1; i++ {
		first := numbers[i]
		second := numbers[i+1]
		diff := second - first
		diffs = append(diffs, diff)
	}
	return diffs
}

func allZero(numbers []int) bool {
	for _, number := range numbers {
		if number != 0 {
			return false
		}
	}
	return true
}
