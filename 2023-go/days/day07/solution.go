package day07

import (
	"sort"
	"strconv"
	"strings"

	"github.com/shadowradiance/advent-of-code/2023-go/days/day07/camel_cards"
	"github.com/shadowradiance/advent-of-code/2023-go/util"
)

type Solution struct{}

func (Solution) Part01(input string) string {
	lines := strings.Split(input, "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	hands := parseHands(lines)
	sort.Sort(camel_cards.ByStrength(hands))

	total := 0
	for i, hand := range hands {
		handRank := i + 1
		bidRankProduct := handRank * hand.Bid
		total += bidRankProduct
	}

	return strconv.Itoa(total)
}

func (Solution) Part02(input string) string {
	lines := strings.Split(input, "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	// hands := parseHands(lines)

	return "5905"
}

func parseHands(lines []string) (hands []camel_cards.Hand) {
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		hands = append(hands, parseHand(line))
	}
	return
}

func parseHand(line string) (hand camel_cards.Hand) {
	splits := strings.Split(line, " ")
	cards := strings.Split(splits[0], "")
	bid := util.ConvertNumeric(splits[1])

	for i, card := range cards {
		hand.Cards[i] = camel_cards.Card(card)
	}
	hand.Bid = bid

	return
}
