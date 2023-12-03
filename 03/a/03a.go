package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {

	lines := read("input/1")

	width := len(lines[0]) + 1
	plan := ""
	for _, line := range lines {
		plan = plan + line + "."
	}

	reNumbers := regexp.MustCompile(`[0-9]+`)
	parts := reNumbers.FindAllStringIndex(plan, -1)

	reSymbols := regexp.MustCompile(`[^0-9\.]`)

	partNumbers := make([]int, 0)
	for _, part := range parts {
		number, _ := strconv.Atoi(plan[part[0]:part[1]])
		surroudings := getSurroundings(part, plan, width)
		hasSymbol := reSymbols.MatchString(surroudings)
		if hasSymbol {
			partNumbers = append(partNumbers, number)
		}
	}

	result := 0
	for _, number := range partNumbers {
		result = result + number
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

func getSurroundings(indexes []int, plan string, width int) string {
	start := indexes[0]
	end := indexes[1]

	ixs := make([]int, 0)
	for _, i := range getRange(start, end) {
		up := i - width
		down := i + width

		left := i - 1
		right := i + 1

		leftUp := up - 1
		rightUp := up + 1
		leftDown := down - 1
		rightDown := down + 1

		ixs = safeAppend(ixs, left, len(plan))
		ixs = safeAppend(ixs, leftUp, len(plan))
		ixs = safeAppend(ixs, up, len(plan))
		ixs = safeAppend(ixs, rightUp, len(plan))
		ixs = safeAppend(ixs, right, len(plan))
		ixs = safeAppend(ixs, rightDown, len(plan))
		ixs = safeAppend(ixs, down, len(plan))
		ixs = safeAppend(ixs, leftDown, len(plan))
	}

	result := ""
	for _, i := range ixs {
		result = result + plan[i:i+1]
	}

	return result
}

func getRange(start int, end int) []int {
	result := make([]int, 0)
	for i := start; i < end; i++ {
		result = append(result, i)
	}
	return result
}

func safeAppend(ixs []int, value int, maxWidth int) []int {
	if value >= maxWidth || value < 0 {
		return ixs
	}
	return append(ixs, value)
}
