package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type card struct {
	id             int
	winningNumbers []int
	numbers        []int
	count          int
}

func main() {

	lines := read("input/1")
	cards := make([]card, 0)
	for _, line := range lines {
		card := parse(line)
		cards = append(cards, card)
	}

	for i, card := range cards {
		matches := countMatches(card)
		for j := 1; j <= matches; j++ {
			cards[i+j].count = cards[i+j].count + card.count
		}
	}
	result := 0
	for _, card := range cards {
		result += card.count
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

func parse(line string) card {
	parts1 := strings.Split(line, ": ")
	parts2 := strings.Split(parts1[1], " | ")

	id := parseCardId(parts1[0])
	winningNumbers := parseNumbers(parts2[0])
	numbers := parseNumbers(parts2[1])

	return card{id: id, winningNumbers: winningNumbers, numbers: numbers, count: 1}
}

func parseCardId(gameRaw string) int {
	parts := strings.Split(gameRaw, " ")
	id, _ := strconv.Atoi(parts[1])
	return id
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

func countMatches(card card) int {
	matches := 0
	for _, number := range card.numbers {
		contained := slices.Contains(card.winningNumbers, number)
		if contained {
			matches++
		}
	}
	return matches
}
