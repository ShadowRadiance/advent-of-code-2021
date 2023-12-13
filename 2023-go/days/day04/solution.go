package day04

import (
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/shadowradiance/advent-of-code/2023-go/util"
)

type Solution struct{}

func (Solution) Part01(input string) string {
	lines := strings.Split(input, "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	cards := buildCards(lines)

	sum := 0
	for _, card := range cards {
		sum += card.score()
	}

	return strconv.Itoa(sum)
}

func (Solution) Part02(input string) string {
	lines := strings.Split(input, "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}
	cards := buildCards(lines)
	counts := make([]int, len(cards)+1) // 0 is unused
	for _, card := range cards {
		counts[card.id] = 1
	}

	for _, card := range cards {
		for j := 1; j <= card.matches(); j++ {
			counts[card.id+j] += counts[card.id]
		}
	}

	sum := 0
	for _, card := range cards {
		sum += counts[card.id]
	}

	return strconv.Itoa(sum)
}

func buildCards(lines []string) []Card {
	cards := make([]Card, 0)
	for _, line := range lines {
		if len(line) != 0 {
			cards = append(cards, NewCard(line))
		}
	}
	return cards
}

type Card struct {
	id             int
	winningNumbers []int
	yourNumbers    []int
}

func NewCard(s string) Card {
	var cardParser = regexp.MustCompile("^Card +(\\d+): +([^|]*) \\| (.*)$")
	var numSplitter = regexp.MustCompile(" +")
	matches := cardParser.FindStringSubmatch(s)

	return Card{
		id:             util.ConvertNumeric(matches[1]),
		winningNumbers: util.MapStringsToIntegers(numSplitter.Split(matches[2], -1)),
		yourNumbers:    util.MapStringsToIntegers(numSplitter.Split(matches[3], -1)),
	}
}

func (c Card) matches() (matches int) {
	matches = 0
	sort.Ints(c.winningNumbers)
	for _, number := range c.yourNumbers {
		// if number in winning numbers
		if util.IntSliceContainsInt(c.winningNumbers, number) {
			matches++
		}
	}
	return
}

func (c Card) score() int {
	matches := c.matches()
	if matches == 0 {
		return 0
	}
	return int(math.Pow(2, float64(matches-1)))
}

type ById []Card

func (a ById) Len() int           { return len(a) }
func (a ById) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ById) Less(i, j int) bool { return a[i].id < a[j].id }
