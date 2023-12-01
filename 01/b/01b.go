package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
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

	line1 := strings.ReplaceAll(line, "one", "one1one")
	line2 := strings.ReplaceAll(line1, "two", "two2two")
	line3 := strings.ReplaceAll(line2, "three", "three3three")
	line4 := strings.ReplaceAll(line3, "four", "four4four")
	line5 := strings.ReplaceAll(line4, "five", "five5five")
	line6 := strings.ReplaceAll(line5, "six", "six6six")
	line7 := strings.ReplaceAll(line6, "seven", "seven7seven")
	line8 := strings.ReplaceAll(line7, "eight", "eight8eight")
	line9 := strings.ReplaceAll(line8, "nine", "nine9nine")

	re := regexp.MustCompile(`[0-9]`)
	parts := re.FindAllString(line9, -1)

	first, _ := strconv.Atoi(parts[0])
	last, _ := strconv.Atoi(parts[len(parts)-1])

	return first*10 + last
}
