package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {

	var lines = read("input/1")

	result := 0
	for _, line := range lines {
		value := calibrationValue(line)
		result += value
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

	var lines = make([]string, 3)

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	return lines
}

func calibrationValue(line string) int {
	if len(line) == 0 {
		return 0
	}
	re := regexp.MustCompile(`[0-9]`)
	parts := re.FindAllString(line, -1)

	first, _ := strconv.Atoi(parts[0])
	last, _ := strconv.Atoi(parts[len(parts)-1])

	return first*10 + last
}
