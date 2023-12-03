package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"unicode"
)

func main() {

	lines := read("input/1")

	width := len(lines[0])
	height := len(lines)

	result := 0
	re := regexp.MustCompile(`\*`)
	for y, line := range lines {
		gearIxs := re.FindAllStringIndex(line, -1)
		for _, gearIx := range gearIxs {
			x := gearIx[0]
			ixs := getSurrounddingIxs(x, y, width, height)

			partNumbers := getPartNumbers(ixs, lines)
			if len(partNumbers) == 2 {
				result += partNumbers[0] * partNumbers[1]
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
		lines = append(lines, line)
	}

	return lines
}

func getPartNumber(x int, y int, lines []string) int {
	return getNumber(x, lines[y])
}

func getNumber(position int, line string) int {
	if !isNumber(line[position]) {
		return -1
	}
	startIndex := position
	for startIndex > 0 {
		nextIndex := startIndex - 1
		isNumber := isNumber(line[nextIndex])
		if !isNumber {
			break
		}
		startIndex = nextIndex
	}

	endIndex := position + 1
	for endIndex < len(line) {
		isNumber := isNumber(line[endIndex])
		if !isNumber {
			break
		}
		endIndex++
	}

	result, _ := strconv.Atoi(line[startIndex:endIndex])
	return result
}

func isNumber(b byte) bool {
	r := rune(b)
	return unicode.IsDigit(r)
}

func getSurrounddingIxs(x int, y int, width int, height int) [][]int {
	ixs := [][]int{
		{x - 1, y},
		{x - 1, y - 1},
		{x, y - 1},
		{x + 1, y - 1},
		{x + 1, y},
		{x + 1, y + 1},
		{x, y + 1},
		{x - 1, y + 1},
	}
	result := make([][]int, 0)
	for _, ix := range ixs {
		_x := ix[0]
		_y := ix[1]
		if 0 <= _x && _x < width && 0 <= _y && _y < height {
			result = append(result, ix)
		}
	}
	return result
}

func getPartNumbers(ixs [][]int, lines []string) []int {
	numbers := make([]int, 0)
	for _, ix := range ixs {
		x := ix[0]
		y := ix[1]
		number := getPartNumber(x, y, lines)
		if number >= 0 && !slices.Contains(numbers, number) {
			numbers = append(numbers, number)
		}
	}
	return numbers
}
