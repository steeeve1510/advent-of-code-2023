package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type position struct {
	x int
	y int
}

type plan struct {
	field []string
	start position
}

func main() {
	lines := read("input/1")
	plan := parsePlan(lines)

	steps := 1
	pos, _ := getNeighborsOfStart(plan.start, plan.field)

	previous := plan.start

	for pos != plan.start {
		nextPos1 := getNextPos(pos, previous, plan.field)
		previous = pos
		pos = nextPos1
		steps++
	}
	fmt.Println(steps / 2)
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

func parsePlan(lines []string) plan {

	field := make([]string, 0)
	start := position{x: -1, y: -1}
	for y, line := range lines {
		startX := strings.IndexRune(line, 'S')
		if startX >= 0 {
			start = position{x: startX, y: y}
		}
		field = append(field, line)
	}
	if start.x < 0 {
		fmt.Println("No start position found!")
	}
	return plan{field: field, start: start}
}

func getNeighbors(pos position, field []string) (position, position) {
	pipe := getPipe(pos, field)
	switch pipe {
	case '|':
		return north(pos), south(pos)
	case '-':
		return east(pos), west(pos)
	case 'L':
		return north(pos), east(pos)
	case 'J':
		return north(pos), west(pos)
	case '7':
		return south(pos), west(pos)
	case 'F':
		return south(pos), east(pos)
	default:
		return errorPos(), errorPos()
	}
}

func getNeighborsOfStart(pos position, field []string) (position, position) {
	neighbors := make([]position, 0)

	north1, north2 := getNeighbors(north(pos), field)
	if containsPos(pos, north1, north2) {
		neighbors = append(neighbors, north(pos))
	}
	east1, east2 := getNeighbors(east(pos), field)
	if containsPos(pos, east1, east2) {
		neighbors = append(neighbors, east(pos))
	}
	south1, south2 := getNeighbors(south(pos), field)
	if containsPos(pos, south1, south2) {
		neighbors = append(neighbors, south(pos))
	}
	west1, west2 := getNeighbors(west(pos), field)
	if containsPos(pos, west1, west2) {
		neighbors = append(neighbors, west(pos))
	}

	return neighbors[0], neighbors[1]
}

func getNextPos(pos position, previous position, field []string) position {
	next1, next2 := getNeighbors(pos, field)
	if previous == next1 {
		return next2
	} else {
		return next1
	}
}

func getPipe(pos position, field []string) rune {
	return rune(field[pos.y][pos.x])
}

func north(pos position) position {
	return position{x: pos.x, y: pos.y - 1}
}

func east(pos position) position {
	return position{x: pos.x + 1, y: pos.y}
}

func south(pos position) position {
	return position{x: pos.x, y: pos.y + 1}
}

func west(pos position) position {
	return position{x: pos.x - 1, y: pos.y}
}

func errorPos() position {
	return position{x: +1, y: -1}
}

func containsPos(pos position, first position, second position) bool {
	return pos == first || pos == second
}
