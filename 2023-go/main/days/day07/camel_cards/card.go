package camel_cards

import (
	"github.com/shadowradiance/advent-of-code/2023-go/util"
)

type Card string

type CardRank int

func (card Card) Rank(part int) CardRank {
	switch card {
	case "A":
		return 14
	case "K":
		return 13
	case "Q":
		return 12
	case "J":
		if part == 1 {
			return 11
		} else {
			return 1
		}
	case "T":
		return 10
	default:
		return CardRank(util.ConvertNumeric(string(card)))
	}
}
