package days

import (
	"github.com/shadowradiance/advent-of-code/2023-go/util"
	"strconv"
	"strings"
)

type Day02 struct{}

type Set struct {
	blue  int
	red   int
	green int
}

type Game struct {
	id   int
	sets []Set
}

func makeSet(s string) (set Set) {
	cubeStrings := strings.Split(s, ", ")
	for _, cubeString := range cubeStrings {
		pair := strings.Split(cubeString, " ")
		switch pair[1] {
		case "blue":
			set.blue, _ = strconv.Atoi(pair[0])
		case "red":
			set.red, _ = strconv.Atoi(pair[0])
		case "green":
			set.green, _ = strconv.Atoi(pair[0])
		}
	}
	return
}

func makeGame(line string) (game Game) {
	parts := strings.Split(line, ": ")
	gameString := parts[0]
	game.id, _ = strconv.Atoi(gameString[5:])
	setStrings := strings.Split(parts[1], "; ")
	for _, setString := range setStrings {
		game.sets = append(game.sets, makeSet(setString))
	}
	return
}

func gamePossible(game Game, bag Set) bool {
	for _, set := range game.sets {
		if set.red > bag.red || set.green > bag.green || set.blue > bag.blue {
			return false
		}
	}
	return true
}

func (Day02) Part01(input string) string {
	lines := strings.Split(input, "\n")

	games := make([]Game, 0, len(lines))
	for _, line := range lines {
		if len(line) > 0 {
			games = append(games, makeGame(line))
		}
	}

	var sum int
	bag := Set{red: 12, green: 13, blue: 14}
	for _, game := range games {
		if gamePossible(game, bag) {
			sum += game.id
		}
	}

	return strconv.Itoa(sum)
}

// ------------------

func powerOfSet(set Set) int {
	return set.red * set.green * set.blue
}

func minimalSetFor(bag Set, set Set) Set {
	return Set{
		red:   util.MaxInt(bag.red, set.red),
		green: util.MaxInt(bag.green, set.green),
		blue:  util.MaxInt(bag.blue, set.blue),
	}
}

func (Day02) Part02(input string) string {
	lines := strings.Split(input, "\n")

	games := make([]Game, 0, len(lines))
	for _, line := range lines {
		if len(line) > 0 {
			games = append(games, makeGame(line))
		}
	}
	var power int
	for _, game := range games {
		minimalBag := Set{}
		for _, set := range game.sets {
			minimalBag = minimalSetFor(minimalBag, set)
		}
		power += powerOfSet(minimalBag)
	}

	return strconv.Itoa(power)
}
