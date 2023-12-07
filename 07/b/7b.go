package main

import (
	"bufio"
	"cmp"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type hand struct {
	cards    string
	bid      int
	strength int
}

func main() {

	lines := read("input/1")
	hands := parse(lines)

	for i, hand := range hands {
		strength := getStrength(hand)
		hand.strength = strength
		hands[i] = hand
	}

	slices.SortFunc(
		hands,
		func(a, b hand) int {
			return cmp.Compare(a.strength, b.strength)
		},
	)

	result := 0
	for i, hand := range hands {
		result += (i + 1) * hand.bid
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

func parse(lines []string) []hand {
	result := make([]hand, 0)
	for _, line := range lines {
		hand := parseHand(line)
		result = append(result, hand)
	}
	return result
}

func parseHand(line string) hand {
	parts := strings.Split(line, " ")
	cards := parts[0]
	bid, _ := strconv.Atoi(parts[1])
	return hand{cards: cards, bid: bid}
}

func getStrength(hand hand) int {
	values := getValues(hand)

	strength := 0
	if slices.Equal(values, []int{1, 1, 1, 1, 1}) {
		strength = 1
	} else if slices.Equal(values, []int{1, 1, 1, 2}) {
		strength = 2
	} else if slices.Equal(values, []int{1, 2, 2}) {
		strength = 3
	} else if slices.Equal(values, []int{1, 1, 3}) {
		strength = 4
	} else if slices.Equal(values, []int{2, 3}) {
		strength = 5
	} else if slices.Equal(values, []int{1, 4}) {
		strength = 6
	} else if slices.Equal(values, []int{5}) {
		strength = 7
	}
	strength = strength * int(math.Pow(10, 12))

	cardScores := "J23456789TQKA"
	for i, card := range hand.cards {
		score := strings.IndexRune(cardScores, card) + 1
		power := (len(hand.cards) - i - 1) * 2
		factor := int(math.Pow(10, float64(power)))
		strength += score * factor
	}

	return strength
}

func getValues(hand hand) []int {
	counts := make(map[rune]int)
	for _, card := range hand.cards {
		_, isSet := counts[card]
		if !isSet {
			counts[card] = 0
		}
		counts[card]++
	}

	jokers, isSet := counts['J']
	if !isSet {
		jokers = 0
	}
	delete(counts, 'J')

	values := make([]int, 0)
	for _, count := range counts {
		values = append(values, count)
	}
	slices.Sort(values)

	if len(values) == 0 {
		values = append(values, 0)
	}
	values[len(values)-1] += jokers
	return values
}
