package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type race struct {
	time     int
	distance int
}

func main() {

	lines := read("input/1")
	races := parse(lines)

	result := 1
	for _, race := range races {
		winCounter := 0
		for i := 0; i <= race.time; i++ {
			distance := getDistance(i, race.time)
			if distance > race.distance {
				winCounter++
			}
		}
		result = result * winCounter
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

func parse(lines []string) []race {
	times := parseLine(lines[0])
	distances := parseLine(lines[1])

	result := make([]race, 0)
	for i, time := range times {
		distance := distances[i]
		result = append(result, race{time: time, distance: distance})
	}
	return result
}

func parseLine(line string) []int {
	parts := strings.Split(line, ":")
	numbers := parseNumbers(parts[1])
	result := make([]int, 0)
	for _, number := range numbers {
		result = append(result, int(number))
	}
	return result
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

func getDistance(buttonTime int, totalTime int) int {
	speed := buttonTime
	remainingTime := totalTime - buttonTime
	return speed * remainingTime
}
