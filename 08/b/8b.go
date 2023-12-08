package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type fork struct {
	left  string
	right string
}

type plan struct {
	instructions string
	directions   map[string]fork
}

type matching struct {
	firstMatch int
	beat       int
}

func main() {
	lines := read("input/1")
	plan := parse(lines)

	positions := make([]string, 0)
	for from := range plan.directions {
		if from[len(from)-1] == 'A' {
			positions = append(positions, from)
		}
	}

	numberOfInstructions := len(plan.instructions)

	matchings := make([]matching, len(positions))
	for i, position := range positions {
		steps := 0
		firstMatch := -1
		for !isAtEnd(position) || firstMatch < 0 {
			if isAtEnd(position) {
				firstMatch = steps
			}
			fork := plan.directions[position]
			instruction := plan.instructions[steps%numberOfInstructions]
			if instruction == 'L' {
				position = fork.left
			} else if instruction == 'R' {
				position = fork.right
			} else {
				fmt.Println("Should not happen")
			}
			positions[i] = position
			steps++
		}
		matchings[i] = matching{firstMatch: firstMatch, beat: steps - firstMatch}
	}

	currentSteps := make([]int, len(positions))
	for i, matching := range matchings {
		currentSteps[i] = matching.firstMatch
	}

	for !allMatch(currentSteps) {
		index := getMinIndex(currentSteps)
		currentSteps[index] += matchings[index].beat
	}
	fmt.Println(currentSteps[0])
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

func parse(lines []string) plan {
	instructions := lines[0]

	directions := make(map[string]fork)
	for i, line := range lines {
		if i == 0 {
			continue
		}
		from, fork := parseDirection(line)
		directions[from] = fork
	}
	return plan{instructions: instructions, directions: directions}
}

func parseDirection(line string) (string, fork) {
	reg := regexp.MustCompile(`[()\s]`)
	sanitized := reg.ReplaceAllString(line, "")

	parts1 := strings.Split(sanitized, "=")
	parts2 := strings.Split(parts1[1], ",")

	from := parts1[0]
	fork := fork{left: parts2[0], right: parts2[1]}

	return from, fork
}

func allMatch(list []int) bool {
	first := list[0]
	for _, entry := range list {
		if entry != first {
			return false
		}
	}
	return true
}

func getMinIndex(list []int) int {
	index := 0
	mininum := list[0]
	for i, entry := range list {
		if entry < mininum {
			index = i
			mininum = entry
		}
	}
	return index
}

func isAtEnd(position string) bool {
	return position[len(position)-1] == 'Z'
}
