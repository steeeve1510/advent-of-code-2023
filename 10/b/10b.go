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
	lines := read("input/2")
	plan := parsePlan(lines)

	copy := emptyField(plan)

	pos, pos2 := getNeighborsOfStart(plan.start, plan.field)
	startPipe := determinePipe(plan.start, pos, pos2)
	setPipe(plan.start, startPipe, copy)
	previous := plan.start
	for pos != plan.start {
		pipe := getPipe(pos, plan.field)
		setPipe(pos, pipe, copy)

		nextPos := getNextPos(pos, previous, plan.field)
		previous = pos
		pos = nextPos
	}

	result := 0
	for _, line := range copy {
		count := countEnclosed(line)
		result += count
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

func getNextPos(pos position, previous position, field []string) position {
	next1, next2 := getNeighbors(pos, field)
	if previous == next1 {
		return next2
	} else {
		return next1
	}
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

	if isValid(north(pos), field) {
		north1, north2 := getNeighbors(north(pos), field)
		if containsPos(pos, north1, north2) {
			neighbors = append(neighbors, north(pos))
		}
	}
	if isValid(east(pos), field) {
		east1, east2 := getNeighbors(east(pos), field)
		if containsPos(pos, east1, east2) {
			neighbors = append(neighbors, east(pos))
		}
	}
	if isValid(south(pos), field) {
		south1, south2 := getNeighbors(south(pos), field)
		if containsPos(pos, south1, south2) {
			neighbors = append(neighbors, south(pos))
		}
	}
	if isValid(west(pos), field) {
		west1, west2 := getNeighbors(west(pos), field)
		if containsPos(pos, west1, west2) {
			neighbors = append(neighbors, west(pos))
		}
	}

	return neighbors[0], neighbors[1]
}

func determinePipe(pos position, next1 position, next2 position) rune {
	if next1 == north(pos) && next2 == south(pos) {
		return '|'
	}
	if next1 == east(pos) && next2 == west(pos) {
		return '-'
	}
	if next1 == north(pos) && next2 == east(pos) {
		return 'L'
	}
	if next1 == north(pos) && next2 == west(pos) {
		return 'J'
	}
	if next1 == south(pos) && next2 == west(pos) {
		return '7'
	}
	if next1 == east(pos) && next2 == south(pos) {
		return 'F'
	}
	return '.'
}

func getPipe(pos position, field []string) rune {
	return rune(field[pos.y][pos.x])
}

func setPipe(pos position, pipe rune, field []string) {
	line := field[pos.y]
	lineInRunes := []rune(line)
	lineInRunes[pos.x] = pipe
	field[pos.y] = string(lineInRunes)
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

func isValid(pos position, field []string) bool {
	height := len(field)
	width := len(field[0])
	return pos.x >= 0 && pos.x < width && pos.y >= 0 && pos.y < height
}

func emptyField(plan plan) []string {
	height := len(plan.field)
	width := len(plan.field[0])
	emptyField := make([]string, height)
	for i := range emptyField {
		emptyField[i] = strings.Repeat(".", width)
	}
	return emptyField
}

func countEnclosed(line string) int {
	count := 0
	inLoop := false
	previousPipe := ' '
	for _, pipe := range line {
		switch pipe {
		case '|':
			inLoop = !inLoop
			previousPipe = pipe
			break
		case 'F':
			inLoop = !inLoop
			previousPipe = pipe
			break
		case 'J':
			if previousPipe != 'F' {
				inLoop = !inLoop
				previousPipe = pipe
			}
			break
		case 'L':
			inLoop = !inLoop
			previousPipe = pipe
			break
		case '7':
			if previousPipe != 'L' {
				inLoop = !inLoop
				previousPipe = pipe
			}
			break
		case '-':
			break
		default:
			if inLoop {
				count++
			}
			break
		}
	}
	return count
}
