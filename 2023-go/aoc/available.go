package aoc

import (
	"github.com/shadowradiance/advent-of-code/2023-go/days"
	"github.com/shadowradiance/advent-of-code/2023-go/days/day01"
	"github.com/shadowradiance/advent-of-code/2023-go/days/day02"
	"github.com/shadowradiance/advent-of-code/2023-go/days/day03"
)

func Available() []days.DayInterface {
	return []days.DayInterface{
		day01.Solution{},
		day02.Solution{},
		day03.Solution{},
	}
}
