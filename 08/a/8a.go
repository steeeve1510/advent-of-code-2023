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

func main() {
	lines := read("input/2")
	plan := parse(lines)

	position := "AAA"
	steps := 0
	for position != "ZZZ" {
		instruction := plan.instructions[steps%len(plan.instructions)]
		fork := plan.directions[position]
		if instruction == 'L' {
			position = fork.left
		} else if instruction == 'R' {
			position = fork.right
		} else {
			fmt.Println("Should not happen")
		}
		steps++
	}

	fmt.Println(steps)
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
