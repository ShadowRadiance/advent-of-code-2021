package main

import (
	"bufio"
	"fmt"
	"github.com/shadowradiance/advent-of-code/2023-go/days"
	"os"
	"strconv"

	"github.com/shadowradiance/advent-of-code/2023-go/aoc"
)

var scanner = bufio.NewScanner(os.Stdin)

func userInput(output string) string {
	fmt.Print(output)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	var err error

	aoc.RegisterDay(days.Day01{})
	//aoc.RegisterDay(days.Day02{})
	//aoc.RegisterDay(days.Day03{})
	//aoc.RegisterDay(days.Day04{})
	//aoc.RegisterDay(days.Day05{})
	//aoc.RegisterDay(days.Day06{})
	//aoc.RegisterDay(days.Day07{})
	//aoc.RegisterDay(days.Day08{})
	//aoc.RegisterDay(days.Day09{})
	//aoc.RegisterDay(days.Day10{})
	//aoc.RegisterDay(days.Day11{})
	//aoc.RegisterDay(days.Day12{})
	//aoc.RegisterDay(days.Day13{})
	//aoc.RegisterDay(days.Day14{})
	//aoc.RegisterDay(days.Day15{})
	//aoc.RegisterDay(days.Day16{})
	//aoc.RegisterDay(days.Day17{})
	//aoc.RegisterDay(days.Day18{})
	//aoc.RegisterDay(days.Day19{})
	//aoc.RegisterDay(days.Day20{})
	//aoc.RegisterDay(days.Day21{})
	//aoc.RegisterDay(days.Day22{})
	//aoc.RegisterDay(days.Day23{})
	//aoc.RegisterDay(days.Day24{})
	//aoc.RegisterDay(days.Day25{})

	dayStr := userInput("Day [Enter for all]: ")

	if dayStr == "" {
		aoc.RunAllParts()
	} else {
		if day, err := strconv.Atoi(dayStr); err == nil {
			partStr := userInput("Part [Enter for all]: ")
			if partStr == "" {
				aoc.RunOnePart(day, 1)
				aoc.RunOnePart(day, 2)
			} else {
				if part, err := strconv.Atoi(partStr); err == nil {
					aoc.RunOnePart(day, part)
				}
			}
		}
	}
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Done!")
}
