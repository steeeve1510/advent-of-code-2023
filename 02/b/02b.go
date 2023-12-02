package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cubes struct {
	amount int
	color  string
}

type draw struct {
	cubes []cubes
}

type game struct {
	id    int
	draws []draw
}

func main() {

	lines := read("input/1")
	games := make([]game, 0)
	for _, line := range lines {
		game := parse(line)
		if game.id < 0 {
			continue
		}
		games = append(games, game)
	}

	result := 0
	for _, game := range games {
		power := getPower(game)
		result = result + power
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

func parse(line string) game {
	if len(line) == 0 {
		return game{id: -1, draws: make([]draw, 0)}
	}

	parts1 := strings.Split(line, ":")
	parts2 := strings.Split(parts1[1], ";")

	gameRaw := parts1[0]
	id := parseGameId(gameRaw)

	draws := make([]draw, 0)
	for _, part := range parts2 {
		draw := parseDraw(part)
		draws = append(draws, draw)
	}

	return game{id: id, draws: draws}
}

func parseGameId(gameRaw string) int {
	parts := strings.Split(gameRaw, " ")
	id, _ := strconv.Atoi(parts[1])
	return id
}

func parseDraw(cubesRaw string) draw {
	parts := strings.Split(cubesRaw, ",")

	result := make([]cubes, 0)
	for _, part := range parts {
		xs := strings.Split(part, " ")
		amount, _ := strconv.Atoi(xs[1])
		c := cubes{amount: amount, color: xs[2]}
		result = append(result, c)
	}

	return draw{cubes: result}
}

func getPower(game game) int {
	allCubes := make([]cubes, 0)
	for _, draw := range game.draws {
		for _, cubes := range draw.cubes {
			allCubes = append(allCubes, cubes)
		}
	}

	red := 0
	green := 0
	blue := 0

	for _, cubes := range allCubes {
		if cubes.color == "red" {
			if cubes.amount > red {
				red = cubes.amount
			}
		} else if cubes.color == "green" {
			if cubes.amount > green {
				green = cubes.amount
			}
		} else if cubes.color == "blue" {
			if cubes.amount > blue {
				blue = cubes.amount
			}
		}
	}

	return red * green * blue
}
